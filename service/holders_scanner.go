package service

import (
	"bytes"
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
	_ "github.com/mattn/go-sqlite3"

	"github.com/vocdoni/census3/db"
	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/service/web3"
	"github.com/vocdoni/census3/state"
	"go.vocdoni.io/dvote/log"
)

var (
	ErrNoDB           = fmt.Errorf("no database instance provided")
	ErrHalted         = fmt.Errorf("scanner loop halted")
	ErrTokenNotExists = fmt.Errorf("token does not exists")
)

// HoldersScanner struct contains the needed parameters to scan the holders of
// the tokens stored on the database (located on 'dataDir/dbFilename'). It
// keeps the database updated scanning the network using the web3 endpoint.
type HoldersScanner struct {
	w3p          web3.NetworkEndpoints
	tokens       []*state.TokenHolders
	mutex        sync.RWMutex
	db           *db.DB
	lastBlock    uint64
	extProviders map[state.TokenType]HolderProvider
	coolDown     time.Duration
}

// NewHoldersScanner function creates a new HolderScanner using the dataDir path
// and the web3 endpoint URI provided. It sets up a sqlite3 database instance
// and gets the number of last block scanned from it.
func NewHoldersScanner(db *db.DB, w3p web3.NetworkEndpoints,
	ext map[state.TokenType]HolderProvider, coolDown time.Duration,
) (*HoldersScanner, error) {
	if db == nil {
		return nil, ErrNoDB
	}
	// create an empty scanner
	s := HoldersScanner{
		w3p:          w3p,
		tokens:       []*state.TokenHolders{},
		db:           db,
		extProviders: ext,
		coolDown:     coolDown,
	}
	// get latest analyzed block
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	lastBlock, err := s.db.QueriesRO.LastBlock(ctx)
	if err == nil {
		s.lastBlock = lastBlock
	}
	return &s, nil
}

// Start function initialises the given scanner until the provided context is
// canceled. It first gets the addresses of the tokens to scan and their current
// token state. It then starts scanning, keeping these lists updated and
// synchronised with the database instance.
func (s *HoldersScanner) Start(ctx context.Context) {
	// monitor for new tokens added and update every token holders
	itCounter := uint64(0)
	for {
		select {
		case <-ctx.Done():
			log.Info(ErrHalted)
			return
		default:
			itCounter++
			startTime := time.Now()
			// get updated list of tokens
			if err := s.getTokensToScan(); err != nil {
				log.Error(err)
				continue
			}
			// scan for new holders of every token
			atSyncGlobal := true
			for index, data := range s.tokens {
				if !data.IsReady() {
					if err := s.calcTokenCreationBlock(ctx, index); err != nil {
						log.Error(err)
						continue
					}
				}
				atSync, err := s.scanHolders(ctx, data.Address(), data.ChainID, []byte(data.ExternalID))
				if err != nil {
					log.Error(err)
					continue
				}
				if !atSync {
					atSyncGlobal = false
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

// getTokensToScan function gets the information of the current tokens to scan,
// including its addresses from the database. If the current database instance
// does not contain any token, it returns nil addresses without error.
// This behaviour helps to deal with this particular case. It also filters the
// tokens to retunr only the ones that are ready to be scanned, which means that
// the token creation block is already calculated.
func (s *HoldersScanner) getTokensToScan() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	tx, err := s.db.RO.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
			log.Errorf("error rolling back transaction when scanner get token addresses: %v", err)
		}
	}()
	qtx := s.db.QueriesRW.WithTx(tx)
	s.tokens = []*state.TokenHolders{}
	// get last created tokens from the database to scan them first
	lastNotSyncedTokens, err := qtx.ListLastNoSyncedTokens(ctx)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return err
	}
	// parse last not synced token addresses
	for _, token := range lastNotSyncedTokens {
		lastBlock := uint64(token.CreationBlock)
		if blockNumber, err := s.db.QueriesRO.LastBlockByTokenID(ctx, token.ID); err == nil {
			lastBlock = blockNumber
		}
		s.tokens = append(s.tokens, new(state.TokenHolders).Init(
			common.BytesToAddress(token.ID), state.TokenType(token.TypeID),
			lastBlock, token.ChainID, token.ExternalID))
	}
	// get old tokens from the database
	oldNotSyncedTokens, err := qtx.ListOldNoSyncedTokens(ctx)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return err
	}
	// get the current block number of every chain
	currentBlockNumbers, err := s.w3p.CurrentBlockNumbers(ctx)
	if err != nil {
		return err
	}
	// sort old not synced tokens by nearest to be synced, that is, the tokens
	// that have the minimum difference between the current block of its chain
	// and the last block scanned by the scanner (retrieved from the database
	// as LastBlock)
	sort.Slice(oldNotSyncedTokens, func(i, j int) bool {
		iLastBlock := uint64(0)
		if oldNotSyncedTokens[i].LastBlock != nil {
			iLastBlock = uint64(oldNotSyncedTokens[i].LastBlock.(int64))
		}
		jLastBlock := uint64(0)
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
		if blockNumber, err := s.db.QueriesRO.LastBlockByTokenID(ctx, token.ID); err == nil {
			lastBlock = blockNumber
		}
		s.tokens = append(s.tokens, new(state.TokenHolders).Init(
			common.BytesToAddress(token.ID), state.TokenType(token.TypeID),
			lastBlock, token.ChainID, token.ExternalID))
	}
	// get last created tokens from the database to scan them first
	syncedTokens, err := qtx.ListSyncedTokens(ctx)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return err
	}
	for _, token := range syncedTokens {
		lastBlock := uint64(token.CreationBlock)
		if blockNumber, err := s.db.QueriesRO.LastBlockByTokenID(ctx, token.ID); err == nil {
			lastBlock = blockNumber
		}
		s.tokens = append(s.tokens, new(state.TokenHolders).Init(
			common.BytesToAddress(token.ID), state.TokenType(token.TypeID),
			lastBlock, token.ChainID, token.ExternalID))
	}
	return nil
}

