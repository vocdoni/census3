package service

import (
	"context"
	"database/sql"
	"errors"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"github.com/vocdoni/census3/contractstate"
	"github.com/vocdoni/census3/db"
	queries "github.com/vocdoni/census3/db/sqlc"
	"go.vocdoni.io/dvote/log"
)

type HoldersScanner struct {
	dataDir string
	web3    string
	tokens  map[common.Address]*contractstate.TokenHolders
	mutex   sync.RWMutex
	sqlc    *queries.Queries

	StartBlock uint64
	LastBlock  uint64
}

func NewHoldersScanner(dataDir string, w3uri string) (*HoldersScanner, error) {
	var s HoldersScanner
	s.dataDir = dataDir
	s.tokens = make(map[common.Address]*contractstate.TokenHolders)
	s.web3 = w3uri

	goose.SetBaseFS(db.Census3Migrations)
	if err := goose.SetDialect("sqlite3"); err != nil {
		return nil, err
	}

	var database *sql.DB
	if err := goose.Up(database, "migrations"); err != nil {
		return nil, err
	}

	s.sqlc = queries.New(database)
	return &s, nil
}

func (s *HoldersScanner) Start(ctx context.Context) {
	// load existing contracts
	log.Infof("loading stored contracts...")
	tokens, err := s.ListTokens()
	if err != nil {
		log.Error(err)
		return
	}
	for _, c := range tokens {
		th, err := s.GetHolders(c)
		if err != nil {
			log.Errorf("cannot get contract details for %s: %v", c, err)
			continue
		}
		s.tokens[c] = new(contractstate.TokenHolders).Init(c, th.Type())
	}
	// monitor for new contracts added and update existing
	for {
		select {
		case <-ctx.Done():
			log.Info("scanner loop halted")
			return
		default:
			tokens, err := s.ListTokens()
			if err != nil {
				log.Error(err)
				continue
			}
			for _, c := range tokens {
				if err := s.scanHolders(ctx, c); err != nil {
					log.Error(err)
				}
			}
			log.Info("waiting until next scan iteration")
			time.Sleep(scanSleepTime)
		}
	}
}

func (s *HoldersScanner) ListTokens() ([]common.Address, error) {
	ctx := context.Background()
	tokens, err := s.sqlc.PaginatedTokens(ctx, queries.PaginatedTokensParams{
		Limit:  -1,
		Offset: 0,
	})

	if err != nil {
		return nil, err
	}

	addresses := make([]common.Address, len(tokens))
	for idx, token := range tokens {
		addresses[idx] = common.BigToAddress((*big.Int)(&token.ID))
	}
	return addresses, nil
}

func (s *HoldersScanner) GetHolders(addr common.Address) (*contractstate.TokenHolders, error) {
	ctx := context.Background()
	token, err := s.sqlc.TokenByID(ctx, db.BigInt(*addr.Big()))
	if err != nil {
		return nil, err
	}

	dbHolders, err := s.sqlc.TokenHoldersByTokenID(ctx,
		queries.TokenHoldersByTokenIDParams{
			TokenID: addr.Bytes(),
			Limit:   -1,
			Offset:  0,
		})
	if err != nil {
		return nil, err
	}

	currentHolders := make([]common.Address, len(dbHolders))
	for idx, holder := range dbHolders {
		currentHolders[idx] = common.BytesToAddress(holder)
	}

	th := new(contractstate.TokenHolders).Init(addr, contractstate.ContractType(token.TypeID))
	th.Append(currentHolders...)
	return th, nil
}

func (s *HoldersScanner) SetHolders(th *contractstate.TokenHolders, timestamp string) error {
	ctx := context.Background()

	// if the last block not exists, create it
	if _, err := s.sqlc.BlockByID(ctx, int64(s.LastBlock)); err != nil {
		if !errors.Is(sql.ErrNoRows, err) {
			return err
		}

		_, err = s.sqlc.CreateBlock(ctx, queries.CreateBlockParams{
			ID:        int64(s.LastBlock),
			Timestamp: timestamp,
			RootHash:  []byte{0},
		})
		if err != nil {
			return err
		}
	}

	// iterate over given holders
	for _, holder := range th.Holders() {
		_, err := s.sqlc.TokenHolderByTokenIDAndHolderID(ctx,
			queries.TokenHolderByTokenIDAndHolderIDParams{
				TokenID:  th.Address().Bytes(),
				HolderID: holder.Bytes(),
			})
		if err != nil {
			// return the error if fails and the error is not 'no rows' err
			if !errors.Is(sql.ErrNoRows, err) {
				return err
			}
			// if the token holder not exists, create it
			_, err = s.sqlc.CreateTokenHolder(ctx, queries.CreateTokenHolderParams{
				TokenID:  th.Address().Bytes(),
				HolderID: holder.Bytes(),
				BlockID:  int64(s.LastBlock),
				Balance:  db.BigInt(*big.NewInt(-1)),
			})
			if err != nil {
				return err
			}
		}
		// if exist, update the the block and the holder
		_, err = s.sqlc.UpdateTokenHolder(ctx, queries.UpdateTokenHolderParams{
			TokenID:  th.Address().Bytes(),
			HolderID: holder.Bytes(),
			BlockID:  int64(s.LastBlock),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *HoldersScanner) scanHolders(ctx context.Context, addr common.Address) error {
	log.Debugf("scanning contract %s", addr)
	th, err := s.GetHolders(addr)
	if err != nil {
		return err
	}

	w3 := contractstate.Web3{}
	ctx2, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := w3.Init(ctx2, s.web3, addr, th.Type()); err != nil {
		return err
	}

	s.mutex.RLock()
	th, ok := s.tokens[addr]
	if !ok {
		log.Infof("initializing contract %s (%s)", addr.Hex())
		th = new(contractstate.TokenHolders).Init(addr, th.Type())
		s.tokens[addr] = th
	}
	s.mutex.RUnlock()

	if s.LastBlock, err = w3.UpdateTokenHolders(ctx, th, s.LastBlock+1); err != nil {
		if strings.Contains(err.Error(), "connection reset") ||
			strings.Contains(err.Error(), "context deadline") ||
			strings.Contains(err.Error(), "read limit exceeded") {
			log.Warnf("connection reset on block %d, will retry on next iteration...", s.StartBlock)

			timestamp, err := w3.BlockTimestamp(ctx2, uint(s.LastBlock))
			if err != nil {
				log.Error(err)
				return nil
			}

			if err := s.SetHolders(th, timestamp); err != nil {
				log.Error(err)
			}
			return nil
		}
		return err
	}

	return nil
}
