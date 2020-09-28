package tokenstate

import (
	"context"
	"encoding/hex"
	"fmt"
	"math"
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

// ScanERC20 scans the Ethereum network and updates the token holders state
func ScanERC20(ts *TokenState, web3url string, fromBlock uint64, contract string, decimals int) error {
	ctx := context.Background()
	client, err := ethclient.Dial(web3url)
	if err != nil {
		log.Fatal(err)
	}
	cid, err := client.ChainID(ctx)
	if err != nil {
		return err
	}
	log.Infof("found network %s", cid)

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

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}
	currentblock := header.Number.Uint64()
	decimalsMul := math.Pow(10, float64((decimals * -1)))
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
		logs, err := client.FilterLogs(ctx2, query)
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
			amount.Mul(amount, big.NewFloat(decimalsMul))
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
