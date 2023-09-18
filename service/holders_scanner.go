package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	_ "github.com/mattn/go-sqlite3"

	"github.com/vocdoni/census3/db"
	queries "github.com/vocdoni/census3/db/sqlc"
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
	w3p       map[uint64]string
	tokens    map[common.Address]*state.TokenHolders
	mutex     sync.RWMutex
	db        *db.DB
	lastBlock uint64
}

// NewHoldersScanner function creates a new HolderScanner using the dataDir path
// and the web3 endpoint URI provided. It sets up a sqlite3 database instance
// and gets the number of last block scanned from it.
func NewHoldersScanner(db *db.DB, w3p map[uint64]string) (*HoldersScanner, error) {
	if db == nil {
		return nil, ErrNoDB
	}
	// create an empty scanner
	s := HoldersScanner{
		w3p:    w3p,
		tokens: make(map[common.Address]*state.TokenHolders),
		db:     db,
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
			tokens, err := s.tokenAddresses()
			if err != nil {
				log.Error(err)
				continue
			}
			// scan for new holders of every token
			atSyncGlobal := true
			for addr, ready := range tokens {
				if !ready {
					if err := s.calcTokenCreationBlock(ctx, addr); err != nil {
						log.Error(err)
						continue
					}
				}
				atSync, err := s.scanHolders(ctx, addr)
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
				time.Sleep(scanSleepTimeOnceSync)
			} else {
				time.Sleep(scanSleepTime)
			}
		}
	}
}

// tokenAddresses function gets the current token addresses from the database
// and returns it as a list of common.Address structs. If the current database
// instance does not contain any token, it returns nil addresses without error.
// This behaviour helps to deal with this particular case.
func (s *HoldersScanner) tokenAddresses() (map[common.Address]bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// get tokens from the database
	tokens, err := s.db.QueriesRO.ListTokens(ctx)
	// if error raises and is no rows error return nil results, if it is not
	// return the error.
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, nil
		}
		return nil, err
	}
	// parse and return token addresses
	results := make(map[common.Address]bool)
	for _, token := range tokens {
		results[common.BytesToAddress(token.ID)] = token.CreationBlock.Valid
	}
	return results, nil
}

