package scanner

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/db"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/helpers/web3"
	"github.com/vocdoni/census3/scanner/filter"
	"github.com/vocdoni/census3/scanner/providers/manager"
	web3provider "github.com/vocdoni/census3/scanner/providers/web3"
	"go.vocdoni.io/dvote/log"
)

// UpdateRequest is a struct to request a token update but also to query about
// the status of a request that is being processed.
type UpdateRequest struct {
	Address       common.Address
	ChainID       uint64
	ExternalID    string
	Type          uint64
	CreationBlock uint64
	EndBlock      uint64
	LastBlock     uint64
	Done          bool
	Initial       bool

	TotalLogs                 uint64
	TotalNewLogs              uint64
	TotalAlreadyProcessedLogs uint64
	LastTotalSupply           *big.Int
}

// Updater is a struct to manage the update requests of the tokens. It will
// iterate over the requests, repeating the process of getting the token holders
// balances and saving them in the database until the last block is greater or
// equal to the end block. The end block is the block number where the token
// holders balances are up to date. The holders providers must include an
// instance of a TokenFilter to store the processed transactions to avoid
// re-processing them, but also rescanning a synced token to find missing
// transactions.
type Updater struct {
	ctx    context.Context
	cancel context.CancelFunc

	db          *db.DB
	networks    *web3.Web3Pool
	providers   *manager.ProviderManager
	queue       map[string]*UpdateRequest
	queueMtx    sync.Mutex
	waiter      sync.WaitGroup
	filtersPath string
}

// NewUpdater creates a new instance of Updater.
func NewUpdater(db *db.DB, networks *web3.Web3Pool, pm *manager.ProviderManager,
	filtersPath string,
) *Updater {
	return &Updater{
		db:          db,
		networks:    networks,
		providers:   pm,
		queue:       make(map[string]*UpdateRequest),
		filtersPath: filtersPath,
	}
}

// Start starts the updater process in a goroutine.
func (u *Updater) Start(ctx context.Context) {
	u.ctx, u.cancel = context.WithCancel(ctx)
	u.waiter.Add(1)
	go func() {
		defer u.waiter.Done()
		for {
			select {
			case <-u.ctx.Done():
				return
			default:
				if u.IsEmpty() {
					time.Sleep(coolDown)
					continue
				}
				if err := u.process(); err != nil {
					log.Errorf("Error processing update request: %v", err)
				}
			}
		}
	}()
}

// Stop stops the updater process.
func (u *Updater) Stop() {
	u.cancel()
	u.waiter.Wait()
}

// RequestStatus returns the status of a request by its ID. If the request is
// done, it will be removed from the queue. If the request is not found, it will
// return an error.
func (u *Updater) RequestStatus(id string, deleteOnDone bool) *UpdateRequest {
	u.queueMtx.Lock()
	defer u.queueMtx.Unlock()
	req, ok := u.queue[id]
	if !ok {
		return nil
	}
	res := *req
	if deleteOnDone && req.Done {
		delete(u.queue, id)
	}
	return &res
}

// SetRequest adds a new request to the queue. It will return an error if the
// request is missing required fields or the block range is invalid. The request
// will be added to the queue with a given ID.
func (u *Updater) SetRequest(id string, req *UpdateRequest) error {
	// check required fields
	if id == "" {
		return fmt.Errorf("missing request ID")
	}
	if req.ChainID == 0 || req.Type == 0 || req.CreationBlock == 0 || req.EndBlock == 0 {
		return fmt.Errorf("missing required fields")
	}
	// ensure the block range is valid
	if req.CreationBlock >= req.EndBlock {
		return fmt.Errorf("invalid block range")
	}
	// set the last block to the creation block to start the process from there
	// if it is not set by the client
	if req.LastBlock == 0 {
		req.LastBlock = req.CreationBlock
	}
	u.queueMtx.Lock()
	defer u.queueMtx.Unlock()
	u.queue[id] = req
	return nil
}

// AddRequest adds a new request to the queue. It will return an error if the
// request is missing required fields or the block range is invalid. The request
// will be added to the queue with a ID generated from the address, chainID and
// externalID, that will be returned to allow the client to query the status of
// the request.
func (u *Updater) AddRequest(req *UpdateRequest) (string, error) {
	id, err := RequestID(req.Address, req.ChainID, req.ExternalID)
	if err != nil {
		return "", err
	}
	if err := u.SetRequest(id, req); err != nil {
		return "", err
	}
	return id, nil
}

// IsEmpty returns true if the queue is empty.
func (u *Updater) IsEmpty() bool {
	u.queueMtx.Lock()
	defer u.queueMtx.Unlock()
	return len(u.queue) == 0
}

