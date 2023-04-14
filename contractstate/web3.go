package contractstate

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	want "github.com/vocdoni/census3/contracts/aragon/want"
	erc1155 "github.com/vocdoni/census3/contracts/erc/erc1155"
	erc20 "github.com/vocdoni/census3/contracts/erc/erc20"
	erc721 "github.com/vocdoni/census3/contracts/erc/erc721"
	erc777 "github.com/vocdoni/census3/contracts/erc/erc777"
	venation "github.com/vocdoni/census3/contracts/nation3/vestedToken"

	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.vocdoni.io/dvote/log"
)

// Web3 holds a reference to a web3 client and a contract, as
// well as storing the contract address, the contract type and
// the network id reported by the endpoint connection
type Web3 struct {
	client *ethclient.Client

	contract        interface{}
	contractType    ContractType
	contractAddress common.Address
}

func (w *Web3) NewContract() (interface{}, error) {
	switch w.contractType {
	case CONTRACT_TYPE_ERC20:
		return erc20.NewERC20Contract(w.contractAddress, w.client)
	case CONTRACT_TYPE_ERC721:
		return erc721.NewERC721Contract(w.contractAddress, w.client)
	case CONTRACT_TYPE_ERC1155:
		return erc1155.NewERC1155Contract(w.contractAddress, w.client)
	case CONTRACT_TYPE_ERC777:
		return erc777.NewERC777Contract(w.contractAddress, w.client)
	case CONTRACT_TYPE_CUSTOM_NATION3_VENATION:
		return venation.NewNation3VestedTokenContract(w.contractAddress, w.client)
	case CONTRACT_TYPE_CUSTOM_ARAGON_WANT:
		return want.NewAragonWrappedANTTokenContract(w.contractAddress, w.client)
	default:
		return nil, fmt.Errorf("unknown contract type %d", w.contractType)
	}
}

// Init creates and client connection and connects to contract given its address
func (w *Web3) Init(ctx context.Context, web3Endpoint string, contractAddress common.Address, contractType ContractType) error {
	var err error
	// connect to ethereum endpoint
	w.client, err = ethclient.Dial(web3Endpoint)
	if err != nil {
		log.Fatal(err)
	}
	switch contractType {
	case CONTRACT_TYPE_ERC20,
		CONTRACT_TYPE_ERC721,
		CONTRACT_TYPE_ERC1155,
		CONTRACT_TYPE_ERC777,
		CONTRACT_TYPE_CUSTOM_NATION3_VENATION,
		CONTRACT_TYPE_CUSTOM_ARAGON_WANT:
		w.contractType = contractType
	default:
		return fmt.Errorf("unknown contract type %d", contractType)
	}
	w.contractAddress = contractAddress
	if w.contract, err = w.NewContract(); err != nil {
		return err
	}
	log.Infof("loaded token contract %s", contractAddress)
	return nil
}

func (w *Web3) Close() {
	w.client.Close()
}

// TokenName wraps the name() function contract call
func (w *Web3) TokenName() (string, error) {
	switch w.contractType {
	case CONTRACT_TYPE_ERC20:
		caller := w.contract.(*erc20.ERC20Contract).ERC20ContractCaller
		return caller.Name(nil)
	case CONTRACT_TYPE_ERC721:
		caller := w.contract.(*erc721.ERC721Contract).ERC721ContractCaller
		return caller.Name(nil)
	case CONTRACT_TYPE_ERC777:
		caller := w.contract.(*erc777.ERC777Contract).ERC777ContractCaller
		return caller.Name(nil)
	case CONTRACT_TYPE_ERC1155:
		return "", nil
	case CONTRACT_TYPE_CUSTOM_NATION3_VENATION:
		caller := w.contract.(*venation.Nation3VestedTokenContract).Nation3VestedTokenContractCaller
		return caller.Name(nil)
	case CONTRACT_TYPE_CUSTOM_ARAGON_WANT:
		caller := w.contract.(*want.AragonWrappedANTTokenContract).AragonWrappedANTTokenContractCaller
		return caller.Name(nil)

	}
	return "", fmt.Errorf("unknown contract type %d", w.contractType)
}

