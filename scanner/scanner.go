package scanner

import (
	"context"
	"database/sql"
	"errors"
	"math/big"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/db"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/helpers/web3"
	"github.com/vocdoni/census3/scanner/providers/manager"
	web3provider "github.com/vocdoni/census3/scanner/providers/web3"
	"go.vocdoni.io/dvote/log"
)

// ScannerToken includes the information of a token that the scanner needs to
// scan it.
type ScannerToken struct {
	Address       common.Address
	ChainID       uint64
	Type          uint64
	ExternalID    string
	LastBlock     uint64
	CreationBlock uint64
	Ready         bool
	Synced        bool
	totalSupply   *big.Int
}

// Scanner is the scanner that scans the tokens and saves the holders in the
// database. It has a list of tokens to scan and a list of providers to get the
// holders of the tokens. It has a cool down time between iterations to avoid
// overloading the providers.
type Scanner struct {
	ctx             context.Context
	cancel          context.CancelFunc
	db              *db.DB
	updater         *Updater
	networks        *web3.Web3Pool
	providerManager *manager.ProviderManager
	coolDown        time.Duration
	filtersPath     string

	tokens                  []*ScannerToken
	tokensMtx               sync.Mutex
	waiter                  sync.WaitGroup
	latestBlockNumbers      sync.Map
	lastUpdatedBlockNumbers time.Time
}

// NewScanner returns a new scanner instance with the required parameters
// initialized.
func NewScanner(db *db.DB, updater *Updater, networks *web3.Web3Pool, pm *manager.ProviderManager,
	coolDown time.Duration, filtersPath string,
) *Scanner {
	return &Scanner{
		db:                      db,
		updater:                 updater,
		networks:                networks,
		providerManager:         pm,
		coolDown:                coolDown,
		filtersPath:             filtersPath,
		tokens:                  []*ScannerToken{},
		tokensMtx:               sync.Mutex{},
		waiter:                  sync.WaitGroup{},
		latestBlockNumbers:      sync.Map{},
		lastUpdatedBlockNumbers: time.Time{},
	}
}

// Start starts the scanner. It starts a loop that scans the tokens in the
// database and saves the holders in the database. It stops when the context is
// cancelled.
func (s *Scanner) Start(ctx context.Context, concurrentTokens int) {
	if concurrentTokens < 1 {
		concurrentTokens = 1
	}
	s.ctx, s.cancel = context.WithCancel(ctx)
	itCounter := 0
	// keep the latest block numbers updated
	s.waiter.Add(1)
	go func() {
		defer s.waiter.Done()
		s.getLatestBlockNumbersUpdates()
	}()
	// start the scanner loop
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// create some variables to track the loop progress
			itCounter++
			startTime := time.Now()
			// get the tokens to scan
			tokens, err := s.TokensToScan(ctx)
			if err != nil {
				log.Error(err)
				continue
			}
			// calculate number of batches
			sem := make(chan struct{}, concurrentTokens)
			defer close(sem)
			// iterate over the tokens to scan
			var atSyncGlobal atomic.Bool
			atSyncGlobal.Store(true)
			for _, token := range tokens {
				log.Infow("checking token in the updater queue",
					"address", token.Address.Hex(),
					"chainID", token.ChainID,
					"externalID", token.ExternalID)
				// get the request ID of the token in the updater queue
				reqID, err := RequestID(token.Address, token.ChainID, token.ExternalID)
				if err != nil {
					log.Error(err)
					continue
				}
				// get the status of the token in the updater queue
				status := s.updater.RequestStatus(reqID, true)
				if status != nil {
					log.Infow("token status in the updater queue",
						"address", token.Address.Hex(),
						"chainID", token.ChainID,
						"externalID", token.ExternalID,
						"lastBlock", status.LastBlock,
						"lastTotalSupply", status.LastTotalSupply,
						"totalNewLogs", status.TotalNewLogs,
						"totalAlreadyProcessedLogs", status.TotalAlreadyProcessedLogs,
						"totalLogs", status.TotalLogs,
						"done", status.Done)
					// if the token is in the updater queue, update the
					// internal token status and continue to the next token
					// only if the token is done
					defer s.updateInternalTokenStatus(*token, status.LastBlock, status.Done, status.LastTotalSupply)
					if status.Done {
						continue
					}
					atSyncGlobal.Store(false)
				}
				// if it has been processed or it is not in the queue, load
				// the last available block number of the network and
				// enqueue it to the updater queue from the last scanned
				// block
				if iLastNetworkBlock, ok := s.latestBlockNumbers.Load(token.ChainID); ok {
					if lastNetworkBlock, ok := iLastNetworkBlock.(uint64); ok {
						if _, err := s.updater.SetRequest(&UpdateRequest{
							Address:       token.Address,
							ChainID:       token.ChainID,
							Type:          token.Type,
							ExternalID:    token.ExternalID,
							CreationBlock: token.CreationBlock,
							EndBlock:      lastNetworkBlock,
							LastBlock:     token.LastBlock,
						}); err != nil {
							log.Warnw("error enqueuing token", "error", err)
							continue
						}
						log.Infow("token enqueued from the scanner",
							"address", token.Address.Hex(),
							"chainID", token.ChainID,
							"externalID", token.ExternalID,
							"from", token.LastBlock,
							"to", lastNetworkBlock)
					}
				}
			}
			log.Infow("scan iteration finished",
				"iteration", itCounter,
				"duration", time.Since(startTime).Seconds(),
				"atSync", atSyncGlobal.Load())
			// if all the tokens are synced, sleep the cool down time, else,
			// sleep the scan sleep time
			if atSyncGlobal.Load() {
				time.Sleep(s.coolDown)
			} else {
				time.Sleep(scanSleepTime)
			}
		}
	}
}

