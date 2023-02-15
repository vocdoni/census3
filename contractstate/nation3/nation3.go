package nation3

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	nation3Passportcontracts "github.com/vocdoni/tokenstate/contracts/nation3/passport"
	nation3Tokencontracts "github.com/vocdoni/tokenstate/contracts/nation3/token"
	nation3VestedTokencontracts "github.com/vocdoni/tokenstate/contracts/nation3/vestedToken"
	"github.com/vocdoni/tokenstate/contractstate"
	"go.vocdoni.io/dvote/log"
	"go.vocdoni.io/dvote/util"
)

const (
	PASSPORT = iota
	VENATION
	NATION3
)

const (
	PASSPORT_TRANSFER_LOG_TOPIC = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	VENATION_DEPOSIT_LOG_TOPIC  = "0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59"
	VENATION_WITHDRAW_LOG_TOPIC = "0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568"
	NATION3_TRANSFER_LOG_TOPIC  = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
)

const (
	maxScanBlocksPerIteration = 1000000
	maxScanLogsPerIteration   = 80000
	blocksToScanAtOnce        = 5000
)

type Contracts struct {
	PassportCaller   *nation3Passportcontracts.Nation3PassportcontractsCaller
	VeNationCaller   *nation3VestedTokencontracts.Nation3VestedTokencontractsCaller
	Nation3Caller    *nation3Tokencontracts.Nation3TokencontractsCaller
	PassportFilterer *nation3Passportcontracts.Nation3PassportcontractsFilterer
	VeNationFilterer *nation3VestedTokencontracts.Nation3VestedTokencontractsFilterer
	Nation3Filterer  *nation3Tokencontracts.Nation3TokencontractsFilterer
	PassportAddress,
	VeNationAddress,
	Nation3Address common.Address
}

type Nation3 struct {
	client    *ethclient.Client
	contacts  *Contracts
	networkID *big.Int
	close     chan (bool)
}

// Init creates and client connection and connects to all Nation3 contracts
// First contract is the Passport contract, second is the VeNation contract, third is the Nation3 contract
// Please respect this order when instantiating the contract addresses
func (n *Nation3) Init(ctx context.Context, web3Endpoint string, contractAddresses [3]string) error {
	var err error
	// connect to ethereum endpoint
	n.client, err = ethclient.Dial(web3Endpoint)
	if err != nil {
		log.Fatal(err)
	}
	n.networkID, err = n.client.ChainID(ctx)
	if err != nil {
		return err
	}
	log.Debugf("found ethereum network id %s", n.networkID.String())

	n.contacts = &Contracts{}
	// passport contract
	c, err := hex.DecodeString(util.TrimHex(contractAddresses[PASSPORT]))
	if err != nil {
		return err
	}
	caddr := common.Address{}
	caddr.SetBytes(c)
	if n.contacts.PassportCaller, err = nation3Passportcontracts.NewNation3PassportcontractsCaller(caddr, n.client); err != nil {
		return err
	}
	if n.contacts.PassportFilterer, err = nation3Passportcontracts.NewNation3PassportcontractsFilterer(caddr, n.client); err != nil {
		return err
	}
	n.contacts.PassportAddress = caddr
	log.Infof("loaded passport contract %s", caddr.String())

	// veNation token contract
	c, err = hex.DecodeString(util.TrimHex(contractAddresses[VENATION]))
	if err != nil {
		return err
	}
	caddr = common.Address{}
	caddr.SetBytes(c)
	if n.contacts.VeNationCaller, err = nation3VestedTokencontracts.NewNation3VestedTokencontractsCaller(caddr, n.client); err != nil {
		return err
	}
	if n.contacts.VeNationFilterer, err = nation3VestedTokencontracts.NewNation3VestedTokencontractsFilterer(caddr, n.client); err != nil {
		return err
	}
	n.contacts.VeNationAddress = caddr
	log.Infof("loaded veNation contract %s", caddr.String())

	// nation3 token contract
	c, err = hex.DecodeString(util.TrimHex(contractAddresses[NATION3]))
	if err != nil {
		return err
	}
	caddr = common.Address{}
	caddr.SetBytes(c)
	if n.contacts.Nation3Caller, err = nation3Tokencontracts.NewNation3TokencontractsCaller(caddr, n.client); err != nil {
		return err
	}
	if n.contacts.Nation3Filterer, err = nation3Tokencontracts.NewNation3TokencontractsFilterer(caddr, n.client); err != nil {
		return err
	}
	n.contacts.Nation3Address = caddr
	log.Infof("loaded nation3 token contract %s", caddr.String())
	return nil
}