// TokenSymbol wraps the symbol() function contract call
func (w *Web3) TokenSymbol() (string, error) {
	switch w.contractType {
	case CONTRACT_TYPE_ERC20:
		caller := w.contract.(*erc20.ERC20Contract).ERC20ContractCaller
		return caller.Symbol(nil)
	case CONTRACT_TYPE_ERC721:
		caller := w.contract.(*erc721.ERC721Contract).ERC721ContractCaller
		return caller.Symbol(nil)
	case CONTRACT_TYPE_ERC777:
		caller := w.contract.(*erc777.ERC777Contract).ERC777ContractCaller
		return caller.Symbol(nil)
	case CONTRACT_TYPE_ERC1155:
		return "", nil
	case CONTRACT_TYPE_CUSTOM_NATION3_VENATION:
		caller := w.contract.(*venation.Nation3VestedTokenContract).Nation3VestedTokenContractCaller
		return caller.Symbol(nil)
	case CONTRACT_TYPE_CUSTOM_ARAGON_WANT:
		caller := w.contract.(*want.AragonWrappedANTTokenContract).AragonWrappedANTTokenContractCaller
		return caller.Symbol(nil)
	}
	return "", fmt.Errorf("unknown contract type %d", w.contractType)
}

// TokenDecimals wraps the decimals() function contract call
func (w *Web3) TokenDecimals() (uint8, error) {
	switch w.contractType {
	case CONTRACT_TYPE_ERC20:
		caller := w.contract.(*erc20.ERC20Contract).ERC20ContractCaller
		return caller.Decimals(nil)
	case CONTRACT_TYPE_ERC721:
		return 0, nil
	case CONTRACT_TYPE_ERC777:
		caller := w.contract.(*erc777.ERC777Contract).ERC777ContractCaller
		return caller.Decimals(nil)
	case CONTRACT_TYPE_ERC1155:
		return 0, nil
	case CONTRACT_TYPE_CUSTOM_NATION3_VENATION:
		caller := w.contract.(*venation.Nation3VestedTokenContract).Nation3VestedTokenContractCaller
		decimals, err := caller.Decimals(nil)
		if err != nil {
			return 0, err
		}
		return uint8(decimals.Uint64()), nil
	case CONTRACT_TYPE_CUSTOM_ARAGON_WANT:
		caller := w.contract.(*want.AragonWrappedANTTokenContract).AragonWrappedANTTokenContractCaller
		return caller.Decimals(nil)
	}
	return 0, fmt.Errorf("unknown contract type %d", w.contractType)
}

// TokenTotalSupply wraps the totalSupply function contract call
func (w *Web3) TokenTotalSupply() (*big.Int, error) {
	switch w.contractType {
	case CONTRACT_TYPE_ERC20:
		caller := w.contract.(*erc20.ERC20Contract).ERC20ContractCaller
		return caller.TotalSupply(nil)
	case CONTRACT_TYPE_ERC721:
		return nil, nil
	case CONTRACT_TYPE_ERC777:
		caller := w.contract.(*erc777.ERC777Contract).ERC777ContractCaller
		return caller.TotalSupply(nil)
	case CONTRACT_TYPE_ERC1155:
		return nil, nil
	case CONTRACT_TYPE_CUSTOM_NATION3_VENATION:
		caller := w.contract.(*venation.Nation3VestedTokenContract).Nation3VestedTokenContractCaller
		return caller.TotalSupply(nil)
	case CONTRACT_TYPE_CUSTOM_ARAGON_WANT:
		caller := w.contract.(*want.AragonWrappedANTTokenContract).AragonWrappedANTTokenContractCaller
		return caller.TotalSupply(nil)
	}
	return nil, fmt.Errorf("unknown contract type %d", w.contractType)
}

