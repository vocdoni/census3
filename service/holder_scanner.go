package service

import (
	"context"
	"database/sql"
	"errors"
	"math/big"
	"path/filepath"
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
		log.Errorw(err, "error during setup goose")
		return nil, err
	}

	database, err := sql.Open("sqlite3", filepath.Join(dataDir, dbFilename))
	if err != nil {
		log.Errorw(err, "error opening database")
		return nil, err
	}

	if err := goose.Up(database, "migrations"); err != nil {
		log.Errorw(err, "error during goose up")
		return nil, err
	}

	s.sqlc = queries.New(database)

	// Get latest analyzed block
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	lastBlock, err := s.sqlc.LastBlock(ctx)
	if err == nil {
		s.StartBlock = uint64(lastBlock)
		s.LastBlock = uint64(lastBlock)
	}

	return &s, nil
}

func (s *HoldersScanner) Start(ctx context.Context) {
	// load existing contracts
	log.Infof("loading stored contracts...")
	tokens, err := s.ListTokens()
	if err != nil {
		if !errors.Is(sql.ErrNoRows, err) {
			log.Error(err)
		}
		return
	}
	for _, addr := range tokens {
		var err error
		if s.tokens[addr], err = s.GetHolders(addr); err != nil {
			log.Errorf("cannot get contract details for %s: %v", addr, err)
			continue
		}
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

func (s *HoldersScanner) AddToken(addr common.Address, tType contractstate.ContractType, startBlock uint64) error {
	w3 := contractstate.Web3{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := w3.Init(ctx, s.web3, addr, tType); err != nil {
		return err
	}
	info, err := w3.GetTokenData()
	if err != nil {
		log.Errorw(err, "error getting token contract data")
		return err
	}
	var (
		name     = new(sql.NullString)
		symbol   = new(sql.NullString)
		decimals = new(sql.NullInt32)
	)
	if err := name.Scan(info.Name); err != nil {
		return err
	}
	if err := symbol.Scan(info.Symbol); err != nil {
		return err
	}
	if err := decimals.Scan(info.Decimals); err != nil {
		return err
	}
	_, err = s.sqlc.CreateToken(ctx, queries.CreateTokenParams{
		ID:            info.Address.Bytes(),
		Name:          *name,
		Symbol:        *symbol,
		Decimals:      *decimals,
		TotalSupply:   info.TotalSupply.Bytes(),
		CreationBlock: int64(startBlock),
		TypeID:        int64(tType),
	})
	if err != nil {
		log.Errorw(err, "error creating token on the database")
		return err
	}
	return nil
}

func (s *HoldersScanner) ListTokens() ([]common.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	tokens, err := s.sqlc.PaginatedTokens(ctx, queries.PaginatedTokensParams{
		Limit:  -1,
		Offset: 0,
	})

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, nil
		}
		return nil, err
	}

	results := make([]common.Address, len(tokens))
	for _, token := range tokens {
		results = append(results, common.BytesToAddress(token.ID))
	}
	return results, nil
}

func (s *HoldersScanner) GetHolders(addr common.Address) (*contractstate.TokenHolders, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	token, err := s.sqlc.TokenByID(ctx, addr.Bytes())
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

	blockNumber, err := s.sqlc.LastBlockByTokenIDAndHolderID(ctx, addr.Bytes())
	if err != nil {
		if !errors.Is(sql.ErrNoRows, err) {
			return nil, err
		}
		blockNumber = 0
	}

	th := new(contractstate.TokenHolders).Init(addr, contractstate.ContractType(token.TypeID), uint64(blockNumber))
	th.Append(currentHolders...)
	return th, nil
}

func (s *HoldersScanner) SetHolders(th *contractstate.TokenHolders, rootHash []byte, timestamp string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// if the last block not exists, create it
	if _, err := s.sqlc.BlockByID(ctx, int64(s.LastBlock)); err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			_, err = s.sqlc.CreateBlock(ctx, queries.CreateBlockParams{
				ID:        int64(s.LastBlock),
				Timestamp: timestamp,
				RootHash:  rootHash,
			})
			if err != nil {
				return err
			}
		} else {
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
			_, err = s.sqlc.CreateHolder(ctx, holder.Bytes())
			if err != nil {
				return err
			}
			// if the token holder not exists, create it
			_, err = s.sqlc.CreateTokenHolder(ctx, queries.CreateTokenHolderParams{
				TokenID:  th.Address().Bytes(),
				HolderID: holder.Bytes(),
				BlockID:  int64(s.LastBlock),
				Balance:  big.NewInt(-1).Bytes(),
			})
			if err != nil {
				return err
			}
			continue
		}
		// if exist, update the the block and the holder
		_, err = s.sqlc.UpdateTokenHolder(ctx, queries.UpdateTokenHolderParams{
			TokenID:  th.Address().Bytes(),
			HolderID: holder.Bytes(),
			BlockID:  int64(s.LastBlock),
			Balance:  big.NewInt(-1).Bytes(),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *HoldersScanner) scanHolders(ctx context.Context, addr common.Address) error {
	log.Debugf("scanning contract %s", addr)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	tokenInfo, err := s.sqlc.TokenByID(ctx, addr.Bytes())
	if err != nil {
		return err
	}
	ttype := contractstate.ContractType(tokenInfo.TypeID)

	if s.LastBlock < uint64(tokenInfo.CreationBlock) {
		s.LastBlock = uint64(tokenInfo.CreationBlock)
	}

	w3 := contractstate.Web3{}
	if err := w3.Init(ctx, s.web3, addr, ttype); err != nil {
		return err
	}

	s.mutex.RLock()
	th, ok := s.tokens[addr]
	if !ok {
		log.Infof("initializing contract %s", addr.Hex())
		th = new(contractstate.TokenHolders).Init(addr, ttype, uint64(tokenInfo.CreationBlock))
		s.tokens[addr] = th
	}
	s.mutex.RUnlock()

	if s.LastBlock, err = w3.UpdateTokenHolders(ctx, th, s.LastBlock+1); err != nil {
		if strings.Contains(err.Error(), "connection reset") ||
			strings.Contains(err.Error(), "context deadline") ||
			strings.Contains(err.Error(), "read limit exceeded") {
			log.Warnf("connection reset on block %d, will retry on next iteration...", s.StartBlock)

			timestamp, err := w3.BlockTimestamp(ctx, uint(s.LastBlock))
			if err != nil {
				log.Error(err)
				return nil
			}
			rootHash, err := w3.BlockRootHash(ctx, uint(s.LastBlock))
			if err != nil {
				log.Error(err)
				return err
			}

			if err := s.SetHolders(th, rootHash, timestamp); err != nil {
				log.Error(err)
			}
			return nil
		}
		return err
	}

	timestamp, err := w3.BlockTimestamp(ctx, uint(s.LastBlock))
	if err != nil {
		log.Error(err)
		return err
	}

	rootHash, err := w3.BlockRootHash(ctx, uint(s.LastBlock))
	if err != nil {
		log.Error(err)
		return err
	}
	if err := s.SetHolders(th, rootHash, timestamp); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
