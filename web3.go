package tokenstate

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/vocdoni/go-dvote/log"
	"gitlab.com/vocdoni/go-dvote/util"
)

var BlocksToScan = 10000

// Web3 holds a reference to a go-ethereum client,
// to an ERC20 like contract and to an ENS.
// It is expected for the ERC20 contract to implement the standard
// optional ERC20 functions: {name, symbol, decimals, totalSupply}
type Web3 struct {
	client    *ethclient.Client
	token     *ERC20BaseContractCaller
	networkID *big.Int
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
	log.Infof("found network %s", w.networkID.String())
	// load token contract
	c, err := hex.DecodeString(util.TrimHex(contractAddress))
	if err != nil {
		return err
	}
	caddr := common.Address{}
	caddr.SetBytes(c)
	if w.token, err = NewERC20BaseContractCaller(caddr, w.client); err != nil {
		return err
	}
	log.Infof("loaded token contract %s", caddr.String())
	return nil
}

func (w *Web3) GetTokenData() (*TokenData, error) {
	td := &TokenData{}
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

type TokenData struct {
	Name        string
	Symbol      string
	Decimals    uint8
	TotalSupply *big.Int
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
func (w *Web3) ScanERC20Holders(ts *TokenState, fromBlock uint64, contract string, decimals uint8) error {
	ctx := context.Background()
	thash := common.Hash{}
	tbytes, err := hex.DecodeString("ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	if err != nil {
		return err
	}
	thash.SetBytes(tbytes)

	from := common.Hash{}
	to := common.Hash{}
	amount := big.NewFloat(0)

	c, err := hex.DecodeString(util.TrimHex(contract))
	if err != nil {
		return err
	}
	caddr := common.Address{}
	caddr.SetBytes(c)

	log.Infof("scaning logs for contract %s", caddr.String())

	header, err := w.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}
	currentblock := header.Number.Uint64()
	var blocks uint64
	blocks = uint64(BlocksToScan)
	newBlocks := make(map[uint64]bool)
	for fromBlock < currentblock {
		log.Infof("analyzing blocks from %d to %d [%d%%]", fromBlock, fromBlock+blocks, (fromBlock*100)/currentblock)
		query := eth.FilterQuery{
			Addresses: []common.Address{caddr},
			Topics:    [][]common.Hash{{thash}},
			FromBlock: big.NewInt(int64(fromBlock)),
			ToBlock:   big.NewInt(int64(fromBlock + blocks)),
		}
		ctx2, cancel := context.WithTimeout(ctx, time.Second*10)
		logs, err := w.client.FilterLogs(ctx2, query)
		if err != nil {
			if strings.Contains(err.Error(), "query returned more than") {
				blocks = blocks / 2
				log.Warnf("too much results on query, decreasing blocks to %d", blocks)
				continue
			} else {
				return err
			}
		}
		defer cancel()

		fromBlock += blocks
		if len(logs) > 0 {
			log.Infof("found %d logs...", len(logs))
		}
		blocksToSave := make(map[uint64]bool)
		for _, l := range logs {
			if _, ok := newBlocks[l.BlockNumber]; !ok {
				if ts.HasBlock(l.BlockNumber) {
					log.Infof("found already processe block %d", fromBlock)
					continue
				}
			}
			blocksToSave[l.BlockNumber] = true
			newBlocks[l.BlockNumber] = true
			from = l.Topics[1]
			to = l.Topics[2]
			if _, ok := amount.SetString(fmt.Sprintf("0x%x", l.Data)); !ok {
				log.Warnf("cannot parse amount")
				continue
			}
			amount.Mul(amount, big.NewFloat(float64(decimals)))
			fromstr := from.String()
			fromstr = fromstr[len(fromstr)-40:]
			tostr := to.String()
			tostr = tostr[len(tostr)-40:]

			ts.Add(tostr, amount)
			ts.Sub(fromstr, amount)
		}
		for k := range blocksToSave {
			ts.Save(k)
		}
	}
	return nil
}
