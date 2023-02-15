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
	"github.com/vocdoni/tokenstate/contractstate/nation3"
	"go.vocdoni.io/dvote/db"
	"go.vocdoni.io/dvote/db/metadb"
	"go.vocdoni.io/dvote/log"
)

type Nation3Scanner struct {
	dataDir    string
	kv         db.Database
	web3       string
	tokens     map[string]*contractstate.ContractState
	tokensLock sync.RWMutex
}

func NewNation3Scanner(dataDir string, w3uri string) (*Nation3Scanner, error) {
	var s Nation3Scanner
	var err error
	s.kv, err = metadb.New(db.TypePebble, dataDir)
	if err != nil {
		return nil, err
	}
	s.dataDir = dataDir
	s.tokens = make(map[string]*contractstate.ContractState)
	s.web3 = w3uri
	return &s, nil
}

func (s *Nation3Scanner) Start(ctx context.Context) {
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
			contractAddresses := [3]string{}
			for idx, c := range s.ListContracts() {
				contractAddresses[idx] = c
			}
			if err := s.scanTokens(ctx, contractAddresses); err != nil {
				log.Error(err)
			}
			time.Sleep(scanSleepTime)
		}
	}
}

func (s *Nation3Scanner) GetContract(contract string) (*TokenInfo, error) {
	tx := s.kv.ReadTx()
	defer tx.Discard()
	ib, err := tx.Get(([]byte(contractPrefix + contract)))
	if err != nil {
		return nil, err
	}
	ti := &TokenInfo{}
	return ti, bare.Unmarshal(ib, ti)

}

func (s *Nation3Scanner) ListContracts() (contracts []string) {
	s.kv.Iterate([]byte(contractPrefix), func(key, value []byte) bool {
		contracts = append(contracts, strings.TrimPrefix(string(key), contractPrefix))
		return true
	})
	return
}

func (s *Nation3Scanner) Balances(contract string) (map[string]*big.Float, error) {
	s.tokensLock.RLock()
	defer s.tokensLock.RUnlock()
	state, ok := s.tokens[contract]
	if !ok {
		return nil, fmt.Errorf("contract %s is unknown", contract)
	}
	return state.List(), nil
}

func (s *Nation3Scanner) Root(contract string, height uint64) ([]byte, error) {
	s.tokensLock.RLock()
	defer s.tokensLock.RUnlock()
	state, ok := s.tokens[contract]
	if !ok {
		return nil, fmt.Errorf("contract %s is unknown", contract)
	}
	return state.Root(height)
}

func (s *Nation3Scanner) QueueExport(contract string) (uint64, error) {
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

func (s *Nation3Scanner) FetchExport(contract string, block uint64) ([]byte, error) {
	s.tokensLock.RLock()
	defer s.tokensLock.RUnlock()
	return os.ReadFile(filepath.Join(
		s.dataDir, "dumps", contract, fmt.Sprintf("%d", block)))
}

func (s *Nation3Scanner) AddContracts(contracts [3]string, startBlock uint64) error {
	for idx, sc := range contracts {
		if c, _ := s.GetContract(sc); c != nil {
			return fmt.Errorf("contract %s already exist", sc)
		}
		tinfo, err := s.getOnChainContractsData(contracts)
		if err != nil {
			return err
		}
		tinfo[idx].StartBlock = startBlock
		tinfo[idx].LastBlock = startBlock
		log.Debugf("adding new contract %+v", *tinfo[idx])
		if err = s.setContract(tinfo[idx]); err != nil {
			return err
		}
	}
	return nil
}

func (s *Nation3Scanner) RescanContract(contract string) error {
	c, err := s.GetContract(contract)
	if err != nil {
		return fmt.Errorf("cannot rescan contract: %w", err)
	}
	c.LastBlock = c.StartBlock
	log.Debugf("queued contract %+v for rescan", c.Contract)
	return s.setContract(c)
}

func (s *Nation3Scanner) setContract(ti *TokenInfo) error {
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

func (s *Nation3Scanner) getOnChainContractsData(contracts [3]string) ([]*TokenInfo, error) {
	w3 := nation3.Nation3{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := w3.Init(ctx, s.web3, contracts); err != nil {
		return nil, err
	}
	tdPassport, err := w3.GetTokenData(nation3.PASSPORT)
	if err != nil {
		return nil, err
	}
	tdVeNation, err := w3.GetTokenData(nation3.VENATION)
	if err != nil {
		return nil, err
	}
	tdNation3, err := w3.GetTokenData(nation3.NATION3)
	if err != nil {
		return nil, err
	}

	tInfo := make([]*TokenInfo, 3)
	tInfo[0] = &TokenInfo{
		Name:        tdPassport.Name,
		Contract:    contracts[0],
		Symbol:      tdPassport.Symbol,
		TotalSupply: tdPassport.TotalSupply.String(),
		Decimals:    tdPassport.Decimals,
	}
	tInfo[1] = &TokenInfo{
		Name:        tdVeNation.Name,
		Contract:    contracts[1],
		Symbol:      tdVeNation.Symbol,
		TotalSupply: tdVeNation.TotalSupply.String(),
		Decimals:    tdVeNation.Decimals,
	}
	tInfo[2] = &TokenInfo{
		Name:        tdNation3.Name,
		Contract:    contracts[2],
		Symbol:      tdNation3.Symbol,
		TotalSupply: tdNation3.TotalSupply.String(),
		Decimals:    tdNation3.Decimals,
	}
	return tInfo, nil
}

func (s *Nation3Scanner) scanTokens(ctx context.Context, contracts [3]string) error {
	w3 := nation3.Nation3{}
	ctx2, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := w3.Init(ctx2, s.web3, contracts); err != nil {
		return err
	}
	for idx, contract := range contracts {
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
		switch idx {
		case nation3.PASSPORT:
			if tinfo.LastBlock, err = w3.ScanPassportHolders(ctx, ts, tinfo.LastBlock+1); err != nil {
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
		case nation3.VENATION:
			if tinfo.LastBlock, err = w3.ScanVeNationHolders(ctx, ts, tinfo.LastBlock+1); err != nil {
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
		case nation3.NATION3:
			if tinfo.LastBlock, err = w3.ScanNation3Holders(ctx, ts, tinfo.LastBlock+1); err != nil {
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
		}
		root, err := s.tokens[contract].LastRoot()
		if err != nil {
			log.Warnf("cannot fetch last root for contract %s: %w", contract, err)
		}
		tinfo.LastRoot = fmt.Sprintf("%x", root)
		if err := s.setContract(tinfo); err != nil {
			return err
		}
		contractCopy := contract
		go func(string) {
			select {
			case <-ctx.Done():
				return
			default:
				// Perform the snapshot and reset of the tree if snapshotBlocks reached
				if tinfo.LastSnapshot+snapshotBlocks <= tinfo.LastBlock {
					if err := s.tokens[contractCopy].Snapshot(); err != nil {
						log.Error(err)
						return
					}
					tinfo.LastSnapshot = tinfo.LastBlock
					if err := s.setContract(tinfo); err != nil {
						log.Error(err)
					}
				}
			}
		}(contractCopy)
	}
	return nil
}
