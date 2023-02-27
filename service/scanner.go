package service

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"git.sr.ht/~sircmpwn/go-bare"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/tokenstate/contractstate"
	"go.vocdoni.io/dvote/db"
	"go.vocdoni.io/dvote/db/metadb"
	"go.vocdoni.io/dvote/log"
)

const (
	/*
	   The key value stores the relation:
	   c_<contractAddres> = #block

	*/
	contractPrefix = "c_"   // KV prefix for identify contracts
	snapshotBlocks = 100000 // a snapshot and reset of the tree is performed every snapshotBlocks
	scanSleepTime  = time.Second * 10
)

type Scanner struct {
	dataDir string
	kv      db.Database
	web3    string
	tokens  map[common.Address]*contractstate.ContractState
	mutex   sync.RWMutex
}

type TokenInfo struct {
	Name         string
	Address      common.Address
	Type         contractstate.ContractType
	Symbol       string
	Decimals     uint8
	TotalSupply  string
	StartBlock   uint64
	LastBlock    uint64
	LastRoot     string
	LastSnapshot uint64
}

func NewScanner(dataDir string, w3uri string) (*Scanner, error) {
	var s Scanner
	var err error
	s.kv, err = metadb.New(db.TypePebble, dataDir)
	if err != nil {
		return nil, err
	}
	s.dataDir = dataDir
	s.tokens = make(map[common.Address]*contractstate.ContractState)
	s.web3 = w3uri
	return &s, nil
}

func (s *Scanner) Start(ctx context.Context) {
	// load existing contracts
	log.Infof("loading stored contracts...")
	for _, c := range s.ListContracts() {
		contract, err := s.GetContract(c)
		if err != nil {
			log.Errorf("cannot get contract details for %s: %v", c, err)
			continue
		}
		s.tokens[c] = new(contractstate.ContractState)
		s.tokens[c].Init(s.dataDir, c, contract.Type, int(contract.Decimals))
	}
	// monitor for new contracts added and update existing
	for {
		select {
		case <-ctx.Done():
			log.Info("scanner loop halted")
			return
		default:
			for _, c := range s.ListContracts() {
				if err := s.scanToken(ctx, c); err != nil {
					log.Error(err)
				}
			}
			log.Info("waiting until next scan iteration")
			time.Sleep(scanSleepTime)
		}
	}
}

func (s *Scanner) GetContract(contractAddress common.Address) (*TokenInfo, error) {
	tx := s.kv.ReadTx()
	defer tx.Discard()
	ib, err := tx.Get(([]byte(contractPrefix + contractAddress.Hex())))
	if err != nil {
		return nil, err
	}
	ti := &TokenInfo{}
	if err := bare.Unmarshal(ib, ti); err != nil {
		return nil, err
	}
	log.Debugf("GetContract: %s details: %+v", contractAddress, ti)
	return ti, nil
}

func (s *Scanner) ListContracts() []common.Address {
	contracts := make([]common.Address, 0)
	s.kv.Iterate([]byte(contractPrefix), func(k, v []byte) bool {
		contracts = append(contracts, common.HexToAddress(strings.TrimPrefix(string(k), contractPrefix)))
		log.Debugf("ListContracts(): found contract %s", k)
		return true
	})
	return contracts
}

func (s *Scanner) Balances(contractAddress common.Address) (map[common.Address]*big.Int, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	ts, ok := s.tokens[contractAddress]
	if !ok {
		return nil, fmt.Errorf("contract %s not added", contractAddress)
	}
	return ts.Holders(), nil
}

func (s *Scanner) Root(contractAddress common.Address, height uint64) ([]byte, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	state, ok := s.tokens[contractAddress]
	if !ok {
		return nil, fmt.Errorf("contract %s is unknown", contractAddress)
	}
	return state.BlockRootHash(height)
}

func (s *Scanner) QueueExport(contractAddress common.Address) (uint64, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	state, ok := s.tokens[contractAddress]
	if !ok {
		return 0, fmt.Errorf("contract %s is unknown", contractAddress)
	}
	ti, err := s.GetContract(contractAddress)
	if err != nil {
		return 0, err
	}
	go func() {
		dump, err := state.ExportTree()
		if err != nil {
			log.Error(err)
			return
		}
		if err := os.MkdirAll(filepath.Join(s.dataDir, "dumps", contractAddress.Hex()), 0750); err != nil {
			log.Error(err)
			return
		}
		if err := os.WriteFile(
			filepath.Join(
				s.dataDir, "dumps", contractAddress.Hex(), fmt.Sprintf("%d", ti.LastBlock)),
			dump, 0640); err != nil {
			log.Error(err)
			return
		}
		log.Infof("export dump for contract %s and block %d saved", contractAddress, ti.LastBlock)
	}()
	return ti.LastBlock, nil
}