// Close closes the Nation3 client connection
func (n *Nation3) Close() {
	n.close <- true
}

// GetTokenData returns the token data for the given token operation
func (n *Nation3) GetTokenData(op uint8) (*contractstate.TokenData, error) {
	var err error
	td := &contractstate.TokenData{}

	switch op {
	case PASSPORT:
		td.Address = n.contacts.PassportAddress.Hex()
		td.Name, err = n.contacts.PassportCaller.Name(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token name: %s", err)
		}
		td.Symbol, err = n.contacts.PassportCaller.Symbol(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token symbol: %s", err)
		}
		td.TotalSupply, err = n.contacts.PassportCaller.TotalSupply(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token total supply: %s", err)
		}
	case VENATION:
		td.Address = n.contacts.VeNationAddress.Hex()
		td.Name, err = n.contacts.VeNationCaller.Name(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token name: %s", err)
		}
		td.Symbol, err = n.contacts.VeNationCaller.Symbol(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token symbol: %s", err)
		}
		decimalsBig, err := n.contacts.VeNationCaller.Decimals(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token decimals: %s", err)
		}
		td.Decimals = uint8(decimalsBig.Uint64())
		td.TotalSupply, err = n.contacts.VeNationCaller.TotalSupply(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token total supply: %s", err)
		}
	case NATION3:
		td.Address = n.contacts.Nation3Address.Hex()
		td.Name, err = n.contacts.Nation3Caller.Name(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token name: %s", err)
		}
		td.Symbol, err = n.contacts.Nation3Caller.Symbol(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token symbol: %s", err)
		}
		td.Decimals, err = n.contacts.Nation3Caller.Decimals(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token decimals: %s", err)
		}
		td.TotalSupply, err = n.contacts.Nation3Caller.TotalSupply(nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get token total supply: %s", err)
		}

	default:
		return nil, fmt.Errorf("invalid operation")
	}

	return td, nil
}

// BalanceOfOrAt returns the balance of the given address, if op = VENATION you can get the
// balance at a the given block
func (n *Nation3) BalanceOfOrAt(ctx context.Context, op uint8, address string, atBlock *big.Int) (*big.Int, error) {
	var err error
	var balance *big.Int

	switch op {
	case PASSPORT:
		balance, err = n.contacts.PassportCaller.BalanceOf(nil, common.HexToAddress(address))
		if err != nil {
			return nil, fmt.Errorf("unable to get passport balance: %s", err)
		}
	case VENATION:
		balance, err = n.contacts.VeNationCaller.BalanceOfAt(nil, common.HexToAddress(address), atBlock)
		if err != nil {
			return nil, fmt.Errorf("unable to get veNation balance: %s", err)
		}
	case NATION3:
		balance, err = n.contacts.Nation3Caller.BalanceOf(nil, common.HexToAddress(address))
		if err != nil {
			return nil, fmt.Errorf("unable to get nation3 balance: %s", err)
		}
	default:
		return nil, fmt.Errorf("invalid operation")
	}

	return balance, nil
}

// GetTotalPassports returns the total number of passports
func (n *Nation3) GetTotalPassports(ctx context.Context) (*big.Int, error) {
	// nextId - 1 = total passports
	nextId, err := n.contacts.PassportCaller.GetNextId(nil)
	if err != nil {
		return nil, fmt.Errorf("unable to get next passport id: %s", err)
	}
	total := big.NewInt(0)
	total.Sub(nextId, big.NewInt(1))
	return total, nil
}

