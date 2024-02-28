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
	"github.com/vocdoni/census3/db/annotations"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/scanner/providers"
	"github.com/vocdoni/census3/scanner/providers/web3"
	"go.vocdoni.io/dvote/log"
)

// ScannerToken includes the information of a token that the scanner needs to
// scan it.
type ScannerToken struct {
	Address     common.Address
	ChainID     uint64
	Type        uint64
	ExternalID  string
	LastBlock   uint64
	Ready       bool
	Synced      bool
	totalSupply *big.Int
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

	tokens                  []*ScannerToken
	tokensMtx               sync.Mutex
	waiter                  sync.WaitGroup
	latestBlockNumbers      sync.Map
	lastUpdatedBlockNumbers time.Time
}

// NewScanner returns a new scanner instance with the required parameters
// initialized.
func NewScanner(db *db.DB, networks web3.NetworkEndpoints, coolDown time.Duration) *Scanner {
	return &Scanner{
		db:                      db,
		networks:                networks,
		providers:               make(map[uint64]providers.HolderProvider),
		coolDown:                coolDown,
		tokens:                  []*ScannerToken{},
		tokensMtx:               sync.Mutex{},
		waiter:                  sync.WaitGroup{},
		latestBlockNumbers:      sync.Map{},
		lastUpdatedBlockNumbers: time.Time{},
	}
}

