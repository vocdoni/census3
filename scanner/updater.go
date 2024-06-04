package scanner

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/db"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/helpers/web3"
	"github.com/vocdoni/census3/scanner/providers/manager"
	web3provider "github.com/vocdoni/census3/scanner/providers/web3"
	"go.vocdoni.io/dvote/log"
	"go.vocdoni.io/dvote/util"
)

const (
	coolDown       = 15 * time.Second
	UPDATE_TIMEOUT = 5 * time.Minute
)

type UpdateRequest struct {
	Address       common.Address
	ChainID       uint64
	Type          uint64
	CreationBlock uint64
	EndBlock      uint64
	LastBlock     uint64
}

type Updater struct {
	ctx    context.Context
	cancel context.CancelFunc

	db        *db.DB
	networks  *web3.Web3Pool
	providers *manager.ProviderManager
	queue     map[string]UpdateRequest
	queueMtx  sync.Mutex
	waiter    sync.WaitGroup
}

func NewUpdater(db *db.DB, networks *web3.Web3Pool, pm *manager.ProviderManager) *Updater {
	return &Updater{
		db:        db,
		networks:  networks,
		providers: pm,
		queue:     make(map[string]UpdateRequest),
	}
}

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
					log.Error("Error processing update request: %w", err)
				}
			}
		}
	}()
}

func (u *Updater) Stop() {
	u.cancel()
	u.waiter.Wait()
}

func (u *Updater) RequestStatus(id string) UpdateRequest {
	u.queueMtx.Lock()
	defer u.queueMtx.Unlock()
	req := u.queue[id]
	if req.LastBlock >= req.EndBlock {
		delete(u.queue, id)
	}
	return u.queue[id]
}

func (u *Updater) AddRequest(req UpdateRequest) {
	if req.ChainID == 0 || req.Type == 0 || req.CreationBlock == 0 || req.EndBlock == 0 {
		return
	}
	if req.CreationBlock >= req.EndBlock || req.LastBlock >= req.EndBlock {
		return
	}
	id := util.RandomHex(16)
	u.queueMtx.Lock()
	defer u.queueMtx.Unlock()
	u.queue[id] = req
}

func (u *Updater) IsEmpty() bool {
	u.queueMtx.Lock()
	defer u.queueMtx.Unlock()
	return len(u.queue) == 0
}

func (u *Updater) process() error {
	// make a copy of current queue
	u.queueMtx.Lock()
	queue := map[string]UpdateRequest{}
	for k, v := range u.queue {
		queue[k] = v
	}
	u.queueMtx.Unlock()
	// iterate over the current queue items
	for id, req := range queue {
		internalCtx, cancel := context.WithTimeout(u.ctx, UPDATE_TIMEOUT)
		defer cancel()
		// get the provider by token type
		provider, err := u.providers.GetProvider(u.ctx, req.Type)
		if err != nil {
			return err
		}
		// if the token is a external token, return an error
		if provider.IsExternal() {
			return fmt.Errorf("external providers are not supported yet")
		}
		// set the reference of the token to update in the provider
		if err := provider.SetRef(web3provider.Web3ProviderRef{
			HexAddress:    req.Address.Hex(),
			ChainID:       req.ChainID,
			CreationBlock: req.CreationBlock,
		}); err != nil {
			return err
		}
		// get current token holders from database
		results, err := u.db.QueriesRO.ListTokenHolders(internalCtx, queries.ListTokenHoldersParams{
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
		if err := provider.SetLastBalances(internalCtx, nil, currentHolders, req.LastBlock); err != nil {
			return err
		}
		// get range balances from the provider, it will check itereate again
		// over transfers logs, checking if there are new transfers using the
		// bloom filter associated to the token
		rangeBalances, newTransfers, lastBlock, synced, totalSupply, err := provider.HoldersBalances(internalCtx, nil, req.EndBlock)
		if err != nil {
			return err
		}
		// update the token last
		if synced {
			req.LastBlock = req.EndBlock
		} else {
			req.LastBlock = lastBlock
		}
		// save the new balances in the database
		created, updated, err := SaveHolders(u.db, internalCtx, ScannerToken{
			Address: req.Address,
			ChainID: req.ChainID,
		}, rangeBalances, newTransfers, lastBlock, synced, totalSupply)
		if err != nil {
			return err
		}
		log.Debugw("missing token holders balances updated",
			"token", req.Address.Hex(),
			"chainID", req.ChainID,
			"created", created,
			"updated", updated)
		// update the request in the queue
		u.queueMtx.Lock()
		u.queue[id] = req
		u.queueMtx.Unlock()
	}
	return nil
}