// ScanPassportHolders scans the Ethereum network and updates the nation3 passport holders state
func (n *Nation3) ScanPassportHolders(ctx context.Context,
	ts *contractstate.ContractState, fromBlock uint64) (uint64, error) {

	thash := common.Hash{}
	tbytes, err := hex.DecodeString(PASSPORT_TRANSFER_LOG_TOPIC)
	if err != nil {
		return 0, err
	}
	thash.SetBytes(tbytes)

	from := common.Hash{}
	to := common.Hash{}
	id := big.NewInt(0)

	c, err := hex.DecodeString(util.TrimHex(ts.Contract))
	if err != nil {
		return 0, err
	}
	caddr := common.Address{}
	caddr.SetBytes(c)

	header, err := n.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	toBlock := header.Number.Uint64()
	if fromBlock >= toBlock {
		log.Infof("no new blocks to scan for %s", ts.Contract)
		return toBlock, nil
	}
	if toBlock-fromBlock > maxScanBlocksPerIteration {
		toBlock = fromBlock + maxScanBlocksPerIteration
	}
	var blocks uint64
	var logCount int
	blocks = uint64(blocksToScanAtOnce)
	newBlocks := make(map[uint64]bool)
	log.Infof("start scan iteration for %s from block %d to %d (%d)",
		ts.Contract, fromBlock, toBlock, toBlock-fromBlock)

	for fromBlock < toBlock {
		select {
		// check if we need to close due context signal
		case <-ctx.Done():
			log.Warnf("scan graceful canceled by context")
			return fromBlock, nil

		default:
			startTime := time.Now()
			log.Infof("analyzing blocks from %d to %d [%d%%]", fromBlock,
				fromBlock+blocks, (fromBlock*100)/toBlock)
			query := eth.FilterQuery{
				Addresses: []common.Address{caddr},
				Topics:    [][]common.Hash{{thash}},
				FromBlock: big.NewInt(int64(fromBlock)),
				ToBlock:   big.NewInt(int64(fromBlock + blocks)),
			}
			ctx2, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			logs, err := n.client.FilterLogs(ctx2, query)
			if err != nil {
				if strings.Contains(err.Error(), "query returned more than") {
					blocks = blocks / 2
					log.Warnf("too much results on query, decreasing blocks to %d", blocks)
					continue
				} else {
					return fromBlock, err
				}
			}

			fromBlock += blocks
			if len(logs) == 0 {
				continue
			}
			logCount += len(logs)
			log.Infof("found %d logs, iteration count %d", len(logs), logCount)
			blocksToSave := make(map[uint64]bool)
			for _, l := range logs {
				if _, ok := newBlocks[l.BlockNumber]; !ok {
					if ts.HasBlock(l.BlockNumber) {
						log.Debugf("found already processed block %d", fromBlock)
						continue
					}
				}
				blocksToSave[l.BlockNumber] = true
				newBlocks[l.BlockNumber] = true
				from = l.Topics[1]
				to = l.Topics[2]
				id.SetBytes(l.Data)
				fromAddr := common.BytesToAddress(from.Bytes())
				toAddr := common.BytesToAddress(to.Bytes())
				if err := ts.Add(toAddr, id); err != nil {
					log.Error(err)
				}
				if err := ts.Sub(fromAddr, id); err != nil {
					log.Error(err)
				}
			}
			for k := range blocksToSave {
				ts.Save(k)
			}
			log.Debugf("saved %d blocks at %.2f blocks/second", len(blocksToSave),
				1000*float32(len(blocksToSave))/float32(time.Now().Sub(startTime).Milliseconds()))
			// check if we need to exit because max logs reached for iteration
			if logCount > maxScanLogsPerIteration {
				return fromBlock, nil
			}
		}
	}
	return toBlock, nil
}