// SetProviders sets the providers that the scanner will use to get the holders
// of the tokens. It also creates the token types in the database if they do not
// exist. It returns an error something goes wrong creating the token types in
// the database.
func (s *Scanner) SetProviders(newProviders ...providers.HolderProvider) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// create a tx to use it in the following queries
	for _, provider := range newProviders {
		// try to create the token type in the database, if it already exists,
		// update its name to ensure that it is correct according to the type id
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
		// include the provider in the scanner
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
			// iterate over the tokens to scan
			atSyncGlobal := true
			for _, token := range tokens {
				log.Infow("scanning token",
					"address", token.Address.Hex(),
					"chainID", token.ChainID,
					"externalID", token.ExternalID,
					"lastBlock", token.LastBlock,
					"ready", token.Ready)
				// scan the token
				holders, newTransfers, lastBlock, synced, totalSupply, err := s.ScanHolders(ctx, token)
				if err != nil {
					log.Error(err)
					continue
				}
				if !synced {
					atSyncGlobal = false
				}
				// save the token holders in the database in a goroutine and
				// continue with the next token
				s.waiter.Add(1)
				go func(t *ScannerToken, h map[common.Address]*big.Int, n, lb uint64, sy bool, ts *big.Int) {
					defer s.waiter.Done()
					if err = s.SaveHolders(ctx, t, h, n, lb, sy, ts); err != nil {
						log.Error(err)
					}
				}(token, holders, newTransfers, lastBlock, synced, totalSupply)
			}
			log.Infow("scan iteration finished",
				"iteration", itCounter,
				"duration", time.Since(startTime).Seconds(),
				"atSync", atSyncGlobal)
			// if all the tokens are synced, sleep the cool down time, else,
			// sleep the scan sleep time
			if atSyncGlobal {
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
		tokens = append(tokens, &ScannerToken{
			Address:     common.BytesToAddress(token.ID),
			ChainID:     token.ChainID,
			Type:        token.TypeID,
			ExternalID:  token.ExternalID,
			LastBlock:   uint64(token.LastBlock),
			Ready:       token.CreationBlock > 0 && token.LastBlock >= token.CreationBlock,
			Synced:      token.Synced,
			totalSupply: totalSupply,
		})
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
			tokens = append(tokens, &ScannerToken{
				Address:     common.BytesToAddress(token.ID),
				ChainID:     token.ChainID,
				Type:        token.TypeID,
				ExternalID:  token.ExternalID,
				LastBlock:   uint64(token.LastBlock),
				Ready:       token.CreationBlock > 0 && token.LastBlock >= token.CreationBlock,
				Synced:      token.Synced,
				totalSupply: totalSupply,
			})
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
		tokens = append(tokens, &ScannerToken{
			Address:     common.BytesToAddress(token.ID),
			ChainID:     token.ChainID,
			Type:        token.TypeID,
			ExternalID:  token.ExternalID,
			LastBlock:   uint64(token.LastBlock),
			Ready:       token.CreationBlock > 0 && token.LastBlock >= token.CreationBlock,
			Synced:      token.Synced,
			totalSupply: totalSupply,
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
	map[common.Address]*big.Int, uint64, uint64, bool, *big.Int, error,
) {
	internalCtx, cancel := context.WithTimeout(ctx, SCAN_TIMEOUT)
	defer cancel()
	// get the correct token holder for the current token
	provider, exists := s.providers[token.Type]
	if !exists {
		return nil, 0, token.LastBlock, token.Synced, nil, fmt.Errorf("token type %d not supported", token.Type)
	}
	// create a tx to use it in the following queries
	tx, err := s.db.RW.BeginTx(internalCtx, nil)
	if err != nil {
		return nil, 0, token.LastBlock, token.Synced, nil, err
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
			return nil, 0, token.LastBlock, token.Synced, nil, err
		}
		// set the last block number of the network in the provider getting it
		// from the latest block numbers cache
		if iLastNetworkBlock, ok := s.latestBlockNumbers.Load(token.ChainID); ok {
			if lastNetworkBlock, ok := iLastNetworkBlock.(uint64); ok {
				provider.SetLastBlockNumber(lastNetworkBlock)
			}
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
				return nil, 0, token.LastBlock, token.Synced, nil, err
			}
			_, err = qtx.UpdateTokenBlocks(internalCtx, queries.UpdateTokenBlocksParams{
				ID:            token.Address.Bytes(),
				ChainID:       token.ChainID,
				ExternalID:    token.ExternalID,
				CreationBlock: int64(creationBlock),
				LastBlock:     int64(creationBlock),
			})
			if err != nil {
				return nil, 0, token.LastBlock, token.Synced, nil, err
			}
			token.LastBlock = creationBlock
		}
	}
	log.Infow("scanning holders",
		"address", token.Address.Hex(),
		"chainID", token.ChainID,
		"externalID", token.ExternalID,
		"lastBlock", token.LastBlock)
	// get the current token holders from the database
	results, err := qtx.ListTokenHolders(internalCtx,
		queries.ListTokenHoldersParams{
			TokenID:    token.Address.Bytes(),
			ChainID:    token.ChainID,
			ExternalID: token.ExternalID,
		})
	if err != nil {
		return nil, 0, token.LastBlock, token.Synced, nil, err
	}
	// set the current holders into the provider and get the new ones
	currentHolders := map[common.Address]*big.Int{}
	for _, result := range results {
		bBalance, ok := new(big.Int).SetString(result.Balance, 10)
		if !ok {
			return nil, 0, token.LastBlock, token.Synced, nil, fmt.Errorf("error parsing token holder balance")
		}
		currentHolders[common.BytesToAddress(result.HolderID)] = bBalance
	}
	// close the database tx and commit it
	if err := tx.Commit(); err != nil {
		return nil, 0, token.LastBlock, token.Synced, nil, err
	}
	// set the current holders into the provider and get the new ones
	if err := provider.SetLastBalances(ctx, []byte(token.ExternalID),
		currentHolders, token.LastBlock,
	); err != nil {
		return nil, 0, token.LastBlock, token.Synced, nil, err
	}
	// get the new holders from the provider
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
	synced bool, totalSupply *big.Int,
) error {
	log.Infow("saving token status and holders",
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
			if totalSupply != nil && totalSupply.Cmp(big.NewInt(0)) > 0 {
				s.tokens[i].totalSupply = totalSupply
				token.totalSupply = totalSupply
			}
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
			log.Errorf("error rolling back tx: %v, token=%s chainID=%d externalID=%s",
				err, token.Address.Hex(), token.ChainID, token.ExternalID)
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
		TotalSupply:       annotations.BigInt(token.totalSupply.String()),
	})
	if err != nil {
		return err
	}
	log.Debugw("token status saved",
		"synced", synced,
		"token", token.Address.Hex(),
		"chainID", token.ChainID,
		"externalID", token.ExternalID,
		"totalSupply", token.totalSupply.String(),
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
		currentTokenHolder, err := qtx.GetTokenHolderEvenZero(ctx, queries.GetTokenHolderEvenZeroParams{
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
		// if both balances are zero, continue with the next holder to prevent
		// UNIQUES constraint errors
		if balance.Cmp(big.NewInt(0)) == 0 && currentBalance.Cmp(big.NewInt(0)) == 0 {
			continue
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
	log.Debugw("committing token holders",
		"token", token.Address.Hex(),
		"chainID", token.ChainID,
		"externalID", token.ExternalID,
		"block", token.LastBlock,
		"synced", token.Synced,
		"created", created,
		"updated", updated)
	// close the database tx and commit it
	if err := tx.Commit(); err != nil {
		return err
	}
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