// saveHolders function updates the current HoldersScanner database with the
// TokenHolders state provided. Updates the holders for associated token and
// the blocks scanned. To do this, it requires the root hash and the timestampt
// of the given TokenHolders state block.
func (s *HoldersScanner) saveHolders(th *state.TokenHolders) error {
	log.Debugw("saving token holders",
		"token", th.Address(),
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
	if exists, err := qtx.ExistsToken(ctx, th.Address().Bytes()); err != nil {
		return fmt.Errorf("error checking if token exists: %w", err)
	} else if !exists {
		return ErrTokenNotExists
	}
	_, err = qtx.UpdateTokenStatus(ctx, queries.UpdateTokenStatusParams{
		Synced: th.IsSynced(),
		ID:     th.Address().Bytes(),
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
	// get correct web3 uri provider
	w3uri, exists := s.w3p[th.ChainID]
	if !exists {
		return fmt.Errorf("chain ID not supported")
	}
	// init web3 contract state
	w3 := state.Web3{}
	if err := w3.Init(ctx, w3uri, th.Address(), th.Type()); err != nil {
		return err
	}
	// get current block number timestamp and root hash, required parameters to
	// create a new block in the database
	timestamp, err := w3.BlockTimestamp(ctx, uint(th.LastBlock()))
	if err != nil {
		return err
	}
	rootHash, err := w3.BlockRootHash(ctx, uint(th.LastBlock()))
	if err != nil {
		return err
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
		currentTokenHolder, err := qtx.TokenHolderByTokenIDAndHolderID(ctx,
			queries.TokenHolderByTokenIDAndHolderIDParams{
				TokenID:  th.Address().Bytes(),
				HolderID: holder.Bytes(),
			})
		if err != nil {
			// return the error if fails and the error is not 'no rows' err
			if !errors.Is(sql.ErrNoRows, err) {
				return err
			}
			_, err = qtx.CreateHolder(ctx, holder.Bytes())
			if err != nil && !strings.Contains(err.Error(), "UNIQUE constraint failed") {
				return err
			}
			// if the token holder not exists, create it
			_, err = qtx.CreateTokenHolder(ctx, queries.CreateTokenHolderParams{
				TokenID:  th.Address().Bytes(),
				HolderID: holder.Bytes(),
				BlockID:  th.LastBlock(),
				Balance:  balance.Bytes(),
			})
			if err != nil {
				return err
			}
			created++
			continue
		}
		// if the holder already exists, calculate the holder balance with the
		// current balance and the new one
		currentBalance := new(big.Int).SetBytes(currentTokenHolder.Balance)
		newBalance := new(big.Int).Add(currentBalance, balance)
		// if the calculated balance is 0 delete it
		if newBalance.Cmp(big.NewInt(0)) <= 0 {
			_, err := qtx.DeleteTokenHolder(ctx,
				queries.DeleteTokenHolderParams{
					TokenID:  th.Address().Bytes(),
					HolderID: holder.Bytes(),
				})
			if err != nil {
				return fmt.Errorf("error deleting token holder: %w", err)
			}
			deleted++
			continue
		}
		// if the calculated balance is not 0, update it
		_, err = qtx.UpdateTokenHolderBalance(ctx, queries.UpdateTokenHolderBalanceParams{
			TokenID:    th.Address().Bytes(),
			HolderID:   holder.Bytes(),
			BlockID:    currentTokenHolder.BlockID,
			NewBlockID: th.LastBlock(),
			Balance:    newBalance.Bytes(),
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
		"block", th.LastBlock(),
		"created", created,
		"updated", updated,
		"deleted", deleted)
	th.FlushHolders()
	return nil
}

// scanHolders function updates the holders of the token identified by the
// address provided. It checks if the address provided already has a
// TokenHolders state cached, if not, it gets the token information from the
// HoldersScanner database and caches it. If something expected fails or the
// scan process ends successfully, the cached information is stored in the
// database. If it has no updates, it does not change anything and returns nil.
func (s *HoldersScanner) scanHolders(ctx context.Context, addr common.Address) (bool, error) {
	log.Debugf("scanning contract %s", addr)
	ctx, cancel := context.WithTimeout(ctx, scanIterationDurationPerToken)
	defer cancel()
	// get the token TokenHolders struct from cache, if it not exists it will
	// be initialized
	s.mutex.RLock()
	th, ok := s.tokens[addr]
	if !ok {
		log.Infof("initializing contract %s", addr.Hex())
		// get token information from the database
		tokenInfo, err := s.db.QueriesRO.TokenByID(ctx, addr.Bytes())
		if err != nil {
			return false, err
		}
		ttype := state.TokenType(tokenInfo.TypeID)
		tokenLastBlock := uint64(tokenInfo.CreationBlock.Int64)
		if blockNumber, err := s.db.QueriesRO.LastBlockByTokenID(ctx, addr.Bytes()); err == nil {
			tokenLastBlock = blockNumber
		}
		th = new(state.TokenHolders).Init(addr, ttype, tokenLastBlock, tokenInfo.ChainID)
		s.tokens[addr] = th
	}
	s.mutex.RUnlock()
	// If the last block of the current scanner is lower than the TokenHolders
	// state block, it seems that the current scanner is out of date and can
	// move on to this block
	if s.lastBlock < th.LastBlock() {
		s.lastBlock = th.LastBlock()
	}
	// get correct web3 uri provider
	w3uri, exists := s.w3p[th.ChainID]
	if !exists {
		return false, fmt.Errorf("chain ID not supported")
	}
	// init web3 contract state
	w3 := state.Web3{}
	if err := w3.Init(ctx, w3uri, addr, th.Type()); err != nil {
		return th.IsSynced(), err
	}
	// try to update the TokenHolders struct and the current scanner last block
	_, err := w3.UpdateTokenHolders(ctx, th)
	if err != nil {
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
		log.Error("warning scanning contract", "token", th.Address().Hex(),
			"block", th.LastBlock(), "error", err)
		return th.IsSynced(), err
	}
	// save TokesHolders state into the database before exit of the function
	return th.IsSynced(), s.saveHolders(th)
}

// calcTokenCreationBlock function attempts to calculate the block number when
// the token contract provided was created and deployed and updates the database
// with the result obtained.
func (s *HoldersScanner) calcTokenCreationBlock(ctx context.Context, addr common.Address) error {
	// set a deadline of 10 seconds from the current context
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	// get the token type
	tokenInfo, err := s.db.QueriesRO.TokenByID(ctx, addr.Bytes())
	if err != nil {
		return fmt.Errorf("error getting token from database: %w", err)
	}
	ttype := state.TokenType(tokenInfo.TypeID)
	// get correct web3 uri provider
	w3uri, exists := s.w3p[tokenInfo.ChainID]
	if !exists {
		return fmt.Errorf("chain ID not supported")
	}
	// init web3 contract state
	w3 := state.Web3{}
	if err := w3.Init(ctx, w3uri, addr, ttype); err != nil {
		return fmt.Errorf("error intializing web3 client for this token: %w", err)
	}
	// get creation block of the current token contract
	creationBlock, err := w3.ContractCreationBlock(ctx)
	if err != nil {
		return fmt.Errorf("error getting token creation block: %w", err)
	}
	dbCreationBlock := new(sql.NullInt64)
	if err := dbCreationBlock.Scan(creationBlock); err != nil {
		return fmt.Errorf("error getting token creation block value: %w", err)
	}
	// save the creation block into the database
	_, err = s.db.QueriesRW.UpdateTokenCreationBlock(ctx,
		queries.UpdateTokenCreationBlockParams{
			ID:            addr.Bytes(),
			CreationBlock: *dbCreationBlock,
		})
	if err != nil {
		return fmt.Errorf("error updating token creation block on the database: %w", err)
	}
	return err
}
