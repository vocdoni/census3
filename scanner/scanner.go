package scanner

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/db"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/internal"
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
func NewScanner(db *db.DB, networks web3.NetworkEndpoints, coolDown time.Duration) *Scanner {
	return &Scanner{
		db:        db,
		networks:  networks,
		providers: make(map[uint64]providers.HolderProvider),
		coolDown:  coolDown,
		tokens:    []*ScannerToken{},
		tokensMtx: sync.Mutex{},
		waiter:    sync.WaitGroup{},
	}
}

// SetProviders sets the providers that the scanner will use to get the holders
// of the tokens.
func (s *Scanner) SetProviders(newProviders ...providers.HolderProvider) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// create a tx to use it in the following queries
	for _, provider := range newProviders {
		if _, err := s.db.QueriesRW.CreateTokenType(ctx, queries.CreateTokenTypeParams{
			ID:       provider.Type(),
			TypeName: provider.TypeName(),
		}); err != nil {
			if !strings.Contains(err.Error(), "UNIQUE constraint failed") {
				return err
			}
			if _, err := s.db.QueriesRW.UpdateTokenType(ctx, queries.UpdateTokenTypeParams{
				ID:       provider.Type(),
				TypeName: provider.TypeName(),
			}); err != nil {
				return err
			}
		}
		s.providers[provider.Type()] = provider
	}
	return nil
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
				log.Infow("scanning token",
					"address", token.Address.Hex(),
					"chainID", token.ChainID,
					"externalID", token.ExternalID,
					"lastBlock", token.LastBlock,
					"ready", token.Ready)
				holders, newTransfers, lastBlock, synced, err := s.ScanHolders(ctx, token)
				if err != nil {
					log.Error(err)
					continue
				}
				if !synced {
					atSyncGlobal = false
				}
				s.waiter.Add(1)
				go func(t *ScannerToken, h map[common.Address]*big.Int, n, lb uint64, sy bool) {
					defer s.waiter.Done()
					if err = s.SaveHolders(ctx, t, h, n, lb, sy); err != nil {
						log.Error(err)
					}
				}(token, holders, newTransfers, lastBlock, synced)
			}
			log.Infow("scan iteration finished",
				"iteration", itCounter,
				"duration", time.Since(startTime).Seconds(),
				"atSync", atSyncGlobal)
			log.Debugf("GetBlockByNumberCounter: %d", internal.GetBlockByNumberCounter.Load())
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
	for _, provider := range s.providers {
		if err := provider.Close(); err != nil {
			log.Error(err)
		}
	}
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
	tokens := []*ScannerToken{}
	// get last created tokens from the database to scan them first
	lastNotSyncedTokens, err := s.db.QueriesRO.ListLastNoSyncedTokens(internalCtx)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return nil, err
	}
	// parse last not synced token addresses
	for _, token := range lastNotSyncedTokens {
		tokens = append(tokens, &ScannerToken{
			Address:    common.BytesToAddress(token.ID),
			ChainID:    token.ChainID,
			Type:       token.TypeID,
			ExternalID: token.ExternalID,
			LastBlock:  uint64(token.LastBlock),
			Ready:      token.CreationBlock > 0 && token.LastBlock >= token.CreationBlock,
			Synced:     token.Synced,
		})
	}
	// get old tokens from the database
	oldNotSyncedTokens, err := s.db.QueriesRO.ListOldNoSyncedTokens(internalCtx)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return nil, err
	}
	// if there are old not synced tokens, sort them by nearest to be synced
	// and parse them, if not, continue to avoid web3 calls
	if len(oldNotSyncedTokens) > 0 {
		// get the current block number of every chain
		currentBlockNumbers, err := s.networks.CurrentBlockNumbers(internalCtx)
		if err != nil {
			return nil, err
		}
		// sort old not synced tokens by nearest to be synced, that is, the tokens
		// that have the minimum difference between the current block of its chain
		// and the last block scanned by the scanner (retrieved from the database
		// as LastBlock)
		sort.Slice(oldNotSyncedTokens, func(i, j int) bool {
			iBlocksReamining := currentBlockNumbers[oldNotSyncedTokens[i].ChainID] -
				uint64(oldNotSyncedTokens[i].LastBlock)
			jBlocksReamining := currentBlockNumbers[oldNotSyncedTokens[j].ChainID] -
				uint64(oldNotSyncedTokens[j].LastBlock)
			return iBlocksReamining < jBlocksReamining
		})
		// parse old not synced token addresses
		for _, token := range oldNotSyncedTokens {
			tokens = append(tokens, &ScannerToken{
				Address:    common.BytesToAddress(token.ID),
				ChainID:    token.ChainID,
				Type:       token.TypeID,
				ExternalID: token.ExternalID,
				LastBlock:  uint64(token.LastBlock),
				Ready:      token.CreationBlock > 0 && token.LastBlock >= token.CreationBlock,
				Synced:     token.Synced,
			})
		}
	}
	// get last created tokens from the database to scan them first
	syncedTokens, err := s.db.QueriesRO.ListSyncedTokens(internalCtx)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return nil, err
	}
	for _, token := range syncedTokens {
		tokens = append(tokens, &ScannerToken{
			Address:    common.BytesToAddress(token.ID),
			ChainID:    token.ChainID,
			Type:       token.TypeID,
			ExternalID: token.ExternalID,
			LastBlock:  uint64(token.LastBlock),
			Ready:      token.CreationBlock > 0 && token.LastBlock >= token.CreationBlock,
			Synced:     token.Synced,
		})
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
func (s *Scanner) ScanHolders(ctx context.Context, token *ScannerToken) (
	map[common.Address]*big.Int, uint64, uint64, bool, error,
) {
	log.Infow("scanning holders",
		"address", token.Address.Hex(),
		"chainID", token.ChainID,
		"externalID", token.ExternalID,
		"lastBlock", token.LastBlock)
	internalCtx, cancel := context.WithTimeout(ctx, SCAN_TIMEOUT)
	defer cancel()
	// get the correct token holder for the current token
	provider, exists := s.providers[token.Type]
	if !exists {
		return nil, 0, token.LastBlock, token.Synced, fmt.Errorf("token type %d not supported", token.Type)
	}
	// create a tx to use it in the following queries
	tx, err := s.db.RW.BeginTx(internalCtx, nil)
	if err != nil {
		return nil, 0, token.LastBlock, token.Synced, err
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
			return nil, 0, token.LastBlock, token.Synced, err
		}
		// if the token is not ready yet (its creation block has not been
		// calculated yet), calculate it, update the token information and
		// return
		if !token.Ready {
			log.Debugw("token not ready yet, calculating creation block and continue",
				"address", token.Address.Hex(),
				"chainID", token.ChainID,
				"externalID", token.ExternalID)
			creationBlock, err := provider.CreationBlock(internalCtx, []byte(token.ExternalID))
			if err != nil {
				return nil, 0, token.LastBlock, token.Synced, err
			}
			_, err = qtx.UpdateTokenBlocks(internalCtx, queries.UpdateTokenBlocksParams{
				ID:            token.Address.Bytes(),
				ChainID:       token.ChainID,
				ExternalID:    token.ExternalID,
				CreationBlock: int64(creationBlock),
				LastBlock:     int64(creationBlock),
			})
			if err != nil {
				return nil, 0, token.LastBlock, token.Synced, err
			}
			// close the database tx and commit it
			return nil, 0, creationBlock, token.Synced, tx.Commit()
		}
	}
	// get the current token holders from the database
	results, err := qtx.ListTokenHolders(internalCtx,
		queries.ListTokenHoldersParams{
			TokenID:    token.Address.Bytes(),
			ChainID:    token.ChainID,
			ExternalID: token.ExternalID,
		})
	if err != nil {
		return nil, 0, token.LastBlock, token.Synced, err
	}
	// set the current holders into the provider and get the new ones
	currentHolders := map[common.Address]*big.Int{}
	for _, result := range results {
		bBalance, ok := new(big.Int).SetString(result.Balance, 10)
		if !ok {
			return nil, 0, token.LastBlock, token.Synced, fmt.Errorf("error parsing token holder balance")
		}
		currentHolders[common.BytesToAddress(result.HolderID)] = bBalance
	}
	// close the database tx and commit it
	if err := tx.Commit(); err != nil {
		return nil, 0, token.LastBlock, token.Synced, err
	}
	// set the current holders into the provider and get the new ones
	if err := provider.SetLastBalances(ctx, []byte(token.ExternalID),
		currentHolders, token.LastBlock,
	); err != nil {
		return nil, 0, token.LastBlock, token.Synced, err
	}
	return provider.HoldersBalances(ctx, []byte(token.ExternalID), token.LastBlock)
}

