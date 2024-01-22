package scanner

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/db"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/scanner/providers"
	"github.com/vocdoni/census3/scanner/providers/web3"
	"go.vocdoni.io/dvote/log"
)

// ScannerToken includes the information of a token that the scanner needs to
// scan it.
type ScannerToken struct {
	Address    common.Address
	ChainID    uint64
	Type       uint64
	ExternalID string
	LastBlock  uint64
	Ready      bool
	Synced     bool
}

// Scanner is the scanner that scans the tokens and saves the holders in the
// database. It has a list of tokens to scan and a list of providers to get the
// holders of the tokens. It has a cool down time between iterations to avoid
// overloading the providers.
type Scanner struct {
	ctx       context.Context
	cancel    context.CancelFunc
	db        *db.DB
	networks  web3.NetworkEndpoints
	providers map[uint64]providers.HolderProvider
	coolDown  time.Duration

	tokens    []*ScannerToken
	tokensMtx sync.Mutex
	waiter    sync.WaitGroup
}

// NewScanner returns a new scanner instance with the required parameters
// initialized.
func NewScanner(db *db.DB, networks web3.NetworkEndpoints,
	providers map[uint64]providers.HolderProvider, coolDown time.Duration,
) *Scanner {
	return &Scanner{
		db:        db,
		networks:  networks,
		providers: providers,
		coolDown:  coolDown,
		tokens:    []*ScannerToken{},
		tokensMtx: sync.Mutex{},
		waiter:    sync.WaitGroup{},
	}
}

// Start starts the scanner. It starts a loop that scans the tokens in the
// database and saves the holders in the database. It stops when the context is
// cancelled.
func (s *Scanner) Start(ctx context.Context) {
	s.ctx, s.cancel = context.WithCancel(ctx)
	itCounter := 0
	for {
		select {
		case <-ctx.Done():
			return
		default:
			itCounter++
			startTime := time.Now()
			tokens, err := s.TokensToScan(ctx)
			if err != nil {
				log.Error(err)
				continue
			}
			atSyncGlobal := true
			for _, token := range tokens {
				holders, lastBlock, synced, err := s.ScanHolders(ctx, token)
				if err != nil {
					log.Error(err)
					continue
				}
				if !synced {
					atSyncGlobal = false
				}
				if len(holders) > 0 {
					s.waiter.Add(1)
					go func(t *ScannerToken, h map[common.Address]*big.Int, lb uint64, sy bool) {
						defer s.waiter.Done()
						if err = s.SaveHolders(ctx, t, h, lb, sy); err != nil {
							log.Error(err)
						}
					}(token, holders, lastBlock, synced)
				}
			}
			log.Infow("scan iteration finished",
				"iteration", itCounter,
				"duration", time.Since(startTime).Seconds(),
				"atSync", atSyncGlobal)
			if atSyncGlobal {
				time.Sleep(s.coolDown)
			} else {
				time.Sleep(scanSleepTime)
			}
		}
	}
}