// TokenBalanceOf wraps the balanceOf function contract call
// CASE ERC1155: args[0] is the tokenID
// CASE NATION3_VENATION: args[0] is the function to call (0: BalanceOf, 1: BalanceOfAt)
// CASE NATION3_VENATION: args[1] is the block number when calling BalanceOfAt, otherwise it is ignored
// CASE WANT: args[0] is the function to call (0: BalanceOf, 1: BalanceOfAt)
// CASE WANT: args[1] is the block number when calling BalanceOfAt, otherwise it is ignored
func (w *Web3) TokenBalanceOf(tokenHolderAddress common.Address, args ...interface{}) (*big.Int, error) {
	switch w.contractType {
	case CONTRACT_TYPE_ERC20:
		caller := w.contract.(*erc20.ERC20Contract).ERC20ContractCaller
		return caller.BalanceOf(nil, tokenHolderAddress)
	case CONTRACT_TYPE_ERC721:
		caller := w.contract.(*erc721.ERC721Contract).ERC721ContractCaller
		return caller.BalanceOf(nil, tokenHolderAddress)
	case CONTRACT_TYPE_ERC777:
		caller := w.contract.(*erc777.ERC777Contract).ERC777ContractCaller
		return caller.BalanceOf(nil, tokenHolderAddress)
	case CONTRACT_TYPE_ERC1155:
		if len(args) != 1 {
			return nil, fmt.Errorf("wrong number of arguments for ERC1155 balanceOf function")
		}
		caller := w.contract.(*erc1155.ERC1155Contract).ERC1155ContractCaller
		return caller.BalanceOf(nil, tokenHolderAddress, big.NewInt(int64(args[0].(uint64))))
	case CONTRACT_TYPE_CUSTOM_NATION3_VENATION:
		if len(args) != 2 {
			return nil, fmt.Errorf("wrong number of arguments for Nation3VestedToken balanceOf function")
		}
		caller := w.contract.(*venation.Nation3VestedTokenContract).Nation3VestedTokenContractCaller
		switch args[0].(int) {
		case 0:
			return caller.BalanceOf(nil, tokenHolderAddress)
		case 1:
			return caller.BalanceOfAt(nil, tokenHolderAddress, big.NewInt(int64(args[1].(uint64))))
		}
	case CONTRACT_TYPE_CUSTOM_ARAGON_WANT:
		if len(args) != 2 {
			return nil, fmt.Errorf("wrong number of arguments for Nation3VestedToken balanceOf function")
		}
		caller := w.contract.(*want.AragonWrappedANTTokenContract).AragonWrappedANTTokenContractCaller
		switch args[0].(int) {
		case 0:
			return caller.BalanceOf(nil, tokenHolderAddress)
		case 1:
			return caller.BalanceOfAt(nil, tokenHolderAddress, big.NewInt(int64(args[1].(uint64))))
		}
	}
	return nil, fmt.Errorf("unknown contract type %d", w.contractType)
}

