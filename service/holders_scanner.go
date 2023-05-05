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
	sqlc      *queries.Queries
	lastBlock uint64
}

// NewHoldersScanner function creates a new HolderScanner using the dataDir path
// and the web3 endpoint URI provided. It sets up a sqlite3 database instance
// and gets the number of last block scanned from it.
func NewHoldersScanner(db *queries.Queries, w3uri string) (*HoldersScanner, error) {
	// create an empty scanner
	s := HoldersScanner{
		tokens: make(map[common.Address]*state.TokenHolders),
		web3:   w3uri,
		sqlc:   db,
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
			for _, addr := range tokens {
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
func (s *HoldersScanner) getTokenAddresses() ([]common.Address, error) {
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
	results := make([]common.Address, len(tokens))
	for idx, token := range tokens {
		results[idx] = common.BytesToAddress(token.ID)
	}
	return results, nil
}

// saveTokenHolders function updates the current HoldersScanner database with the
// TokenHolders state provided. Updates the holders for associated token and
// the blocks scanned. To do this, it requires the root hash and the timestampt
// of the given TokenHolders state block.
func (s *HoldersScanner) saveTokenHolders(th *state.TokenHolders) error {
	log.Debugf("saving token holders state for token %s at block %d", th.Address(), th.LastBlock())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
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
		log.Error(err)
		return nil
	}
	rootHash, err := w3.BlockRootHash(ctx, uint(th.LastBlock()))
	if err != nil {
		log.Error(err)
		return err
	}
	// if the current HoldersScanner last block not exists in the database,
	// create it
	if _, err := s.sqlc.BlockByID(ctx, int64(th.LastBlock())); err != nil {
		if !errors.Is(sql.ErrNoRows, err) {
			return err
		}
		_, err = s.sqlc.CreateBlock(ctx, queries.CreateBlockParams{
			ID:        int64(th.LastBlock()),
			Timestamp: timestamp,
			RootHash:  rootHash,
		})
		if err != nil {
			return err
		}
	}
	created, updated, deleted := 0, 0, 0
	log.Debugf("new token holders: %d", len(th.Holders()))
	// iterate over given holders
	//  - if the holder not exists, create it
	//  - if the holder already exists, calculate the new balance with the
	//    current balance
	//		- if the calculated balance is 0 delete it
	//		- if the calculated balance is not 0, update it
	for holder, balance := range th.Holders() {
		currentTokenHolder, err := s.sqlc.TokenHolderByTokenIDAndHolderID(ctx,
			queries.TokenHolderByTokenIDAndHolderIDParams{
				TokenID:  th.Address().Bytes(),
				HolderID: holder.Bytes(),
			})
		if err != nil {
			// return the error if fails and the error is not 'no rows' err
			if !errors.Is(sql.ErrNoRows, err) {
				return err
			}
			_, err = s.sqlc.CreateHolder(ctx, holder.Bytes())
			if err != nil {
				if strings.Contains(err.Error(), "UNIQUE constraint failed") {
					continue
				}
				return err
			}
			// if the token holder not exists, create it
			_, err = s.sqlc.CreateTokenHolder(ctx, queries.CreateTokenHolderParams{
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
		newBalance := new(big.Int).Add(
			new(big.Int).SetBytes(currentTokenHolder.Balance), balance)
		// if the calculated balance is 0 delete it
		if newBalance.Cmp(big.NewInt(0)) <= 0 {
			_, err := s.sqlc.DeleteTokenHolder(ctx,
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
		_, err = s.sqlc.UpdateTokenHolder(ctx, queries.UpdateTokenHolderParams{
			TokenID:  th.Address().Bytes(),
			HolderID: holder.Bytes(),
			BlockID:  int64(th.LastBlock()),
			Balance:  newBalance.Bytes(),
		})
		if err != nil {
			return fmt.Errorf("error updating token holder: %w", err)
		}
		updated++
	}

	log.Debugw("token holders state updated",
		"created", created,
		"updated", updated,
		"deleted", deleted)
	th.FlushHolders()
	log.Debug("token holders state reset")
	return nil
}

// scanHolders function updates the holders of the token identified by the
// address provided. It checks if the address provided already has a
// TokenHolders state cached, if not, it gets the token information from the
// HoldersScanner database and caches it. If something expected fails or the
// scan process ends succesfully, the cached information is stored in the
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
		tokenLastBlock := uint64(tokenInfo.CreationBlock)
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
			if err := s.saveTokenHolders(th); err != nil {
				log.Error(err)
			}
			return nil
		}
		// if unexpected error raises, log it as error and return it.
		log.Error("warning scanning contract", "token", th.Address().Hex(),
			"block", th.LastBlock(), "error", err)
		return err
	}
	// save TokesHolders state into the database before exit of the function
	if err := s.saveTokenHolders(th); err != nil {
		log.Error(err)
	}
	return nil
}