// Stop stops the scanner. It cancels the context and waits for the scanner to
// finish. It also closes the providers.
func (s *Scanner) Stop() {
	s.cancel()
	s.waiter.Wait()
}

// TokensToScan returns the tokens that the scanner has to scan. It returns the
// the tokens to scan from the database in the following order:
//  1. The tokens that were created in the last 60 minutes and are not synced.
//  2. The rest of no synced tokens, sorted by the difference between their
//     block number and the last block number of their chain.
//  3. The tokens that were synced in previous iterations.
func (s *Scanner) TokensToScan(ctx context.Context) ([]*ScannerToken, error) {
	internalCtx, cancel := context.WithTimeout(ctx, READ_TIMEOUT)
	defer cancel()
	tokens := []*ScannerToken{}
	// get last created tokens from the database to scan them first (1)
	lastNotSyncedTokens, err := s.db.QueriesRO.ListLastNoSyncedTokens(internalCtx)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return nil, err
	}
	// parse last not synced token addresses
	for _, token := range lastNotSyncedTokens {
		totalSupply, ok := new(big.Int).SetString(string(token.TotalSupply), 10)
		if !ok {
			totalSupply = nil
		}
		st := &ScannerToken{
			Address:       common.BytesToAddress(token.ID),
			ChainID:       token.ChainID,
			Type:          token.TypeID,
			ExternalID:    token.ExternalID,
			LastBlock:     uint64(token.LastBlock),
			CreationBlock: uint64(token.CreationBlock),
			Ready:         token.CreationBlock > 0 && token.LastBlock >= token.CreationBlock,
			Synced:        token.Synced,
			totalSupply:   totalSupply,
		}
		if err := s.prepareToken(st); err != nil {
			log.Warnw("error preparing token", "error", err)
			continue
		}
		tokens = append(tokens, st)
	}
	// get old not synced tokens from the database (2)
	oldNotSyncedTokens, err := s.db.QueriesRO.ListOldNoSyncedTokens(internalCtx)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return nil, err
	}
	// if there are old not synced tokens, sort them by nearest to be synced
	// and parse them, if not, continue to avoid web3 calls
	if len(oldNotSyncedTokens) > 0 {
		// sort old not synced tokens by nearest to be synced, that is, the tokens
		// that have the minimum difference between the current block of its chain
		// and the last block scanned by the scanner (retrieved from the database
		// as LastBlock)
		sort.Slice(oldNotSyncedTokens, func(i, j int) bool {
			iRawLastBlock, ok := s.latestBlockNumbers.Load(oldNotSyncedTokens[i].ChainID)
			if !ok {
				return false
			}
			iLastBlock, ok := iRawLastBlock.(uint64)
			if !ok {
				return false
			}
			jRawLastBlock, ok := s.latestBlockNumbers.Load(oldNotSyncedTokens[j].ChainID)
			if !ok {
				return false
			}
			jLastBlock, ok := jRawLastBlock.(uint64)
			if !ok {
				return false
			}
			iBlocksReamining := iLastBlock - uint64(oldNotSyncedTokens[i].LastBlock)
			jBlocksReamining := jLastBlock - uint64(oldNotSyncedTokens[j].LastBlock)
			return iBlocksReamining < jBlocksReamining
		})
		// parse old not synced token addresses
		for _, token := range oldNotSyncedTokens {
			totalSupply, ok := new(big.Int).SetString(string(token.TotalSupply), 10)
			if !ok {
				totalSupply = nil
			}
			st := &ScannerToken{
				Address:       common.BytesToAddress(token.ID),
				ChainID:       token.ChainID,
				Type:          token.TypeID,
				ExternalID:    token.ExternalID,
				LastBlock:     uint64(token.LastBlock),
				CreationBlock: uint64(token.CreationBlock),
				Ready:         token.CreationBlock > 0 && token.LastBlock >= token.CreationBlock,
				Synced:        token.Synced,
				totalSupply:   totalSupply,
			}
			if err := s.prepareToken(st); err != nil {
				log.Warnw("error preparing token", "error", err)
				continue
			}
			tokens = append(tokens, st)
		}
	}
	// get synced tokens from the database to scan them last (3)
	syncedTokens, err := s.db.QueriesRO.ListSyncedTokens(internalCtx)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return nil, err
	}
	for _, token := range syncedTokens {
		totalSupply, ok := new(big.Int).SetString(string(token.TotalSupply), 10)
		if !ok {
			totalSupply = nil
		}
		st := &ScannerToken{
			Address:       common.BytesToAddress(token.ID),
			ChainID:       token.ChainID,
			Type:          token.TypeID,
			ExternalID:    token.ExternalID,
			LastBlock:     uint64(token.LastBlock),
			CreationBlock: uint64(token.CreationBlock),
			Ready:         token.CreationBlock > 0 && token.LastBlock >= token.CreationBlock,
			Synced:        token.Synced,
			totalSupply:   totalSupply,
		}
		if err := s.prepareToken(st); err != nil {
			log.Warnw("error preparing token", "error", err)
			continue
		}
		tokens = append(tokens, st)
	}
	// update the tokens to scan in the scanner and return them
	s.tokensMtx.Lock()
	s.tokens = tokens
	s.tokensMtx.Unlock()
	return tokens, nil
}