// SaveHolders saves the given holders in the database. It updates the token
// synced status if it is different from the received one. Then, it creates,
// updates or deletes the token holders in the database depending on the
// calculated balance.
// WARNING: the following code could produce holders with negative balances
// in the database. This is because the scanner does not know if the token
// holder is a contract or not, so it does not know if the balance is
// correct or not. The scanner assumes that the balance is correct and
// updates it in the database:
//  1. To get the correct holders from the database you must filter the
//     holders with negative balances.
//  2. To get the correct balances you must use the contract methods to get
//     the balances of the holders.
func (s *Scanner) SaveHolders(ctx context.Context, token *ScannerToken,
	holders map[common.Address]*big.Int, newTransfers, lastBlock uint64,
	synced bool,
) error {
	log.Debugw("saving token status and holders",
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
	tokenInfo, err := qtx.GetToken(internalCtx,
		queries.GetTokenParams{
			ID:         token.Address.Bytes(),
			ChainID:    token.ChainID,
			ExternalID: token.ExternalID,
		})
	if err != nil {
		return err
	}
	// update the balance synced status and last block in the database
	_, err = qtx.UpdateTokenStatus(internalCtx, queries.UpdateTokenStatusParams{
		ID:                token.Address.Bytes(),
		ChainID:           token.ChainID,
		ExternalID:        token.ExternalID,
		Synced:            synced,
		LastBlock:         int64(lastBlock),
		AnalysedTransfers: tokenInfo.AnalysedTransfers + int64(newTransfers),
	})
	if err != nil {
		return err
	}
	log.Debugw("token status saved",
		"synced", synced,
		"token", token.Address.Hex(),
		"chainID", token.ChainID,
		"externalID", token.ExternalID,
		"block", lastBlock)
	if len(holders) == 0 {
		log.Debugw("no holders to save, skipping...",
			"token", token.Address.Hex(),
			"chainID", token.ChainID,
			"externalID", token.ExternalID)
		return tx.Commit()
	}
	// create, update or delete token holders
	created, updated := 0, 0
	for addr, balance := range holders {
		// get the current token holder from the database
		currentTokenHolder, err := qtx.GetTokenHolder(ctx, queries.GetTokenHolderParams{
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
		// parse the current balance of the holder
		currentBalance, ok := new(big.Int).SetString(currentTokenHolder.Balance, 10)
		if !ok {
			return fmt.Errorf("error parsing current token holder balance")
		}
		// calculate the new balance of the holder by adding the current balance
		// and the new balance
		newBalance := new(big.Int).Add(currentBalance, balance)
		// update the token holder in the database with the new balance.
		// WANING: the balance could be negative so you must filter the holders
		// with negative balances to get the correct holders from the database.
		_, err = qtx.UpdateTokenHolderBalance(ctx, queries.UpdateTokenHolderBalanceParams{
			TokenID:    token.Address.Bytes(),
			ChainID:    token.ChainID,
			ExternalID: token.ExternalID,
			HolderID:   addr.Bytes(),
			BlockID:    currentTokenHolder.BlockID,
			NewBlockID: lastBlock,
			Balance:    newBalance.String(),
		})
		if err != nil {
			return fmt.Errorf("error updating token holder: %w", err)
		}
		updated++
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
		"updated", updated)
	return nil
}
