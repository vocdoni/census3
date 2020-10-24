package service

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"git.sr.ht/~sircmpwn/go-bare"
	"github.com/vocdoni/tokenstate/tokenstate"
	"gitlab.com/vocdoni/go-dvote/db"
	"gitlab.com/vocdoni/go-dvote/log"
)

/*
  c<contractAddres> = #block

*/

const contractPrefix = "c_"

var scanSleepTime = time.Second * 30

type Scanner struct {
	dataDir    string
	kv         *db.BadgerDB
	web3       string
	tokens     map[string]*tokenstate.TokenState
	tokensLock sync.RWMutex
	close      chan (bool)
}

type TokenInfo struct {
	Name        string
	Contract    string
	Symbol      string
	Decimals    uint8
	TotalSupply string
	Block       uint64
}

func NewScanner(dataDir string) (*Scanner, error) {
	var s Scanner
	var err error
	s.kv, err = db.NewBadgerDB(dataDir)
	if err != nil {
		return nil, err
	}
	s.dataDir = dataDir
	s.tokens = make(map[string]*tokenstate.TokenState)
	s.close = make(chan (bool))
	return &s, nil
}

func (s *Scanner) Start(w3uri string) {
	s.web3 = w3uri
	s.tokensLock.Lock()
	for _, c := range s.ListContracts() {
		log.Infof("initializing contract %s", c)
		s.tokens[c] = &tokenstate.TokenState{}
		s.tokens[c].Init(s.dataDir, c)
	}
	s.tokensLock.Unlock()
	for {
		select {
		case <-s.close:
			return
		default:
			for _, c := range s.ListContracts() {

				s.scanToken(c)
			}
			time.Sleep(scanSleepTime)
		}
	}
}

func (s *Scanner) Close() {
	s.close <- true
}

func (s *Scanner) GetContract(contract string) (*TokenInfo, error) {
	ib, err := s.kv.Get([]byte(contractPrefix + contract))
	if err != nil {
		return nil, err
	}
	ti := &TokenInfo{}
	err = bare.Unmarshal(ib, ti)
	return ti, err

}

func (s *Scanner) ListContracts() (contracts []string) {
	it := s.kv.NewIterator().(*db.BadgerIterator)
	defer it.Release()
	it.Seek([]byte(contractPrefix))
	for it.Next() {
		contracts = append(contracts, strings.TrimPrefix(string(it.Key()), contractPrefix))
	}
	return
}

func (s *Scanner) AddContract(contract string) error {
	if c, _ := s.GetContract(contract); c != nil {
		return fmt.Errorf("contract %s already exist", contract)
	}
	tinfo, err := s.getOnChainContractData(contract)
	if err != nil {
		return err
	}
	log.Debugf("adding new contract %+v", tinfo)
	if err = s.setContract(tinfo); err != nil {
		return err
	}
	return nil
}

func (s *Scanner) setContract(ti *TokenInfo) error {
	tibytes, err := bare.Marshal(ti)
	if err != nil {
		return err
	}

	if err := s.kv.Put([]byte(contractPrefix+ti.Contract), tibytes); err != nil {
		return err
	}
	return nil
}

func (s *Scanner) getOnChainContractData(contract string) (*TokenInfo, error) {
	ts := tokenstate.Web3{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := ts.Init(ctx, s.web3, contract); err != nil {
		return nil, err
	}
	td, err := ts.GetTokenData()
	if err != nil {
		return nil, err
	}
	return &TokenInfo{Name: td.Name, Contract: contract, Symbol: td.Symbol,
		TotalSupply: td.TotalSupply.String(), Decimals: td.Decimals}, nil

}

func (s *Scanner) scanToken(contract string) {
	w3 := tokenstate.Web3{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	w3.Init(ctx, s.web3, contract)
	cancel()
	tinfo, err := s.GetContract(contract)
	if err != nil {
		log.Error(err)
		return
	}
	log.Debugf("loaded contract data: %+v", tinfo)
	s.tokensLock.RLock()
	ts := s.tokens[contract]
	if ts == nil {
		log.Info("initializing contract %s", contract)
		s.tokens[contract] = &tokenstate.TokenState{}
		s.tokens[contract].Init(s.dataDir, contract)
	}
	s.tokensLock.RUnlock()
	log.Infof("start scanning for token %s from block %d", tinfo.Name, tinfo.Block)
	if tinfo.Block, err = w3.ScanERC20Holders(ts, tinfo.Block, contract); err != nil {
		if strings.Contains(err.Error(), "connection reset") ||
			strings.Contains(err.Error(), "context deadline") ||
			strings.Contains(err.Error(), "read limit exceeded") {
			log.Warnf("connection reset, got until block %d, will retry on next iteration", tinfo.Block)
			if err = s.setContract(tinfo); err != nil {
				log.Error(err)
			}
			return
		}
		log.Error(err)
		return
	}
	log.Infof("successful scanned %s until block %d", tinfo.Name, tinfo.Block)
	if err = s.setContract(tinfo); err != nil {
		log.Error(err)
	}
}
