package contractstate

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/vocdoni/tokenstate/contracts"

	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.vocdoni.io/dvote/log"
	"go.vocdoni.io/dvote/util"
)

const (
	erc20LogTopicTransfer     = "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	maxScanBlocksPerIteration = 2000000
	maxScanLogsPerIteration   = 50000
	blocksToScanAtOnce        = 10000
)

// Web3 holds a reference to a go-ethereum client,
// to an ERC20 like contract and to an ENS.
// It is expected for the ERC20 contract to implement the standard
// optional ERC20 functions: {name, symbol, decimals, totalSupply}
type Web3 struct {
	client    *ethclient.Client
	token     *contracts.ContractsCaller
	tokenAddr string
	networkID *big.Int
	close     chan (bool)
}

// Init creates and client connection and connects to an ERC20 contract given its address
func (w *Web3) Init(ctx context.Context, web3Endpoint, contractAddress string) error {
	var err error
	// connect to ethereum endpoint
	w.client, err = ethclient.Dial(web3Endpoint)
	if err != nil {
		log.Fatal(err)
	}
	w.networkID, err = w.client.ChainID(ctx)
	if err != nil {
		return err
	}
	log.Debugf("found ethereum network id %s", w.networkID.String())
	// load token contract
	c, err := hex.DecodeString(util.TrimHex(contractAddress))
	if err != nil {
		return err
	}
	caddr := common.Address{}
	caddr.SetBytes(c)
	if w.token, err = contracts.NewContractsCaller(caddr, w.client); err != nil {
		return err
	}
	w.tokenAddr = contractAddress
	log.Infof("loaded token contract %s", caddr.String())
	return nil
}

func (w *Web3) Close() {
	w.close <- true
}

func (w *Web3) GetTokenData() (*TokenData, error) {
	td := &TokenData{Address: w.tokenAddr}
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

func (w *Web3) Balance(ctx context.Context, address string) (*big.Int, error) {
	return w.client.BalanceAt(ctx, common.HexToAddress(address), nil)
}

type TokenData struct {
	Address     string   `json:"address"`
	Name        string   `json:"name"`
	Symbol      string   `json:"symbol"`
	Decimals    uint8    `json:"decimals"`
	TotalSupply *big.Int `json:"totalSupply,omitempty"`
}

func (t *TokenData) String() string {
	return fmt.Sprintf(`{"name":%s,"symbol":%s,"decimals":%s,"totalSupply":%s}`,
		t.Name, t.Symbol, string(t.Decimals), t.TotalSupply.String())
}

// TokenName wraps the name() function contract call
func (w *Web3) TokenName() (string, error) {
	return w.token.Name(nil)
}

// TokenSymbol wraps the symbol() function contract call
func (w *Web3) TokenSymbol() (string, error) {
	return w.token.Symbol(nil)
}

// TokenDecimals wraps the decimals() function contract call
func (w *Web3) TokenDecimals() (uint8, error) {
	return w.token.Decimals(nil)
}

// TokenTotalSupply wraps the totalSupply function contract call
func (w *Web3) TokenTotalSupply() (*big.Int, error) {
	return w.token.TotalSupply(nil)
}

// ScanERC20Holders scans the Ethereum network and updates the token holders state
func (w *Web3) ScanERC20Holders(ctx context.Context, ts *ContractState,
	fromBlock uint64, contract string) (uint64, error) {
	thash := common.Hash{}
	tbytes, err := hex.DecodeString(erc20LogTopicTransfer)
	if err != nil {
		return 0, err
	}
	thash.SetBytes(tbytes)

	from := common.Hash{}
	to := common.Hash{}
	amount := big.NewInt(0)

	c, err := hex.DecodeString(util.TrimHex(contract))
	if err != nil {
		return 0, err
	}
	caddr := common.Address{}
	caddr.SetBytes(c)

	header, err := w.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	toBlock := header.Number.Uint64()
	if toBlock-fromBlock > maxScanBlocksPerIteration {
		toBlock = fromBlock + maxScanBlocksPerIteration
	}
	var blocks uint64
	var logCount int
	blocks = uint64(blocksToScanAtOnce)
	newBlocks := make(map[uint64]bool)
	log.Infof("start scan iteration for %s from block %d to %d (%d)",
		contract, fromBlock, toBlock, toBlock-fromBlock)

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
			logs, err := w.client.FilterLogs(ctx2, query)
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