// Stop stops the scanner. It cancels the context and waits for the scanner to
// finish.
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
	internalCtx, cancel := context.WithTimeout(ctx, SCAN_TIMEOUT)
	defer cancel()
	// create a tx to use it in the following queries
	tx, err := s.db.RW.BeginTx(internalCtx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
			log.Error(err)
			return
		}
	}()
	qtx := s.db.QueriesRW.WithTx(tx)

	tokens := []*ScannerToken{}
	// get last created tokens from the database to scan them first
	lastNotSyncedTokens, err := qtx.ListLastNoSyncedTokens(ctx)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return nil, err
	}
	// parse last not synced token addresses
	for _, token := range lastNotSyncedTokens {
		lastBlock := uint64(token.CreationBlock)
		if blockNumber, err := s.db.QueriesRO.LastBlockByTokenID(ctx, token.ID); err == nil {
			lastBlock = blockNumber
		}
		tokens = append(tokens, &ScannerToken{
			Address:    common.BytesToAddress(token.ID),
			ChainID:    token.ChainID,
			Type:       token.TypeID,
			ExternalID: token.ExternalID,
			LastBlock:  lastBlock,
			Ready:      token.CreationBlock > 0,
			Synced:     token.Synced,
		})
	}
	// get old tokens from the database
	oldNotSyncedTokens, err := qtx.ListOldNoSyncedTokens(ctx)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return nil, err
	}
	// get the current block number of every chain
	currentBlockNumbers, err := s.networks.CurrentBlockNumbers(ctx)
	if err != nil {
		return nil, err
	}
	// sort old not synced tokens by nearest to be synced, that is, the tokens
	// that have the minimum difference between the current block of its chain
	// and the last block scanned by the scanner (retrieved from the database
	// as LastBlock)
	sort.Slice(oldNotSyncedTokens, func(i, j int) bool {
		iLastBlock := uint64(oldNotSyncedTokens[i].CreationBlock)
		if oldNotSyncedTokens[i].LastBlock != nil {
			iLastBlock = uint64(oldNotSyncedTokens[i].LastBlock.(int64))
		}
		jLastBlock := uint64(oldNotSyncedTokens[j].CreationBlock)
		if oldNotSyncedTokens[j].LastBlock != nil {
			jLastBlock = uint64(oldNotSyncedTokens[j].LastBlock.(int64))
		}
		iBlocksReamining := currentBlockNumbers[oldNotSyncedTokens[i].ChainID] - uint64(iLastBlock)
		jBlocksReamining := currentBlockNumbers[oldNotSyncedTokens[j].ChainID] - uint64(jLastBlock)
		return iBlocksReamining < jBlocksReamining
	})
	// parse old not synced token addresses
	for _, token := range oldNotSyncedTokens {
		lastBlock := uint64(token.CreationBlock)
		if token.LastBlock != nil {
			lastBlock = uint64(token.LastBlock.(int64))
		}
		tokens = append(tokens, &ScannerToken{
			Address:    common.BytesToAddress(token.ID),
			ChainID:    token.ChainID,
			Type:       token.TypeID,
			ExternalID: token.ExternalID,
			LastBlock:  lastBlock,
			Ready:      token.CreationBlock > 0,
			Synced:     token.Synced,
		})
	}
	// get last created tokens from the database to scan them first
	syncedTokens, err := qtx.ListSyncedTokens(ctx)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return nil, err
	}
	for _, token := range syncedTokens {
		lastBlock := uint64(token.CreationBlock)
		if blockNumber, err := s.db.QueriesRO.LastBlockByTokenID(ctx, token.ID); err == nil {
			lastBlock = blockNumber
		}
		s.tokens = append(s.tokens, &ScannerToken{
			Address:    common.BytesToAddress(token.ID),
			ChainID:    token.ChainID,
			Type:       token.TypeID,
			ExternalID: token.ExternalID,
			LastBlock:  lastBlock,
			Ready:      token.CreationBlock > 0,
			Synced:     token.Synced,
		})
	}
	// close the database tx and commit it
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	// update the tokens to scan in the scanner and return them
	s.tokensMtx.Lock()
	s.tokens = tokens
	s.tokensMtx.Unlock()
	return tokens, nil
}