// ScanVeNationHolders scans the Ethereum network and updates the venation token holders state
func (n *Nation3) ScanVeNationHolders(ctx context.Context,
	ts *contractstate.ContractState, fromBlock uint64) (uint64, error) {

	// deposit
	thash := common.Hash{}
	tbytes, err := hex.DecodeString(VENATION_DEPOSIT_LOG_TOPIC)
	if err != nil {
		return 0, err
	}
	thash.SetBytes(tbytes)
	// whitdraw
	thash2 := common.Hash{}
	tbytes2, err := hex.DecodeString(VENATION_WITHDRAW_LOG_TOPIC)
	if err != nil {
		return 0, err
	}
	thash2.SetBytes(tbytes2)

	from := common.Hash{}

	c, err := hex.DecodeString(util.TrimHex(ts.Contract))
	if err != nil {
		return 0, err
	}
	caddr := common.Address{}
	caddr.SetBytes(c)

	header, err := n.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	toBlock := header.Number.Uint64()
	if fromBlock >= toBlock {
		log.Infof("no new blocks to scan for %s", ts.Contract)
		return toBlock, nil
	}
	if toBlock-fromBlock > maxScanBlocksPerIteration {
		toBlock = fromBlock + maxScanBlocksPerIteration
	}
	var blocks uint64
	var logCount int
	blocks = uint64(blocksToScanAtOnce)
	newBlocks := make(map[uint64]bool)
	log.Infof("start scan iteration for %s from block %d to %d (%d)",
		ts.Contract, fromBlock, toBlock, toBlock-fromBlock)

	for fromBlock < toBlock {
		select {
		// check if we need to close due context signal
		case <-ctx.Done():
			log.Warnf("scan graceful canceled by context")
			return fromBlock, nil

		default:
			startTime := time.Now()
			log.Infof("analyzing blocks from %d to %d [%d%%]", fromBlock,
				fromBlock+blocks, (fromBlock*100)/toBlock)
			query := eth.FilterQuery{
				Addresses: []common.Address{caddr},
				Topics:    [][]common.Hash{{thash}, {thash2}},
				FromBlock: big.NewInt(int64(fromBlock)),
				ToBlock:   big.NewInt(int64(fromBlock + blocks)),
			}
			ctx2, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			logs, err := n.client.FilterLogs(ctx2, query)
			if err != nil {
				if strings.Contains(err.Error(), "query returned more than") {
					blocks = blocks / 2
					log.Warnf("too much results on query, decreasing blocks to %d", blocks)
					continue
				} else {
					return fromBlock, err
				}
			}

			fromBlock += blocks
			if len(logs) == 0 {
				continue
			}
			logCount += len(logs)
			log.Infof("found %d logs, iteration count %d", len(logs), logCount)
			blocksToSave := make(map[uint64]bool)
			for _, l := range logs {
				switch l.Topics[0] {
				case thash: // DEPOSIT
					if _, ok := newBlocks[l.BlockNumber]; !ok {
						if ts.HasBlock(l.BlockNumber) {
							log.Debugf("found already processed block %d", fromBlock)
							continue
						}
					}
					blocksToSave[l.BlockNumber] = true
					newBlocks[l.BlockNumber] = true
					from = l.Topics[1]
					fromAddr := common.BytesToAddress(from.Bytes())
					logData, err := n.contacts.VeNationFilterer.FilterDeposit(&bind.FilterOpts{
						Start: l.BlockNumber,
						End:   &l.BlockNumber,
					}, []common.Address{fromAddr}, nil)
					if err != nil {
						log.Error(err)
					}
					if err := ts.Add(logData.Event.Provider, logData.Event.Value); err != nil {
						log.Error(err)
					}
				case thash2: // WITHDRAW
					if _, ok := newBlocks[l.BlockNumber]; !ok {
						if ts.HasBlock(l.BlockNumber) {
							log.Debugf("found already processed block %d", fromBlock)
							continue
						}
					}
					blocksToSave[l.BlockNumber] = true
					newBlocks[l.BlockNumber] = true
					from = l.Topics[1]
					fromAddr := common.BytesToAddress(from.Bytes())
					logData, err := n.contacts.VeNationFilterer.FilterDeposit(&bind.FilterOpts{
						Start: l.BlockNumber,
						End:   &l.BlockNumber,
					}, []common.Address{fromAddr}, nil)
					if err != nil {
						log.Error(err)
					}
					if err := ts.Sub(logData.Event.Provider, logData.Event.Value); err != nil {
						log.Error(err)
					}
				}
			}
			for k := range blocksToSave {
				ts.Save(k)
			}
			log.Debugf("saved %d blocks at %.2f blocks/second", len(blocksToSave),
				1000*float32(len(blocksToSave))/float32(time.Now().Sub(startTime).Milliseconds()))
			// check if we need to exit because max logs reached for iteration
			if logCount > maxScanLogsPerIteration {
				return fromBlock, nil
			}
		}
	}
	return toBlock, nil
}