func (s *Scanner) FetchExport(contractAddress common.Address, block uint64) ([]byte, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return os.ReadFile(filepath.Join(
		s.dataDir, "dumps", contractAddress.Hex(), fmt.Sprintf("%d", block)))
}

func (s *Scanner) AddContract(contractAddress common.Address, contractType contractstate.ContractType, startBlock uint64) error {
	if contractType == contractstate.CONTRACT_TYPE_UNKNOWN {
		return fmt.Errorf("unknown contract type")
	}
	if c, _ := s.GetContract(contractAddress); c != nil {
		return fmt.Errorf("contract %s already exist", contractAddress)
	}
	tinfo, err := s.getOnChainContractData(contractAddress, contractType)
	if err != nil {
		return err
	}
	tinfo.StartBlock = startBlock
	tinfo.LastBlock = startBlock
	log.Debugf("adding new contract %+v", *tinfo)
	if err = s.setContract(tinfo); err != nil {
		return err
	}
	return nil
}

func (s *Scanner) RescanContract(contractAddress common.Address) error {
	c, err := s.GetContract(contractAddress)
	if err != nil {
		return fmt.Errorf("cannot rescan contract: %w", err)
	}
	c.LastBlock = c.StartBlock
	log.Debugf("queued contract %+v for rescan", c.Address)
	return s.setContract(c)
}

func (s *Scanner) setContract(ti *TokenInfo) error {
	tibytes, err := bare.Marshal(ti)
	if err != nil {
		return err
	}
	wtx := s.kv.WriteTx()
	defer wtx.Discard()
	if err := wtx.Set([]byte(contractPrefix+ti.Address.Hex()), tibytes); err != nil {
		return err
	}
	return wtx.Commit()
}

func (s *Scanner) getOnChainContractData(contractAddress common.Address, contractType contractstate.ContractType) (*TokenInfo, error) {
	w3 := contractstate.Web3{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := w3.Init(ctx, s.web3, contractAddress, contractType); err != nil {
		return nil, err
	}
	td, err := w3.GetTokenData()
	if err != nil {
		return nil, err
	}
	return &TokenInfo{Name: td.Name, Address: contractAddress, Type: contractType, Symbol: td.Symbol,
		TotalSupply: td.TotalSupply.String(), Decimals: td.Decimals}, nil
}

func (s *Scanner) scanToken(ctx context.Context, contractAddress common.Address) error {
	log.Debugf("scanning contract %s", contractAddress)
	tinfo, err := s.GetContract(contractAddress)
	if err != nil {
		return err
	}

	w3 := contractstate.Web3{}
	ctx2, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := w3.Init(ctx2, s.web3, contractAddress, tinfo.Type); err != nil {
		return err
	}

	s.mutex.RLock()
	ts, ok := s.tokens[contractAddress]
	if !ok {
		log.Infof("initializing contract %s (%s)", contractAddress.Hex(), tinfo.Name)
		s.tokens[contractAddress] = new(contractstate.ContractState)
		s.tokens[contractAddress].Init(s.dataDir, contractAddress, tinfo.Type, int(tinfo.Decimals))
		ts = s.tokens[contractAddress]
	}
	s.mutex.RUnlock()

	if tinfo.LastBlock, err = w3.ScanTokenHolders(ctx, ts, tinfo.LastBlock+1); err != nil {
		if strings.Contains(err.Error(), "connection reset") ||
			strings.Contains(err.Error(), "context deadline") ||
			strings.Contains(err.Error(), "read limit exceeded") {
			log.Warnf("connection reset on block %d, will retry on next iteration...", tinfo.StartBlock)
			if err := s.setContract(tinfo); err != nil {
				log.Error(err)
			}
			return nil
		}
		return err
	}
	log.Infof("successful scanned %s until block %d", tinfo.Name, tinfo.LastBlock)
	root, err := s.tokens[contractAddress].LastRoot()
	if err != nil {
		log.Warnf("cannot fetch last root for contract %s: %w", contractAddress, err)
	}
	tinfo.LastRoot = fmt.Sprintf("%x", root)
	if err := s.setContract(tinfo); err != nil {
		return err
	}
	go func() {
		select {
		case <-ctx.Done():
			return
		default:
			// Perform the snapshot and reset of the tree if snapshotBlocks reached
			if tinfo.LastSnapshot+snapshotBlocks <= tinfo.LastBlock {
				if err := s.tokens[contractAddress].Snapshot(); err != nil {
					log.Error(err)
					return
				}
				tinfo.LastSnapshot = tinfo.LastBlock
				if err := s.setContract(tinfo); err != nil {
					log.Error(err)
				}
			}
		}
	}()
	return nil
}