// ScanHolders scans the holders of the given token. It get the current holders
// from the database, set them into the provider and get the new ones. It
// returns the new holders, the last block scanned and if the token is synced
// after the scan.
func (s *Scanner) ScanHolders(ctx context.Context, token *ScannerToken) (map[common.Address]*big.Int, uint64, bool, error) {
	log.Infow("scanning holders",
		"address", token.Address.Hex(),
		"chainID", token.ChainID,
		"externalID", token.ExternalID)
	internalCtx, cancel := context.WithTimeout(ctx, SCAN_TIMEOUT)
	defer cancel()
	// get the correct token holder for the current token
	provider, exists := s.providers[token.ChainID]
	if !exists {
		return nil, token.LastBlock, token.Synced, fmt.Errorf("no provider for chain %d", token.ChainID)
	}
	// create a tx to use it in the following queries
	tx, err := s.db.RW.BeginTx(internalCtx, nil)
	if err != nil {
		return nil, token.LastBlock, token.Synced, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
			log.Error(err)
		}
	}()
	qtx := s.db.QueriesRW.WithTx(tx)
	// if the provider is not an external one, instance the current token
	if !provider.IsExternal() {
		if err := provider.SetRef(web3.Web3ProviderRef{
			HexAddress: token.Address.Hex(),
			ChainID:    token.ChainID,
		}); err != nil {
			return nil, token.LastBlock, token.Synced, err
		}
		// if the token is not ready yet (its creation block has not been
		// calculated yet), calculate it, update the token information and
		// return
		if !token.Ready {
			creationBlock, err := provider.CreationBlock(internalCtx, []byte(token.ExternalID))
			if err != nil {
				return nil, token.LastBlock, token.Synced, err
			}
			_, err = qtx.UpdateTokenCreationBlock(internalCtx, queries.UpdateTokenCreationBlockParams{
				ID:            token.Address.Bytes(),
				ChainID:       token.ChainID,
				ExternalID:    token.ExternalID,
				CreationBlock: int64(creationBlock),
			})
			if err != nil {
				return nil, token.LastBlock, token.Synced, err
			}
			// close the database tx and commit it
			return nil, token.LastBlock, token.Synced, tx.Commit()
		}
	}
	// get the current token holders from the database
	results, err := qtx.TokenHoldersByTokenIDAndChainIDAndExternalID(internalCtx,
		queries.TokenHoldersByTokenIDAndChainIDAndExternalIDParams{
			TokenID:    token.Address.Bytes(),
			ChainID:    token.ChainID,
			ExternalID: token.ExternalID,
		})
	if err != nil {
		return nil, token.LastBlock, token.Synced, err
	}
	currentHolders := map[common.Address]*big.Int{}
	for _, result := range results {
		currentHolders[common.BytesToAddress(result.ID)] = big.NewInt(0).SetBytes([]byte(result.Balance))
	}
	// close the database tx and commit it
	if err := tx.Commit(); err != nil {
		return nil, token.LastBlock, token.Synced, err
	}
	// set the current holders into the provider and get the new ones
	if err := provider.SetLastBalances(ctx, []byte(token.ExternalID), currentHolders, token.LastBlock); err != nil {
		return nil, token.LastBlock, token.Synced, err
	}
	return provider.HoldersBalances(ctx, []byte(token.ExternalID), token.LastBlock)
}