// ScanNation3Holders scans the Ethereum network and updates the nation3 holders state
func (n *Nation3) ScanNation3Holders(ctx context.Context,
	ts *contractstate.ContractState, fromBlock uint64) (uint64, error) {

	thash := common.Hash{}
	tbytes, err := hex.DecodeString(NATION3_TRANSFER_LOG_TOPIC)
	if err != nil {
		return 0, err
	}
	thash.SetBytes(tbytes)

	from := common.Hash{}
	to := common.Hash{}
	amount := big.NewInt(0)

	c, err := hex.DecodeString(util.TrimHex(ts.Contract))
	if err != nil {
		return 0, err
	}
	caddr := common.Address{}
	caddr.SetBytes(c)

	header, err := n.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	toBlock := header.Number.Uint64()
	if fromBlock >= toBlock {
		log.Infof("no new blocks to scan for %s", ts.Contract)
		return toBlock, nil
	}
	if toBlock-fromBlock > maxScanBlocksPerIteration {
		toBlock = fromBlock + maxScanBlocksPerIteration
	}
	var blocks uint64
	var logCount int
	blocks = uint64(blocksToScanAtOnce)
	newBlocks := make(map[uint64]bool)
	log.Infof("start scan iteration for %s from block %d to %d (%d)",
		ts.Contract, fromBlock, toBlock, toBlock-fromBlock)

	for fromBlock < toBlock {
		select {
		// check if we need to close due context signal
		case <-ctx.Done():
			log.Warnf("scan graceful canceled by context")
			return fromBlock, nil

		default:
			startTime := time.Now()
			log.Infof("analyzing blocks from %d to %d [%d%%]", fromBlock,
				fromBlock+blocks, (fromBlock*100)/toBlock)
			query := eth.FilterQuery{
				Addresses: []common.Address{caddr},
				Topics:    [][]common.Hash{{thash}},
				FromBlock: big.NewInt(int64(fromBlock)),
				ToBlock:   big.NewInt(int64(fromBlock + blocks)),
			}
			ctx2, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			logs, err := n.client.FilterLogs(ctx2, query)
			if err != nil {
				if strings.Contains(err.Error(), "query returned more than") {
					blocks = blocks / 2
					log.Warnf("too much results on query, decreasing blocks to %d", blocks)
					continue
				} else {
					return fromBlock, err
				}
			}

			fromBlock += blocks
			if len(logs) == 0 {
				continue
			}
			logCount += len(logs)
			log.Infof("found %d logs, iteration count %d", len(logs), logCount)
			blocksToSave := make(map[uint64]bool)
			for _, l := range logs {
				if _, ok := newBlocks[l.BlockNumber]; !ok {
					if ts.HasBlock(l.BlockNumber) {
						log.Debugf("found already processed block %d", fromBlock)
						continue
					}
				}
				blocksToSave[l.BlockNumber] = true
				newBlocks[l.BlockNumber] = true
				from = l.Topics[1]
				to = l.Topics[2]
				amount.SetBytes(l.Data)
				fromAddr := common.BytesToAddress(from.Bytes())
				toAddr := common.BytesToAddress(to.Bytes())
				if err := ts.Add(toAddr, amount); err != nil {
					log.Error(err)
				}
				if err := ts.Sub(fromAddr, amount); err != nil {
					log.Error(err)
				}
			}
			for k := range blocksToSave {
				ts.Save(k)
			}
			log.Debugf("saved %d blocks at %.2f blocks/second", len(blocksToSave),
				1000*float32(len(blocksToSave))/float32(time.Now().Sub(startTime).Milliseconds()))
			// check if we need to exit because max logs reached for iteration
			if logCount > maxScanLogsPerIteration {
				return fromBlock, nil
			}
		}
	}
	return toBlock, nil
}