// ScantokenHolders scans the Ethereum network and updates the token holders state
// Returns the last block number scanned
func (w *Web3) ScanTokenHolders(ctx context.Context, ts *ContractState, fromBlockNumber uint64) (uint64, error) {
	// fetch the last block header
	lastBlockHeader, err := w.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	log.Debugf("last block number: %d", lastBlockHeader.Number.Uint64())
	// check if there are new blocks to scan
	lastBlockNumber := lastBlockHeader.Number.Uint64()
	if fromBlockNumber >= lastBlockNumber {
		log.Infof("no new blocks to scan for %s", ts.Address().String())
		return fromBlockNumber, nil
	}
	// check if we need to scan more than MAX_SCAN_BLOCKS_PER_ITERATION
	// if so, scan only MAX_SCAN_BLOCKS_PER_ITERATION blocks
	if lastBlockNumber-fromBlockNumber > MAX_SCAN_BLOCKS_PER_ITERATION {
		lastBlockNumber = fromBlockNumber + MAX_SCAN_BLOCKS_PER_ITERATION
	}

	log.Infof("start scan iteration for %s from block %d to %d", ts.Address().Hex(), fromBlockNumber, lastBlockNumber)
	blocks := BLOCKS_TO_SCAN_AT_ONCE
	logCount := 0
	newBlocksMap := make(map[uint64]bool)

	for fromBlockNumber < lastBlockNumber {
		select {
		// check if we need to close due context signal
		case <-ctx.Done():
			log.Warnf("scan graceful canceled by context")
			return fromBlockNumber, nil

		default:
			startTime := time.Now()
			log.Infof("analyzing blocks from %d to %d [%d%%]", fromBlockNumber,
				fromBlockNumber+blocks, (fromBlockNumber*100)/lastBlockNumber)

			// create the filter query
			query := eth.FilterQuery{
				Addresses: []common.Address{ts.Address()},
				FromBlock: big.NewInt(int64(fromBlockNumber)),
				ToBlock:   big.NewInt(int64(fromBlockNumber + blocks)),
			}
			// set the topics to filter depending on the contract type
			switch ts.Type() {
			case CONTRACT_TYPE_ERC20, CONTRACT_TYPE_ERC777, CONTRACT_TYPE_ERC721:
				query.Topics = [][]common.Hash{{common.HexToHash(LOG_TOPIC_ERC20_TRANSFER)}}
			case CONTRACT_TYPE_ERC1155:
				query.Topics = [][]common.Hash{
					{
						common.HexToHash(LOG_TOPIC_ERC1155_TRANSFER_SINGLE),
						common.HexToHash(LOG_TOPIC_ERC1155_TRANSFER_BATCH),
					},
				}
			case CONTRACT_TYPE_CUSTOM_NATION3_VENATION:
				query.Topics = [][]common.Hash{
					{
						common.HexToHash(LOG_TOPIC_VENATION_DEPOSIT),
						common.HexToHash(LOG_TOPIC_VENATION_WITHDRAW),
					},
				}
			case CONTRACT_TYPE_CUSTOM_ARAGON_WANT:
				query.Topics = [][]common.Hash{
					{
						common.HexToHash(LOG_TOPIC_WANT_DEPOSIT),
						common.HexToHash(LOG_TOPIC_WANT_WITHDRAWAL),
					},
				}
			default:
				return 0, fmt.Errorf("unknown contract type %d", ts.Type())
			}
			// execute the filter query
			ctx2, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			logs, err := w.client.FilterLogs(ctx2, query)
			if err != nil {
				// if we have too much results, decrease the blocks to scan
				if strings.Contains(err.Error(), "query returned more than") {
					blocks = blocks / 2
					log.Warnf("too much results on query, decreasing blocks to %d", blocks)
					continue
				} else {
					return fromBlockNumber, err
				}
			}
			fromBlockNumber += blocks
			if len(logs) == 0 {
				continue
			}
			logCount += len(logs)
			log.Infof("found %d logs, iteration count %d", len(logs), logCount)
			blocksToSave := make(map[uint64]bool)
			// iterate over the logs and update the token holders state
			for _, l := range logs {
				if _, ok := newBlocksMap[l.BlockNumber]; !ok {
					if ts.HasBlock(l.BlockNumber) {
						log.Debugf("found already processed block %d", fromBlockNumber)
						continue
					}
				}
				blocksToSave[l.BlockNumber] = true
				newBlocksMap[l.BlockNumber] = true
				// update the token holders state with the log data
				switch ts.Type() {
				case CONTRACT_TYPE_ERC20:
					filter := w.contract.(*erc20.ERC20Contract).ERC20ContractFilterer
					logData, err := filter.ParseTransfer(l)
					if err != nil {
						log.Errorf("error parsing log data %s", err)
						continue
					}
					if err := ts.Add(logData.To, logData.Value); err != nil {
						log.Errorf("error adding to token holder %s", err)
						continue
					}
					if err := ts.Sub(logData.From, logData.Value); err != nil {
						log.Errorf("error subtracting to token holder %s", err)
						continue
					}
				case CONTRACT_TYPE_ERC777: // stores the total count per address, not all identifiers
					filter := w.contract.(*erc777.ERC777Contract).ERC777ContractFilterer
					logData, err := filter.ParseTransfer(l)
					if err != nil {
						log.Errorf("error parsing log data %s", err)
						continue
					}
					if err := ts.Add(logData.To, big.NewInt(1)); err != nil {
						log.Errorf("error adding to token holder %s", err)
						continue
					}
					if err := ts.Sub(logData.From, big.NewInt(1)); err != nil {
						log.Errorf("error subtracting to token holder %s", err)
						continue
					}
				case CONTRACT_TYPE_ERC721: // stores the total count per address, not all identifiers
					filter := w.contract.(*erc721.ERC721Contract).ERC721ContractFilterer
					logData, err := filter.ParseTransfer(l)
					if err != nil {
						log.Errorf("error parsing log data %s", err)
						continue
					}
					if err := ts.Add(logData.To, big.NewInt(1)); err != nil {
						log.Errorf("error adding to token holder %s", err)
						continue
					}
					if err := ts.Sub(logData.From, big.NewInt(1)); err != nil {
						log.Errorf("error subtracting to token holder %s", err)
						continue
					}
					// case CONTRACT_TYPE_ERC1155:
					// TODO
					/*
						filter := w.contract.(*erc1155.ERC1155Contract).ERC1155ContractFilterer
						if l.Topics[0] == common.HexToHash(LOG_TOPIC_ERC1155_TRANSFER_SINGLE) {
							logData, err := filter.ParseTransferSingle(l)
							if err != nil {
								log.Error("error parsing log data", err)
								continue
							}
							// TODO
						} else if l.Topics[0] == common.HexToHash(LOG_TOPIC_ERC1155_TRANSFER_BATCH) {
							logData, err := filter.ParseTransferBatch(l)
							if err != nil {
								log.Error("error parsing log data", err)
								continue
							}
							// TODO
						}
					*/
				case CONTRACT_TYPE_CUSTOM_NATION3_VENATION:
					// This token contract is a bit special, token balances
					// are updated every block based on the contract state.
					switch l.Topics[0] {
					case common.HexToHash(LOG_TOPIC_VENATION_DEPOSIT):
						provider := common.HexToAddress(l.Topics[1].Hex())
						value := big.NewInt(0).SetBytes(l.Data[:32])
						if err := ts.Add(provider, value); err != nil {
							log.Errorf("error adding to token holder %s", err)
						}
					case common.HexToHash(LOG_TOPIC_VENATION_WITHDRAW):
						provider := common.HexToAddress(l.Topics[1].Hex())
						value := big.NewInt(0).SetBytes(l.Data[:32])
						if err := ts.Sub(provider, value); err != nil {
							log.Errorf("error subtracting to token holder %s", err)
						}
					}
				case CONTRACT_TYPE_CUSTOM_ARAGON_WANT:
					filter := w.contract.(*want.AragonWrappedANTTokenContract).AragonWrappedANTTokenContractFilterer
					switch l.Topics[0] {
					case common.HexToHash(LOG_TOPIC_WANT_DEPOSIT):
						logData, err := filter.ParseDeposit(l)
						if err != nil {
							log.Errorf("error parsing log data %s", err)
							continue
						}
						if err := ts.Add(logData.Entity, logData.Amount); err != nil {
							log.Errorf("error adding to token holder %s", err)
							continue
						}
					case common.HexToHash(LOG_TOPIC_WANT_WITHDRAWAL):
						logData, err := filter.ParseWithdrawal(l)
						if err != nil {
							log.Errorf("error parsing log data %s", err)
							continue
						}
						if err := ts.Sub(logData.Entity, logData.Amount); err != nil {
							log.Errorf("error subtracting to token holder %s", err)
							continue
						}
					}
				}
			}
			for k := range blocksToSave {
				if err := ts.SaveBlock(k); err != nil {
					log.Errorf("error saving block %d %s", k, err)
					continue
				}
			}
			log.Debugf("saved %d blocks at %.2f blocks/second", len(blocksToSave),
				1000*float32(len(blocksToSave))/float32(time.Since(startTime).Milliseconds()))
			// check if we need to exit because max logs reached for iteration
			if logCount > MAX_SCAN_LOGS_PER_ITERATION {
				return fromBlockNumber, nil
			}
		}
	}
	return lastBlockNumber, nil
}