// SaveHolders saves the given holders in the database. It updates the token
// synced status if it is different from the received one. Then, it creates,
// updates or deletes the token holders in the database depending on the
// calculated balance.
func (s *Scanner) SaveHolders(ctx context.Context, token *ScannerToken,
	holders map[common.Address]*big.Int, lastBlock uint64, synced bool) error {
	log.Debugw("saving token holders",
		"token", token.Address.Hex(),
		"chainID", token.ChainID,
		"externalID", token.ExternalID,
		"block", lastBlock,
		"holders", len(holders))

	s.tokensMtx.Lock()
	for i, t := range s.tokens {
		if t.Address == token.Address && t.ChainID == token.ChainID && t.ExternalID == token.ExternalID {
			s.tokens[i].LastBlock = lastBlock
			s.tokens[i].Synced = synced
			break
		}
	}
	s.tokensMtx.Unlock()

	internalCtx, cancel := context.WithTimeout(ctx, SAVE_TIMEOUT)
	defer cancel()
	// create a tx to use it in the following queries
	tx, err := s.db.RW.BeginTx(internalCtx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
			log.Error(err)
		}
	}()
	qtx := s.db.QueriesRW.WithTx(tx)
	// get the current token information from the database
	tokenInfo, err := qtx.TokenByIDAndChainIDAndExternalID(internalCtx,
		queries.TokenByIDAndChainIDAndExternalIDParams{
			ID:         token.Address.Bytes(),
			ChainID:    token.ChainID,
			ExternalID: token.ExternalID,
		})
	if err != nil {
		return err
	}
	// if the token synced status is not the same that the received one, update
	// it in the database
	if tokenInfo.Synced != synced {
		_, err = qtx.UpdateTokenStatus(internalCtx, queries.UpdateTokenStatusParams{
			ID:         token.Address.Bytes(),
			ChainID:    token.ChainID,
			ExternalID: token.ExternalID,
			Synced:     synced,
		})
		if err != nil {
			return err
		}
		if synced {
			log.Infow("token synced",
				"token", token.Address.Hex(),
				"chainID", token.ChainID,
				"externalID", token.ExternalID)
		}
	}
	created, updated, deleted := 0, 0, 0
	for addr, balance := range holders {
		switch balance.Cmp(big.NewInt(0)) {
		case -1:
			// if the calculated balance is negative,try todelete the token holder
			if _, err := qtx.DeleteTokenHolder(ctx, queries.DeleteTokenHolderParams{
				TokenID:    token.Address.Bytes(),
				ChainID:    token.ChainID,
				ExternalID: token.ExternalID,
				HolderID:   addr.Bytes(),
			}); err != nil {
				return fmt.Errorf("error deleting token holder: %w", err)
			}
			deleted++
		case 1:
			// get the current token holder from the database
			currentTokenHolder, err := qtx.TokenHolderByTokenIDAndHolderIDAndChainIDAndExternalID(ctx,
				queries.TokenHolderByTokenIDAndHolderIDAndChainIDAndExternalIDParams{
					TokenID:    token.Address.Bytes(),
					ChainID:    token.ChainID,
					ExternalID: token.ExternalID,
					HolderID:   addr.Bytes(),
				})
			if err != nil {
				if !errors.Is(sql.ErrNoRows, err) {
					return err
				}
				// if the token holder not exists, create it
				_, err = qtx.CreateTokenHolder(ctx, queries.CreateTokenHolderParams{
					TokenID:    token.Address.Bytes(),
					ChainID:    token.ChainID,
					ExternalID: token.ExternalID,
					HolderID:   addr.Bytes(),
					BlockID:    lastBlock,
					Balance:    balance.String(),
				})
				if err != nil {
					return err
				}
				created++
				continue
			}
			// parse the current token holder balance and compare it with the
			// calculated one, if they are the same, continue
			currentBalance, ok := new(big.Int).SetString(currentTokenHolder.Balance, 10)
			if !ok {
				return fmt.Errorf("error parsing current token holder balance")
			}
			if currentBalance.Cmp(balance) == 0 {
				continue
			}
			// if the calculated balance is not 0 or less and it is different
			// from the current one, update it in the database
			_, err = qtx.UpdateTokenHolderBalance(ctx, queries.UpdateTokenHolderBalanceParams{
				TokenID:    token.Address.Bytes(),
				ChainID:    token.ChainID,
				ExternalID: token.ExternalID,
				HolderID:   addr.Bytes(),
				BlockID:    currentTokenHolder.BlockID,
				NewBlockID: lastBlock,
				Balance:    balance.String(),
			})
			if err != nil {
				return fmt.Errorf("error updating token holder: %w", err)
			}
			updated++
		}
	}
	// close the database tx and commit it
	if err := tx.Commit(); err != nil {
		return err
	}
	log.Debugw("token holders saved",
		"token", token.Address.Hex(),
		"chainID", token.ChainID,
		"externalID", token.ExternalID,
		"block", token.LastBlock,
		"synced", token.Synced,
		"created", created,
		"updated", updated,
		"deleted", deleted)
	return nil
}