func (s *Scanner) updateInternalTokenStatus(token ScannerToken, lastBlock uint64,
	synced bool, totalSupply *big.Int,
) {
	s.tokensMtx.Lock()
	for i, t := range s.tokens {
		if t.Address == token.Address && t.ChainID == token.ChainID && t.ExternalID == token.ExternalID {
			s.tokens[i].LastBlock = lastBlock
			s.tokens[i].Synced = synced
			if totalSupply != nil && totalSupply.Cmp(big.NewInt(0)) > 0 {
				s.tokens[i].totalSupply = totalSupply
				token.totalSupply = totalSupply
			}
			break
		}
	}
	s.tokensMtx.Unlock()
}

func (s *Scanner) prepareToken(token *ScannerToken) error {
	ctx, cancel := context.WithTimeout(s.ctx, UPDATE_TIMEOUT)
	defer cancel()
	// get the provider by token type
	provider, err := s.providerManager.GetProvider(ctx, token.Type)
	if err != nil {
		return err
	}
	// if the token is not ready yet (its creation block has not been
	// calculated yet), calculate it, update the token information and
	// return
	if !provider.IsExternal() && !token.Ready {
		if err := provider.SetRef(web3provider.Web3ProviderRef{
			HexAddress:    token.Address.Hex(),
			ChainID:       token.ChainID,
			CreationBlock: token.CreationBlock,
		}); err != nil {
			return err
		}
		log.Debugw("token not ready yet, calculating creation block and continue",
			"address", token.Address.Hex(),
			"chainID", token.ChainID,
			"externalID", token.ExternalID)
		creationBlock, err := provider.CreationBlock(ctx, []byte(token.ExternalID))
		if err != nil {
			return err
		}
		_, err = s.db.QueriesRW.UpdateTokenBlocks(ctx, queries.UpdateTokenBlocksParams{
			ID:            token.Address.Bytes(),
			ChainID:       token.ChainID,
			ExternalID:    token.ExternalID,
			CreationBlock: int64(creationBlock),
			LastBlock:     int64(creationBlock),
		})
		if err != nil {
			return err
		}
		token.LastBlock = creationBlock
	}
	return nil
}

