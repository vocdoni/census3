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

	queries "github.com/vocdoni/census3/db/sqlc"
	"github.com/vocdoni/census3/state"
	"go.vocdoni.io/dvote/log"
)

// HoldersScanner struct contains the needed parameters to scan the holders of
// the tokens stored on the database (located on 'dataDir/dbFilename'). It
// keeps the database updated scanning the network using the web3 endpoint.
type HoldersScanner struct {
	web3      string
	tokens    map[common.Address]*state.TokenHolders
	mutex     sync.RWMutex
	db        *sql.DB
	sqlc      *queries.Queries
	lastBlock uint64
}

// NewHoldersScanner function creates a new HolderScanner using the dataDir path
// and the web3 endpoint URI provided. It sets up a sqlite3 database instance
// and gets the number of last block scanned from it.
func NewHoldersScanner(db *sql.DB, q *queries.Queries, w3uri string) (*HoldersScanner, error) {
	// create an empty scanner
	s := HoldersScanner{
		tokens: make(map[common.Address]*state.TokenHolders),
		web3:   w3uri,
		db:     db,
		sqlc:   q,
	}
	// get latest analyzed block
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	lastBlock, err := s.sqlc.LastBlock(ctx)
	if err == nil {
		s.lastBlock = uint64(lastBlock)
	}
	return &s, nil
}

// Start function initialises the given scanner until the provided context is
// canceled. It first gets the addresses of the tokens to scan and their current
// token state. It then starts scanning, keeping these lists updated and
// synchronised with the database instance.
func (s *HoldersScanner) Start(ctx context.Context) {
	// monitor for new tokens added and update every token holders
	for {
		select {
		case <-ctx.Done():
			log.Info("scanner loop halted")
			return
		default:
			// get updated list of tokens
			tokens, err := s.getTokenAddresses()
			if err != nil {
				log.Error(err)
				continue
			}
			// scan for new holders of every token
			for addr, ready := range tokens {
				if !ready {
					if err := s.calcTokenCreationBlock(ctx, addr); err != nil {
						log.Error(err)
						continue
					}
				}
				if err := s.scanHolders(ctx, addr); err != nil {
					log.Error(err)
				}
			}
			log.Info("waiting until next scan iteration")
			time.Sleep(scanSleepTime)
		}
	}
}

// getTokenAddresses function gets the current token addresses from the database
// and returns it as a list of common.Address structs. If the current database
// instance does not contain any token, it returns nil addresses without error.
// This behaviour helps to deal with this particular case.
func (s *HoldersScanner) getTokenAddresses() (map[common.Address]bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// get tokens from the database
	tokens, err := s.sqlc.ListTokens(ctx)
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

// saveTokenHolders function updates the current HoldersScanner database with the
// TokenHolders state provided. Updates the holders for associated token and
// the blocks scanned. To do this, it requires the root hash and the timestampt
// of the given TokenHolders state block.
func (s *HoldersScanner) saveTokenHolders(th *state.TokenHolders) error {
	log.Debugw("saving token holders",
		"token", th.Address(),
		"block", th.LastBlock(),
		"holders", len(th.Holders()))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// begin a transaction for group sql queries
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Errorw(err, "holders transaction rollback failed")
		}
	}()
	qtx := s.sqlc.WithTx(tx)
	_, err = qtx.UpdateTokenStatus(ctx, queries.UpdateTokenStatusParams{
		Synced: th.IsSynced(),
		ID:     th.Address().Bytes(),
	})
	if err != nil {
		return fmt.Errorf("error updating token: %w", err)
	}
	// if not token holders received, skip
	if len(th.Holders()) == 0 {
		log.Debug("nothing to save. skipping...")
		return nil
	}
	// init web3 contract state
	w3 := state.Web3{}
	if err := w3.Init(ctx, s.web3, th.Address(), th.Type()); err != nil {
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
	if _, err := qtx.BlockByID(ctx, int64(th.LastBlock())); err != nil {
		if !errors.Is(sql.ErrNoRows, err) {
			return err
		}
		_, err = qtx.CreateBlock(ctx, queries.CreateBlockParams{
			ID:        int64(th.LastBlock()),
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
				BlockID:  int64(th.LastBlock()),
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
			NewBlockID: int64(th.LastBlock()),
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
func (s *HoldersScanner) scanHolders(ctx context.Context, addr common.Address) error {
	log.Debugf("scanning contract %s", addr)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	// get the token TokenHolders struct from cache, if it not exists it will
	// be initialized
	s.mutex.RLock()
	th, ok := s.tokens[addr]
	if !ok {
		log.Infof("initializing contract %s", addr.Hex())
		// get token information from the database
		tokenInfo, err := s.sqlc.TokenByID(ctx, addr.Bytes())
		if err != nil {
			return err
		}
		ttype := state.TokenType(tokenInfo.TypeID)
		tokenLastBlock := uint64(tokenInfo.CreationBlock.Int32)
		if blockNumber, err := s.sqlc.LastBlockByTokenID(ctx, addr.Bytes()); err == nil {
			tokenLastBlock = uint64(blockNumber)
		}
		th = new(state.TokenHolders).Init(addr, ttype, tokenLastBlock)
		s.tokens[addr] = th
	}
	s.mutex.RUnlock()
	// If the last block of the current scanner is lower than the TokenHolders
	// state block, it seems that the current scanner is out of date and can
	// move on to this block
	if s.lastBlock < th.LastBlock() {
		s.lastBlock = th.LastBlock()
	}
	// init web3 contract state
	w3 := state.Web3{}
	if err := w3.Init(ctx, s.web3, addr, th.Type()); err != nil {
		return err
	}
	// try to update the TokenHolders struct and the current scanner last block
	var err error
	if _, err = w3.UpdateTokenHolders(ctx, th); err != nil {
		if strings.Contains(err.Error(), "no new blocks") {
			// if no new blocks error raises, log it as debug and return nil
			log.Debugw("no new blocks to scan", "token", th.Address())
			return nil
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
			return s.saveTokenHolders(th)
		}
		// if unexpected error raises, log it as error and return it.
		log.Error("warning scanning contract", "token", th.Address().Hex(),
			"block", th.LastBlock(), "error", err)
		return err
	}
	// save TokesHolders state into the database before exit of the function
	return s.saveTokenHolders(th)
}

// calcTokenCreationBlock function attempts to calculate the block number when
// the token contract provided was created and deployed and updates the database
// with the result obtained.
func (s *HoldersScanner) calcTokenCreationBlock(ctx context.Context, addr common.Address) error {
	// set a deadline of 10 seconds from the current context
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	// get the token type
	tokenInfo, err := s.sqlc.TokenByID(ctx, addr.Bytes())
	if err != nil {
		return err
	}
	ttype := state.TokenType(tokenInfo.TypeID)
	// init web3 contract state
	w3 := state.Web3{}
	if err := w3.Init(ctx, s.web3, addr, ttype); err != nil {
		return err
	}
	// get creation block of the current token contract
	creationBlock, err := w3.ContractCreationBlock(ctx)
	if err != nil {
		return err
	}
	dbCreationBlock := new(sql.NullInt32)
	if err := dbCreationBlock.Scan(creationBlock); err != nil {
		return err
	}
	// save the creation block into the database
	_, err = s.sqlc.UpdateTokenCreationBlock(ctx,
		queries.UpdateTokenCreationBlockParams{
			ID:            addr.Bytes(),
			CreationBlock: *dbCreationBlock,
		})
	return err
}
