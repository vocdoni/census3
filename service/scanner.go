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
	dataDir    string
	kv         db.Database
	web3       string
	tokens     map[string]*contractstate.ContractState
	snapshots  map[string]*contractstate.ContractState
	tokensLock sync.RWMutex
}

type TokenInfo struct {
	Name         string
	Contract     string
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
	s.tokens = make(map[string]*contractstate.ContractState)
	s.snapshots = make(map[string]*contractstate.ContractState)
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
		s.tokens[c].Init(s.dataDir, c, int(contract.Decimals))
	}
	// monitor for new contracts added and update existing
	for {
		select {
		case <-ctx.Done():
			return
		default:
			for _, c := range s.ListContracts() {
				if err := s.scanToken(ctx, c); err != nil {
					log.Error(err)
				}
			}
			time.Sleep(scanSleepTime)
		}
	}
}

func (s *Scanner) GetContract(contract string) (*TokenInfo, error) {
	tx := s.kv.ReadTx()
	defer tx.Discard()
	ib, err := tx.Get(([]byte(contractPrefix + contract)))
	if err != nil {
		return nil, err
	}
	ti := &TokenInfo{}
	return ti, bare.Unmarshal(ib, ti)

}

func (s *Scanner) ListContracts() (contracts []string) {
	s.kv.Iterate([]byte(contractPrefix), func(key, value []byte) bool {
		contracts = append(contracts, strings.TrimPrefix(string(key), contractPrefix))
		return true
	})
	return
}

func (s *Scanner) Balances(contract string) (map[string]*big.Float, error) {
	s.tokensLock.RLock()
	defer s.tokensLock.RUnlock()
	state, ok := s.tokens[contract]
	if !ok {
		return nil, fmt.Errorf("contract %s is unknown", contract)
	}
	return state.List(), nil
}

func (s *Scanner) Root(contract string, height uint64) ([]byte, error) {
	s.tokensLock.RLock()
	defer s.tokensLock.RUnlock()
	state, ok := s.tokens[contract]
	if !ok {
		return nil, fmt.Errorf("contract %s is unknown", contract)
	}
	return state.Root(height)
}

func (s *Scanner) QueueExport(contract string) (uint64, error) {
	s.tokensLock.RLock()
	defer s.tokensLock.RUnlock()
	state, ok := s.tokens[contract]
	if !ok {
		return 0, fmt.Errorf("contract %s is unknown", contract)
	}
	ti, err := s.GetContract(contract)
	if err != nil {
		return 0, err
	}
	go func() {
		dump, err := state.Export()
		if err != nil {
			log.Error(err)
			return
		}
		if err := os.MkdirAll(filepath.Join(s.dataDir, "dumps", contract), 750); err != nil {
			log.Error(err)
			return
		}
		if err := os.WriteFile(
			filepath.Join(
				s.dataDir, "dumps", contract, fmt.Sprintf("%d", ti.LastBlock)),
			dump, 640); err != nil {
			log.Error(err)
			return
		}
		log.Infof("export dump for contract %s and block %d saved", contract, ti.LastBlock)
	}()
	return ti.LastBlock, nil
}

func (s *Scanner) FetchExport(contract string, block uint64) ([]byte, error) {
	s.tokensLock.RLock()
	defer s.tokensLock.RUnlock()
	return os.ReadFile(filepath.Join(
		s.dataDir, "dumps", contract, fmt.Sprintf("%d", block)))
}

func (s *Scanner) AddContract(contract string, startBlock uint64) error {
	if c, _ := s.GetContract(contract); c != nil {
		return fmt.Errorf("contract %s already exist", contract)
	}
	tinfo, err := s.getOnChainContractData(contract)
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

func (s *Scanner) RescanContract(contract string) error {
	c, err := s.GetContract(contract)
	if err != nil {
		return fmt.Errorf("cannot rescan contract: %w", err)
	}
	c.LastBlock = c.StartBlock
	log.Debugf("queued contract %+v for rescan", c.Contract)
	return s.setContract(c)
}

func (s *Scanner) setContract(ti *TokenInfo) error {
	tibytes, err := bare.Marshal(ti)
	if err != nil {
		return err
	}
	wtx := s.kv.WriteTx()
	defer wtx.Discard()
	if err := wtx.Set([]byte(contractPrefix+ti.Contract), tibytes); err != nil {
		return err
	}
	return wtx.Commit()
}

func (s *Scanner) getOnChainContractData(contract string) (*TokenInfo, error) {
	w3 := contractstate.Web3{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := w3.Init(ctx, s.web3, contract); err != nil {
		return nil, err
	}
	td, err := w3.GetTokenData()
	if err != nil {
		return nil, err
	}
	return &TokenInfo{Name: td.Name, Contract: contract, Symbol: td.Symbol,
		TotalSupply: td.TotalSupply.String(), Decimals: td.Decimals}, nil

}

// ChildSnapshot uses the list of holders for parent to update the balances on child.
func (s *Scanner) ChildSnapshot(ctx context.Context, parentContract, childContract string, atBlock uint64) error {
	log.Infof("snapshotting child contract %s at block %d", childContract, atBlock)
	state, ok := s.tokens[parentContract]
	if !ok {
		return fmt.Errorf("contract %s is unknown", parentContract)
	}
	w3 := contractstate.Web3{}
	ctx2, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := w3.Init(ctx2, s.web3, childContract); err != nil {
		return err
	}
	if err := s.AddContract(childContract, atBlock); err != nil {
		return err
	}
	info, err := s.GetContract(childContract)
	if err != nil {
		return err
	}
	s.snapshots[childContract] = new(contractstate.ContractState)
	s.snapshots[childContract].InitSnapshot(s.dataDir, childContract, int(info.Decimals), atBlock)
	return w3.SnapshotChildERC20Holders(ctx, state, s.tokens[childContract], atBlock)
}

func (s *Scanner) scanToken(ctx context.Context, contract string) error {
	w3 := contractstate.Web3{}
	ctx2, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := w3.Init(ctx2, s.web3, contract); err != nil {
		return err
	}
	tinfo, err := s.GetContract(contract)
	if err != nil {
		return err
	}
	log.Debugf("loaded contract data: %+v", *tinfo)
	s.tokensLock.RLock()
	ts, ok := s.tokens[contract]
	if !ok {
		log.Infof("initializing contract %s (%s)", contract, tinfo.Name)
		s.tokens[contract] = new(contractstate.ContractState)
		s.tokens[contract].Init(s.dataDir, contract, int(tinfo.Decimals))
		ts = s.tokens[contract]
	}
	s.tokensLock.RUnlock()

	if tinfo.LastBlock, err = w3.ScanERC20Holders(ctx, ts, tinfo.LastBlock+1); err != nil {
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
	root, err := s.tokens[contract].LastRoot()
	if err != nil {
		log.Warnf("cannot fetch last root for contract %s: %v", contract, err)
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
				if err := s.tokens[contract].Snapshot(); err != nil {
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