// RequestID returns the ID of a request given the address, chainID and external
// ID. The raw ID is a string with the format "chainID:address:externalID". The
// resulting ID is the first 4 bytes of the hash of the raw ID using the sha256
// algorithm, encoded in hexadecimal.
func RequestID(address common.Address, chainID uint64, externalID string) (string, error) {
	rawID := fmt.Sprintf("%d:%s:%s", chainID, address.Hex(), externalID)
	hashFn := sha256.New()
	if _, err := hashFn.Write([]byte(rawID)); err != nil {
		return "", err
	}
	bHash := hashFn.Sum(nil)
	return hex.EncodeToString(bHash[:4]), nil
}

// process iterates over the current queue items, getting the token holders
// balances and saving them in the database until the last block is greater or
// equal to the end block. It updates th status of the request in the queue. It
// will return an error if the provider is not found, the token is external or
// there is an error getting the token holders balances.
func (u *Updater) process() error {
	// make a copy of current queue
	u.queueMtx.Lock()
	queue := map[string]*UpdateRequest{}
	for k, v := range u.queue {
		queue[k] = v
	}
	u.queueMtx.Unlock()
	// iterate over the current queue items
	for id, req := range queue {
		// check if the request is done
		if req.Done {
			continue
		}
		log.Infow("rescanning token",
			"address", req.Address.Hex(),
			"from", req.CreationBlock,
			"to", req.EndBlock,
			"current", req.LastBlock)
		ctx, cancel := context.WithTimeout(u.ctx, UPDATE_TIMEOUT)
		defer cancel()
		// get the provider by token type
		provider, err := u.providers.GetProvider(ctx, req.Type)
		if err != nil {
			return err
		}
		// if the token is a external token, return an error
		if !provider.IsExternal() {
			// load filter of the token from the database
			filter, err := filter.LoadFilter(u.filtersPath, req.Address, req.ChainID, req.ExternalID)
			if err != nil {
				return err
			}
			// commit the filter when the function finishes
			defer func() {
				if err := filter.Commit(); err != nil {
					log.Error(err)
					return
				}
			}()
			// set the reference of the token to update in the provider
			if err := provider.SetRef(web3provider.Web3ProviderRef{
				HexAddress:    req.Address.Hex(),
				ChainID:       req.ChainID,
				CreationBlock: req.CreationBlock,
				Filter:        filter,
			}); err != nil {
				return err
			}
		}
		// update the last block number of the provider to the last block of
		// the request
		provider.SetLastBlockNumber(req.EndBlock)
		// get current token holders from database
		results, err := u.db.QueriesRO.ListTokenHolders(ctx, queries.ListTokenHoldersParams{
			TokenID: req.Address.Bytes(),
			ChainID: req.ChainID,
		})
		if err != nil {
			return nil
		}
		currentHolders := map[common.Address]*big.Int{}
		for _, holder := range results {
			bBalance, ok := new(big.Int).SetString(holder.Balance, 10)
			if !ok {
				return fmt.Errorf("error parsing holder balance from database")
			}
			currentHolders[common.Address(holder.HolderID)] = bBalance
		}
		// set the current holders in the provider
		if err := provider.SetLastBalances(ctx, nil, currentHolders, req.LastBlock); err != nil {
			return err
		}
		// update with expected results in the queue once the function finishes
		defer func() {
			log.Infow("updating request in the queue", "lastBlock", req.LastBlock, "done", req.Done)
			u.queueMtx.Lock()
			u.queue[id] = req
			u.queueMtx.Unlock()
		}()
		// get range balances from the provider, it will check itereate again
		// over transfers logs, checking if there are new transfers using the
		// bloom filter associated to the token
		balances, delta, err := provider.HoldersBalances(ctx, nil, req.LastBlock)
		// update the token last block in the request before checking the error
		if delta != nil {
			req.TotalLogs += delta.LogsCount
			req.TotalNewLogs += delta.NewLogsCount
			req.TotalAlreadyProcessedLogs += delta.AlreadyProcessedLogsCount
			req.LastTotalSupply = delta.TotalSupply

			req.Done = delta.Synced
			if delta.Synced {
				req.LastBlock = req.EndBlock
			} else if delta.Block >= req.LastBlock {
				req.LastBlock = delta.Block
			}
		}
		if err != nil {
			return err
		}
		log.Infow("new logs received",
			"address", req.Address.Hex(),
			"from", req.LastBlock,
			"lastBlock", delta.Block,
			"newLogs", delta.NewLogsCount,
			"alreadyProcessedLogs", delta.AlreadyProcessedLogsCount,
			"totalLogs", delta.LogsCount)
		// save the new balances in the database
		created, updated, err := SaveHolders(u.db, ctx, ScannerToken{
			Address: req.Address,
			ChainID: req.ChainID,
		}, balances, delta.NewLogsCount, delta.Block, delta.Synced, delta.TotalSupply)
		if err != nil {
			return err
		}
		log.Debugw("token holders balances updated",
			"token", req.Address.Hex(),
			"chainID", req.ChainID,
			"created", created,
			"updated", updated)
	}
	return nil
}
