package state

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

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
	contractType    TokenType
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
func (w *Web3) Init(ctx context.Context, web3Endpoint string, contractAddress common.Address, contractType TokenType) error {
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
			return nil, fmt.Errorf("wrong number of arguments for AragonWrappedANT balanceOf function")
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

// BlockTimestamp function returns the string timestampt of the provided block
// number. The timestamp will be in RFC3339 format.
func (w *Web3) BlockTimestamp(ctx context.Context, blockNumber uint) (string, error) {
	blockHeader, err := w.client.HeaderByNumber(ctx, new(big.Int).SetInt64(int64(blockNumber)))
	if err != nil {
		return "", err
	}
	timeLayout := "2006-01-02T15:04:05Z07:00"
	return time.Unix(int64(blockHeader.Time), 0).Format(timeLayout), nil
}

// BlockRootHash functions returns the root hash of the provided block number in
// bytes.
func (w *Web3) BlockRootHash(ctx context.Context, blockNumber uint) ([]byte, error) {
	blockHeader, err := w.client.HeaderByNumber(ctx, new(big.Int).SetInt64(int64(blockNumber)))
	if err != nil {
		return nil, err
	}
	return blockHeader.Root.Bytes(), nil
}

// UpdateTokenHolders function checks the transfer logs of the given contract
// (in the TokenHolders struct) from the given block number. It gets all
// addresses (candidates to holders) and their balances from the given block
// number to the latest block number and submit the results using
// Web3.submitTokenHolders function.
func (w *Web3) UpdateTokenHolders(ctx context.Context, th *TokenHolders) (uint64, error) {
	// fetch the last block header
	lastBlockHeader, err := w.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	log.Debugf("last block number: %d", lastBlockHeader.Number.Uint64())
	// check if there are new blocks to scan
	lastBlockNumber := lastBlockHeader.Number.Uint64()
	fromBlockNumber := th.LastBlock() + 1
	if fromBlockNumber >= lastBlockNumber {
		return fromBlockNumber, fmt.Errorf("no new blocks to scan for %s", th.Address().String())
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
	holdersCandidates := HoldersCandidates{}
	for fromBlockNumber < lastBlockNumber {
		select {
		// check if we need to close due context signal
		case <-ctx.Done():
			log.Warnf("scan graceful canceled by context")
			th.BlockDone(fromBlockNumber)
			return fromBlockNumber, w.submitTokenHolders(th, holdersCandidates, th.LastBlock())
		default:
			startTime := time.Now()
			log.Infof("analyzing blocks from %d to %d [%d%%]", fromBlockNumber,
				fromBlockNumber+blocks, (fromBlockNumber*100)/lastBlockNumber)
			// get transfer logs for the following n blocks
			logs, err := w.getTransferLogs(th, fromBlockNumber, blocks)
			if err != nil {
				// if we have too much results, decrease the blocks to scan
				if strings.Contains(err.Error(), "query returned more than") {
					blocks /= 2
					log.Warnf("too much results on query, decreasing blocks to %d", blocks)
					continue
				}
				return 0, err
			}
			fromBlockNumber += blocks
			// after updating the starter block number for the next iteration,
			// if there are no logs, mark the starter block as done and proceed
			// to the next iteration.
			if len(logs) == 0 {
				th.BlockDone(fromBlockNumber)
				continue
			}
			logCount += len(logs)
			log.Infof("found %d logs, iteration count %d", len(logs), logCount)
			blocksToSave := make(map[uint64]bool)
			// iterate over the logs and update the token holders state
			for _, currentLog := range logs {
				// If the current log block number is already scanned proceed to
				// the next iteration.
				if _, ok := newBlocksMap[currentLog.BlockNumber]; !ok {
					if th.HasBlock(currentLog.BlockNumber) {
						log.Debugf("found already processed block %d", fromBlockNumber)
						continue
					}
				}
				// update the holders candidates with the current log
				holdersCandidates = w.updateHolderCandidates(holdersCandidates, th.Type(), currentLog)
				blocksToSave[currentLog.BlockNumber] = true
				newBlocksMap[currentLog.BlockNumber] = true
			}
			log.Debugf("saved %d blocks at %.2f blocks/second", len(blocksToSave),
				1000*float32(len(blocksToSave))/float32(time.Since(startTime).Milliseconds()))
			// check if we need to exit because max logs reached for iteration
			if len(holdersCandidates) > MAX_NEW_HOLDER_CANDIDATES_PER_ITERATION {
				log.Debug("MAX_NEW_HOLDER_CANDIDATES_PER_ITERATION limit reached... stop scanning")
				th.BlockDone(fromBlockNumber)
				return fromBlockNumber, w.submitTokenHolders(th, holdersCandidates, fromBlockNumber)
			}
			if logCount > MAX_SCAN_LOGS_PER_ITERATION {
				log.Debug("MAX_SCAN_LOGS_PER_ITERATION limit reached... stop scanning")
				th.BlockDone(fromBlockNumber)
				return fromBlockNumber, w.submitTokenHolders(th, holdersCandidates, fromBlockNumber)
			}
		}
	}
	th.BlockDone(lastBlockNumber)
	return lastBlockNumber, w.submitTokenHolders(th, holdersCandidates, lastBlockNumber)
}

// getTransferLogs function queries to the web3 endpoint for the transfer logs
// of the token provided, that are included in the range of blocks defined by
// the from block number provided to the following number of blocks given.
func (w *Web3) getTransferLogs(th *TokenHolders, fromBlock, nblocks uint64) ([]types.Log, error) {
	// create the filter query
	query := eth.FilterQuery{
		Addresses: []common.Address{th.Address()},
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(fromBlock + nblocks)),
	}
	// set the topics to filter depending on the contract type
	switch th.Type() {
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
		return nil, fmt.Errorf("unknown contract type %d", th.Type())
	}
	// execute the filter query
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return w.client.FilterLogs(ctx, query)
}

// updateHolderCandidates function updates the given holder candidates with the
// values of the log provided. It uses uses functions appropriate to the token
// type indicated.
func (w *Web3) updateHolderCandidates(hc HoldersCandidates, ttype TokenType, currentLog types.Log) HoldersCandidates {
	// update the token holders state with the log data
	switch ttype {
	case CONTRACT_TYPE_ERC20:
		filter := w.contract.(*erc20.ERC20Contract).ERC20ContractFilterer
		logData, err := filter.ParseTransfer(currentLog)
		if err != nil {
			log.Errorf("error parsing log data %s", err)
			return hc
		}
		if toBalance, exists := hc[logData.To]; exists {
			hc[logData.To] = new(big.Int).Add(toBalance, logData.Value)
		} else {
			hc[logData.To] = logData.Value
		}
		if fromBalance, exists := hc[logData.From]; exists {
			hc[logData.From] = new(big.Int).Sub(fromBalance, logData.Value)
		} else {
			hc[logData.From] = new(big.Int).Neg(logData.Value)
		}
	case CONTRACT_TYPE_ERC777: // stores the total count per address, not all identifiers
		filter := w.contract.(*erc777.ERC777Contract).ERC777ContractFilterer
		logData, err := filter.ParseTransfer(currentLog)
		if err != nil {
			log.Errorf("error parsing log data %s", err)
			return hc
		}
		if toBalance, exists := hc[logData.To]; exists {
			hc[logData.To] = new(big.Int).Add(toBalance, big.NewInt(1))
		} else {
			hc[logData.To] = big.NewInt(1)
		}
		if fromBalance, exists := hc[logData.From]; exists {
			hc[logData.From] = new(big.Int).Sub(fromBalance, big.NewInt(1))
		} else {
			hc[logData.From] = new(big.Int).Neg(big.NewInt(1))
		}
	case CONTRACT_TYPE_ERC721: // stores the total count per address, not all identifiers
		filter := w.contract.(*erc721.ERC721Contract).ERC721ContractFilterer
		logData, err := filter.ParseTransfer(currentLog)
		if err != nil {
			log.Errorf("error parsing log data %s", err)
			return hc
		}
		if toBalance, exists := hc[logData.To]; exists {
			hc[logData.To] = new(big.Int).Add(toBalance, big.NewInt(1))
		} else {
			hc[logData.To] = big.NewInt(1)
		}
		if fromBalance, exists := hc[logData.From]; exists {
			hc[logData.From] = new(big.Int).Sub(fromBalance, big.NewInt(1))
		} else {
			hc[logData.From] = new(big.Int).Neg(big.NewInt(1))
		}
	case CONTRACT_TYPE_CUSTOM_NATION3_VENATION:
		// This token contract is a bit special, token balances
		// are updated every block based on the contract state.
		switch currentLog.Topics[0] {
		case common.HexToHash(LOG_TOPIC_VENATION_DEPOSIT):
			provider := common.HexToAddress(currentLog.Topics[1].Hex())
			value := big.NewInt(0).SetBytes(currentLog.Data[:32])
			if toBalance, exists := hc[provider]; exists {
				hc[provider] = new(big.Int).Add(toBalance, value)
			} else {
				hc[provider] = value
			}
		case common.HexToHash(LOG_TOPIC_VENATION_WITHDRAW):
			provider := common.HexToAddress(currentLog.Topics[1].Hex())
			value := big.NewInt(0).SetBytes(currentLog.Data[:32])
			if fromBalance, exists := hc[provider]; exists {
				hc[provider] = new(big.Int).Sub(fromBalance, value)
			} else {
				hc[provider] = new(big.Int).Neg(value)
			}
		}
	case CONTRACT_TYPE_CUSTOM_ARAGON_WANT:
		filter := w.contract.(*want.AragonWrappedANTTokenContract).AragonWrappedANTTokenContractFilterer
		switch currentLog.Topics[0] {
		case common.HexToHash(LOG_TOPIC_WANT_DEPOSIT):
			logData, err := filter.ParseDeposit(currentLog)
			if err != nil {
				log.Errorf("error parsing log data %s", err)
				return hc
			}
			if toBalance, exists := hc[logData.Entity]; exists {
				hc[logData.Entity] = new(big.Int).Add(toBalance, logData.Amount)
			} else {
				hc[logData.Entity] = logData.Amount
			}
		case common.HexToHash(LOG_TOPIC_WANT_WITHDRAWAL):
			logData, err := filter.ParseWithdrawal(currentLog)
			if err != nil {
				log.Errorf("error parsing log data %s", err)
				return hc
			}
			if fromBalance, exists := hc[logData.Entity]; exists {
				hc[logData.Entity] = new(big.Int).Sub(fromBalance, logData.Amount)
			} else {
				hc[logData.Entity] = new(big.Int).Neg(logData.Amount)
			}
		}
	}
	return hc
}

// submitTokenHolders function checks each candidate to token holder provided,
// and removes any with a zero balance before store them. It also checks the
// balances of the current holders, deleting those with no funds.
func (w *Web3) submitTokenHolders(th *TokenHolders, candidates HoldersCandidates, blockNumber uint64) error {
	// remove null address from candidates
	delete(candidates, common.HexToAddress(NULL_ADDRESS))
	// delete holder candidates without funds
	for addr, balance := range candidates {
		if balance.Cmp(big.NewInt(0)) != 0 {
			th.Append(addr, balance)
		}
	}
	return nil
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

// GetContractCreationBlock function calculates the block number where the
// current was created. It tries to calculate it using the first block (0) and
// the current last block.
func (w *Web3) GetContractCreationBlock(ctx context.Context) (uint64, error) {
	lastBlockHeader, err := w.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	log.Debugf("last block number: %d", lastBlockHeader.Number.Uint64())
	return w.getCreationBlock(ctx, 0, lastBlockHeader.Number.Uint64())
}

// getCreationBlock function finds the block number of a contract between the
// bounds provided as start and end blocks.
func (w *Web3) getCreationBlock(ctx context.Context, start, end uint64) (uint64, error) {
	// if both block numbers are equal, return its value as birthblock
	if start == end {
		return start, nil
	}
	// find the middle block between start and end blocks and get the contract
	// code at this block
	midBlock := (start + end) / 2
	codeLen, err := w.SourceCodeLenAt(ctx, midBlock)
	if err != nil {
		return 0, err
	}
	// if any code is found, keep trying with the lower half of blocks until
	// find the first. if not, keep trying with the upper half
	if codeLen > 2 {
		return w.getCreationBlock(ctx, start, midBlock)
	} else {
		return w.getCreationBlock(ctx, midBlock+1, end)
	}
}

// SourceCodeLenAt function returns the length of the current contratct bytecode
// at the block number provided.
func (w *Web3) SourceCodeLenAt(ctx context.Context, atBlockNumber uint64) (int, error) {
	blockNumber := new(big.Int).SetUint64(atBlockNumber)
	sourceCode, err := w.client.CodeAt(ctx, w.contractAddress, blockNumber)
	return len(sourceCode), err
}