// saveHolders function updates the current HoldersScanner database with the
// TokenHolders state provided. Updates the holders for associated token and
// the blocks scanned. To do this, it requires the root hash and the timestampt
// of the given TokenHolders state block.
func (s *HoldersScanner) saveHolders(th *state.TokenHolders) error {
	log.Debugw("saving token holders",
		"token", th.Address(),
		"chainID", th.ChainID,
		"externalID", th.ExternalID,
		"block", th.LastBlock(),
		"holders", len(th.Holders()))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// begin a transaction for group sql queries
	tx, err := s.db.RW.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()
	qtx := s.db.QueriesRW.WithTx(tx)
	if exists, err := qtx.ExistsTokenByChainIDAndExternalID(ctx, queries.ExistsTokenByChainIDAndExternalIDParams{
		ID:         th.Address().Bytes(),
		ChainID:    th.ChainID,
		ExternalID: th.ExternalID,
	}); err != nil {
		return fmt.Errorf("error checking if token exists: %w", err)
	} else if !exists {
		return ErrTokenNotExists
	}
	_, err = qtx.UpdateTokenStatus(ctx, queries.UpdateTokenStatusParams{
		Synced:     th.IsSynced(),
		ID:         th.Address().Bytes(),
		ChainID:    th.ChainID,
		ExternalID: th.ExternalID,
	})
	if err != nil {
		return fmt.Errorf("error updating token: %w", err)
	}
	// if not token holders received, skip
	if len(th.Holders()) == 0 {
		log.Debug("no holders to save. skip scanning and saving...")
		// save btw to update if token is synced
		if err := tx.Commit(); err != nil {
			return err
		}
		return nil
	}
	// get the block timestamp and root hash from the holder provider, it does
	// not matter if it is an external provider or a web3 one
	var timestamp string
	var rootHash []byte
	if provider, isExternal := s.extProviders[th.Type()]; isExternal {
		timestamp, err = provider.BlockTimestamp(ctx, th.LastBlock())
		if err != nil {
			return err
		}
		rootHash, err = provider.BlockRootHash(ctx, th.LastBlock())
		if err != nil {
			return err
		}
	} else {
		// get correct web3 uri provider
		w3Endpoint, exists := s.w3p.EndpointByChainID(th.ChainID)
		if !exists {
			return fmt.Errorf("chain ID not supported")
		}
		// init web3 contract state
		w3 := state.Web3{}
		if err := w3.Init(ctx, w3Endpoint, th.Address(), th.Type()); err != nil {
			return err
		}
		// get current block number timestamp and root hash, required parameters to
		// create a new block in the database
		timestamp, err = w3.BlockTimestamp(ctx, uint(th.LastBlock()))
		if err != nil {
			return err
		}
		rootHash, err = w3.BlockRootHash(ctx, uint(th.LastBlock()))
		if err != nil {
			return err
		}
	}
	// if the current HoldersScanner last block not exists in the database,
	// create it
	if _, err := qtx.BlockByID(ctx, th.LastBlock()); err != nil {
		if !errors.Is(sql.ErrNoRows, err) {
			return err
		}
		_, err = qtx.CreateBlock(ctx, queries.CreateBlockParams{
			ID:        th.LastBlock(),
			Timestamp: timestamp,
			RootHash:  rootHash,
		})
		if err != nil {
			return err
		}
	}
	// iterate over given holders
	//  - if the holder not exists, create it
	//  - if the holder already exists, calculate the new balance with the
	//    current balance
	//		- if the calculated balance is 0 delete it
	//		- if the calculated balance is not 0, update it
	created, updated, deleted := 0, 0, 0
	for holder, balance := range th.Holders() {
		currentTokenHolder, err := qtx.TokenHolderByTokenIDAndHolderIDAndChainIDAndExternalID(ctx,
			queries.TokenHolderByTokenIDAndHolderIDAndChainIDAndExternalIDParams{
				TokenID:    th.Address().Bytes(),
				HolderID:   holder.Bytes(),
				ChainID:    th.ChainID,
				ExternalID: th.ExternalID,
			})
		if err != nil {
			// return the error if fails and the error is not 'no rows' err
			if !errors.Is(sql.ErrNoRows, err) {
				return err
			}
			// if the token holder not exists and the balance is 0 or less, skip
			// it
			if balance.Cmp(big.NewInt(0)) != 1 {
				continue
			}
			_, err = qtx.CreateHolder(ctx, holder.Bytes())
			if err != nil && !strings.Contains(err.Error(), "UNIQUE constraint failed") {
				return err
			}
			// if the token holder not exists, create it
			_, err = qtx.CreateTokenHolder(ctx, queries.CreateTokenHolderParams{
				TokenID:    th.Address().Bytes(),
				HolderID:   holder.Bytes(),
				BlockID:    th.LastBlock(),
				Balance:    balance.String(),
				ChainID:    th.ChainID,
				ExternalID: th.ExternalID,
			})
			if err != nil {
				return err
			}
			created++
			continue
		}
		// if the calculated balance is 0 or less delete it
		if balance.Cmp(big.NewInt(0)) != 1 {
			if _, err := qtx.DeleteTokenHolder(ctx, queries.DeleteTokenHolderParams{
				TokenID:    th.Address().Bytes(),
				HolderID:   holder.Bytes(),
				ChainID:    th.ChainID,
				ExternalID: th.ExternalID,
			}); err != nil {
				return fmt.Errorf("error deleting token holder: %w", err)
			}
			deleted++
			continue
		}
		// if the calculated balance is the same as the current balance, skip it
		currentBalance, ok := new(big.Int).SetString(currentTokenHolder.Balance, 10)
		if !ok {
			return fmt.Errorf("error parsing current balance: %w", err)
		}
		if currentBalance.Cmp(balance) == 0 {
			continue
		}
		// if the calculated balance is not 0 or less, update it
		_, err = qtx.UpdateTokenHolderBalance(ctx, queries.UpdateTokenHolderBalanceParams{
			TokenID:    th.Address().Bytes(),
			HolderID:   holder.Bytes(),
			BlockID:    currentTokenHolder.BlockID,
			NewBlockID: th.LastBlock(),
			Balance:    balance.String(),
			ChainID:    th.ChainID,
			ExternalID: th.ExternalID,
		})
		if err != nil {
			return fmt.Errorf("error updating token holder: %w", err)
		}
		updated++
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	log.Debugw("token holders saved",
		"token", th.Address(),
		"chainID", th.ChainID,
		"externalID", th.ExternalID,
		"block", th.LastBlock(),
		"created", created,
		"updated", updated,
		"deleted", deleted)
	th.FlushHolders()
	return nil
}

// cachedToken function returns the TokenHolders struct associated to the
// address, chainID and externalID provided. If it does not exists, it creates
// a new one and caches it getting the token information from the database.
func (s *HoldersScanner) cachedToken(ctx context.Context, addr common.Address,
	chainID uint64, externalID []byte,
) (*state.TokenHolders, error) {
	// get the token TokenHolders struct from cache, if it not exists it will
	// be initialized
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var th *state.TokenHolders
	for index, token := range s.tokens {
		if bytes.Equal(token.Address().Bytes(), addr.Bytes()) &&
			token.ChainID == chainID && bytes.Equal([]byte(token.ExternalID), externalID) {
			tokenHolders, err := s.db.QueriesRO.TokenHoldersByTokenIDAndChainIDAndExternalID(ctx,
				queries.TokenHoldersByTokenIDAndChainIDAndExternalIDParams{
					TokenID:    addr.Bytes(),
					ChainID:    chainID,
					ExternalID: string(externalID),
				})
			if err != nil {
				return nil, err
			}
			token.FlushHolders()
			for _, holder := range tokenHolders {
				balance, ok := new(big.Int).SetString(holder.Balance, 10)
				if !ok {
					return nil, fmt.Errorf("error parsing balance: %w", err)
				}
				token.Append(common.BytesToAddress(holder.ID), balance)
			}
			s.tokens[index] = token
			return token, nil
		}
	}
	log.Infow("initializing token", "address", addr, "chainID", chainID, "externalID", externalID)
	// get token information from the database
	tokenInfo, err := s.db.QueriesRO.TokenByIDAndChainIDAndExternalID(ctx, queries.TokenByIDAndChainIDAndExternalIDParams{
		ID:         addr.Bytes(),
		ChainID:    chainID,
		ExternalID: string(externalID),
	})
	if err != nil {
		return nil, err
	}
	ttype := state.TokenType(tokenInfo.TypeID)
	tokenLastBlock := uint64(tokenInfo.CreationBlock)
	if blockNumber, err := s.db.QueriesRO.LastBlockByTokenID(ctx, addr.Bytes()); err == nil {
		tokenLastBlock = blockNumber
	}
	th = new(state.TokenHolders).Init(addr, ttype, tokenLastBlock, tokenInfo.ChainID, tokenInfo.ExternalID)
	tokenHolders, err := s.db.QueriesRO.TokenHoldersByTokenIDAndChainIDAndExternalID(ctx,
		queries.TokenHoldersByTokenIDAndChainIDAndExternalIDParams{
			TokenID:    addr.Bytes(),
			ChainID:    chainID,
			ExternalID: string(externalID),
		})
	if err != nil {
		return nil, err
	}
	for _, holder := range tokenHolders {
		balance, ok := new(big.Int).SetString(holder.Balance, 10)
		if !ok {
			return nil, fmt.Errorf("error parsing balance: %w", err)
		}
		th.Append(common.BytesToAddress(holder.ID), balance)
	}
	s.tokens = append(s.tokens, th)
	return th, nil
}

// scanHolders function updates the holders of the token identified by the
// address provided. It checks if the address provided already has a
// TokenHolders state cached, if not, it gets the token information from the
// HoldersScanner database and caches it. If something expected fails or the
// scan process ends successfully, the cached information is stored in the
// database. If it has no updates, it does not change anything and returns nil.
func (s *HoldersScanner) scanHolders(ctx context.Context, addr common.Address, chainID uint64, externalID []byte) (bool, error) {
	log.Infow("scanning holders", "address", addr, "chainID", chainID, "externalID", externalID)
	ctx, cancel := context.WithTimeout(ctx, scanIterationDurationPerToken)
	defer cancel()
	th, err := s.cachedToken(ctx, addr, chainID, externalID)
	if err != nil {
		return false, err
	}
	// If the last block of the current scanner is lower than the TokenHolders
	// state block, it seems that the current scanner is out of date and can
	// move on to this block
	if s.lastBlock < th.LastBlock() {
		s.lastBlock = th.LastBlock()
	}
	// if the token type has a external provider associated, get the holders
	// from it and append it to the TokenHolders struct, then save it into the
	// database and return
	if provider, isExternal := s.extProviders[th.Type()]; isExternal {
		if err := provider.SetLastBalances(ctx, []byte(th.ExternalID), th.Holders(), th.LastBlock()); err != nil {
			return false, err
		}
		externalBalances, err := provider.HoldersBalances(ctx, []byte(th.ExternalID), s.lastBlock-th.LastBlock())
		if err != nil {
			return false, err
		}
		for holder, balance := range externalBalances {
			th.Append(holder, balance)
		}
		blockNumber, err := provider.LatestBlockNumber(ctx, []byte(th.ExternalID))
		if err != nil {
			return false, err
		}
		th.BlockDone(blockNumber)
		th.Synced()
		return true, s.saveHolders(th)
	}
	// get correct web3 uri provider
	w3URI, exists := s.w3p.EndpointByChainID(th.ChainID)
	if !exists {
		return false, fmt.Errorf("chain ID not supported")
	}
	// init web3 contract state
	w3 := state.Web3{}
	if err := w3.Init(ctx, w3URI, addr, th.Type()); err != nil {
		return th.IsSynced(), err
	}
	// try to update the TokenHolders struct and the current scanner last block
	if _, err := w3.UpdateTokenHolders(ctx, th); err != nil {
		if strings.Contains(err.Error(), "no new blocks") {
			// if no new blocks error raises, log it as debug and return nil
			log.Debugw("no new blocks to scan", "token", th.Address())
			return true, nil
		}
		if strings.Contains(err.Error(), "connection reset") ||
			strings.Contains(err.Error(), "context deadline") ||
			strings.Contains(err.Error(), "read limit exceeded") ||
			strings.Contains(err.Error(), "limit reached") {
			// if connection error raises, log it as warning and try to save
			// current TokenHolders state and return nil
			log.Warnw("warning scanning contract", "token", th.Address().Hex(),
				"block", th.LastBlock(), "error", err)
			// save TokesHolders state into the database before exit of the function
			return th.IsSynced(), s.saveHolders(th)
		}
		// if unexpected error raises, log it as error and return it.
		log.Errorw(fmt.Errorf("warning scanning contract: %v", err),
			fmt.Sprintf("token=%s block%d", th.Address().Hex(), th.LastBlock()))
		return th.IsSynced(), err
	}
	// save TokesHolders state into the database before exit of the function
	return th.IsSynced(), s.saveHolders(th)
}

// calcTokenCreationBlock function attempts to calculate the block number when
// the token contract provided was created and deployed and updates the database
// with the result obtained.
func (s *HoldersScanner) calcTokenCreationBlock(ctx context.Context, index int) error {
	if len(s.tokens) < index {
		return fmt.Errorf("token not found")
	}
	addr := s.tokens[index].Address()
	chainID := s.tokens[index].ChainID
	// set a deadline of 10 seconds from the current context
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	// get the token type
	tokenInfo, err := s.db.QueriesRO.TokenByIDAndChainID(ctx,
		queries.TokenByIDAndChainIDParams{
			ID:      addr.Bytes(),
			ChainID: chainID,
		})
	if err != nil {
		return fmt.Errorf("error getting token from database: %w", err)
	}
	creationBlock := uint64(0)
	ttype := state.TokenType(tokenInfo.TypeID)
	if provider, isExternal := s.extProviders[ttype]; isExternal {
		creationBlock, err = provider.CreationBlock(ctx, []byte(tokenInfo.ExternalID))
		if err != nil {
			return fmt.Errorf("error getting token creation block: %w", err)
		}
	} else {
		// get correct web3 uri provider
		w3URI, exists := s.w3p.EndpointByChainID(tokenInfo.ChainID)
		if !exists {
			return fmt.Errorf("chain ID not supported")
		}
		// init web3 contract state
		w3 := state.Web3{}
		if err := w3.Init(ctx, w3URI, addr, ttype); err != nil {
			return fmt.Errorf("error intializing web3 client for this token: %w", err)
		}
		// get creation block of the current token contract
		creationBlock, err = w3.ContractCreationBlock(ctx)
		if err != nil {
			return fmt.Errorf("error getting token creation block: %w", err)
		}
	}
	// save the creation block into the database
	_, err = s.db.QueriesRW.UpdateTokenCreationBlock(ctx,
		queries.UpdateTokenCreationBlockParams{
			ID:            addr.Bytes(),
			CreationBlock: int64(creationBlock),
		})
	if err != nil {
		return fmt.Errorf("error updating token creation block on the database: %w", err)
	}
	s.tokens[index].BlockDone(creationBlock)
	return err
}