// SaveHolders saves the given holders in the database. It calls the SaveHolders
// helper function to save the holders and the token status in the database. It
// prints the number of created and updated token holders if there are any, else
// it prints that there are no holders to save. If some error occurs, it returns
// the error.
func (s *Scanner) SaveHolders(ctx context.Context, token ScannerToken,
	holders map[common.Address]*big.Int, newTransfers, lastBlock uint64,
	synced bool, totalSupply *big.Int,
) error {
	log.Infow("saving token status and holders",
		"token", token.Address.Hex(),
		"chainID", token.ChainID,
		"externalID", token.ExternalID,
		"block", lastBlock,
		"holders", len(holders))
	internalCtx, cancel := context.WithTimeout(ctx, SAVE_TIMEOUT)
	defer cancel()
	// print the number of created and updated token holders if there are any,
	// else, print that there are no holders to save
	if len(holders) == 0 {
		log.Debugw("no holders to save",
			"token", token.Address.Hex(),
			"chainID", token.ChainID,
			"externalID", token.ExternalID)
	} else {
		created, updated, err := SaveHolders(s.db, internalCtx, token, holders, newTransfers, lastBlock, synced, totalSupply)
		if err != nil {
			return err
		}
		log.Debugw("committing token holders",
			"token", token.Address.Hex(),
			"chainID", token.ChainID,
			"externalID", token.ExternalID,
			"block", token.LastBlock,
			"synced", token.Synced,
			"created", created,
			"updated", updated)
	}
	log.Debugw("token status saved",
		"synced", synced,
		"token", token.Address.Hex(),
		"chainID", token.ChainID,
		"externalID", token.ExternalID,
		"totalSupply", token.totalSupply.String(),
		"block", lastBlock)
	return nil
}

// getLatestBlockNumbersUpdates gets the latest block numbers of every chain
// and stores them in the scanner. It is executed in a goroutine and it is
// executed every blockNumbersCooldown. It is used to avoid overloading the
// providers with requests to get the latest block number.
func (s *Scanner) getLatestBlockNumbersUpdates() {
	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			if time.Since(s.lastUpdatedBlockNumbers) > blockNumbersCooldown {
				log.Info("getting latest block numbers")
				latestBlockNumbers, err := s.networks.CurrentBlockNumbers(s.ctx)
				if err != nil {
					log.Error(err)
					continue
				}
				for chainID, blockNumber := range latestBlockNumbers {
					s.latestBlockNumbers.Store(chainID, blockNumber)
				}
				s.lastUpdatedBlockNumbers = time.Now()
			}
		}
	}
}