// UpdateTokenHolders function checks the transfer logs of the given contract
// (in the TokenHolders struct) from the given block number. It gets all
// addresses (candidates to holders) and their balances from the given block
// number to the latest block number. Any address with a balance not equal to
// zero will is stored as a new holder. It also checks the balances of the
// current holders, deleting those that were not found.
func (w *Web3) UpdateTokenHolders(ctx context.Context, th *TokenHolders, fromBlockNumber uint64) (uint64, error) {
	// fetch the last block header
	lastBlockHeader, err := w.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	log.Debugf("last block number: %d", lastBlockHeader.Number.Uint64())
	// check if there are new blocks to scan
	lastBlockNumber := lastBlockHeader.Number.Uint64()
	if fromBlockNumber >= lastBlockNumber {
		log.Infof("no new blocks to scan for %s", th.Address().String())
		return fromBlockNumber, nil
	}
	// check if we need to scan more than MAX_SCAN_BLOCKS_PER_ITERATION
	// if so, scan only MAX_SCAN_BLOCKS_PER_ITERATION blocks
	if lastBlockNumber-fromBlockNumber > MAX_SCAN_BLOCKS_PER_ITERATION {
		lastBlockNumber = fromBlockNumber + MAX_SCAN_BLOCKS_PER_ITERATION
	}
	blocks := BLOCKS_TO_SCAN_AT_ONCE
	log.Infof("start scan iteration for %s from block %d to %d", th.Address().Hex(), fromBlockNumber, lastBlockNumber)

	// get logs and get new candidates to holder. A valid candidate is every
	// address with a positive balance at the end of logs review. It requires
	// take into account the countability of the candidates' balances.
	logCount := 0
	newBlocksMap := make(map[uint64]bool)
	holdersCandidates := make(map[common.Address]*big.Int)
	for fromBlockNumber < lastBlockNumber {
		select {
		// check if we need to close due context signal
		case <-ctx.Done():
			log.Warnf("scan graceful canceled by context")
			return fromBlockNumber, nil
		default:
			log.Infof("analyzing blocks from %d to %d [%d%%]", fromBlockNumber,
				fromBlockNumber+blocks, (fromBlockNumber*100)/lastBlockNumber)
			// create the filter query
			query := eth.FilterQuery{
				Addresses: []common.Address{th.Address()},
				FromBlock: big.NewInt(int64(fromBlockNumber)),
				ToBlock:   big.NewInt(int64(fromBlockNumber + blocks)),
			}
			// set the topics to filter depending on the contract type
			switch th.Type() {
			case CONTRACT_TYPE_ERC20:
				query.Topics = [][]common.Hash{{common.HexToHash(LOG_TOPIC_ERC20_TRANSFER)}}
			default:
				return 0, fmt.Errorf("unknown contract type %d", th.Type())
			}
			// execute the filter query
			ctx2, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			logs, err := w.client.FilterLogs(ctx2, query)
			if err != nil {
				// if we have too much results, decrease the blocks to scan
				if strings.Contains(err.Error(), "query returned more than") {
					blocks = blocks / 2
					log.Warnf("too much results on query, decreasing blocks to %d", blocks)
					continue
				} else {
					return fromBlockNumber, err
				}
			}
			fromBlockNumber += blocks
			if len(logs) == 0 {
				continue
			}
			logCount += len(logs)
			log.Infof("found %d logs, iteration count %d", len(logs), logCount)
			// iterate over the logs and update the token holders state
			for _, l := range logs {
				if done, ok := newBlocksMap[l.BlockNumber]; ok && done {
					continue
				}
				newBlocksMap[l.BlockNumber] = true
				// update the token holders state with the log data
				switch th.Type() {
				case CONTRACT_TYPE_ERC20:
					filter := w.contract.(*erc20.ERC20Contract).ERC20ContractFilterer
					logData, err := filter.ParseTransfer(l)
					if err != nil {
						log.Errorf("error parsing log data %s", err)
						continue
					}
					if toCandidate, exists := holdersCandidates[logData.To]; exists {
						holdersCandidates[logData.To] = new(big.Int).Add(toCandidate, logData.Value)
					} else {
						holdersCandidates[logData.To] = logData.Value
					}
					if fromCandidate, exists := holdersCandidates[logData.From]; exists {
						holdersCandidates[logData.From] = new(big.Int).Sub(fromCandidate, logData.Value)
					} else {
						holdersCandidates[logData.From] = new(big.Int).Neg(logData.Value)
					}
				}
			}
			// check if we need to exit because max logs reached for iteration
			if logCount > MAX_SCAN_LOGS_PER_ITERATION {
				return fromBlockNumber, nil
			}
		}
	}
	// delete holder candidates without funds
	newHolders := make([]common.Address, 0)
	for addr, amount := range holdersCandidates {
		if amount.Cmp(big.NewInt(0)) != 0 {
			newHolders = append(newHolders, addr)
		}
	}
	// get current holders who are not candidates
	holdersToCheck := make([]common.Address, 0)
	for _, currentHolder := range th.Holders() {
		if _, exists := holdersCandidates[currentHolder]; !exists {
			holdersToCheck = append(holdersToCheck, currentHolder)
		}
	}
	// check the balances of the current holders that are not in the candidates,
	// and remove these that have not founds
	switch th.Type() {
	case CONTRACT_TYPE_ERC20:
		balanceOf := w.contract.(*erc20.ERC20Contract).BalanceOf
		for _, holder := range holdersToCheck {
			amount, err := balanceOf(&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(lastBlockNumber)}, holder)
			if err != nil {
				return 0, err
			}
			if amount.Cmp(big.NewInt(0)) == 0 {
				th.Del(holder)
			}
		}
	}
	// add the candidate holders to the current holders
	th.Append(newHolders...)
	return lastBlockNumber, nil
}

type TokenData struct {
	Address     common.Address
	Type        ContractType
	Name        string
	Symbol      string
	Decimals    uint8
	TotalSupply *big.Int
}

func (w *Web3) GetTokenData() (*TokenData, error) {
	td := &TokenData{
		Address: w.contractAddress,
		Type:    w.contractType,
	}
	var err error
	if td.Name, err = w.TokenName(); err != nil {
		return nil, fmt.Errorf("unable to get token data: %s", err)
	}

	if td.Symbol, err = w.TokenSymbol(); err != nil {
		return nil, fmt.Errorf("unable to get token data: %s", err)
	}

	if td.Decimals, err = w.TokenDecimals(); err != nil {
		return nil, fmt.Errorf("unable to get token data: %s", err)
	}

	if td.TotalSupply, err = w.TokenTotalSupply(); err != nil {
		return nil, fmt.Errorf("unable to get token data: %s", err)
	}
	return td, nil
}

func (t *TokenData) String() string {
	return fmt.Sprintf(`{"address":%s, "type":%s "name":%s,"symbol":%s,"decimals":%s,"totalSupply":%s}`,
		t.Address, t.Type.String(), t.Name, t.Symbol, string(t.Decimals), t.TotalSupply.String())
}
