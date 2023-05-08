package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vocdoni/census3/api"
	erc20 "github.com/vocdoni/census3/contracts/erc/erc20"
)

func main() {
	startTime := time.Now()
	// 1. Get token address to check and url to web3
	tokenAddress := flag.String("addr", "", "")
	url := flag.String("url", "", "")
	census3 := flag.String("api", "http://localhost:7788", "")
	flag.Parse()
	// 2. Init web3
	web3, err := ethclient.Dial(*url)
	if err != nil {
		log.Fatal(err)
	}
	contract, err := erc20.NewERC20Contract(common.HexToAddress(*tokenAddress), web3)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	lastBlockHeader, err := web3.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	lastBlockNumber := lastBlockHeader.Number
	// 3. Call to local API to get holders candidates and its balances
	holdersEndpoint := fmt.Sprintf("%s/api/debug/token/%s/holders", *census3, *tokenAddress)
	resp, err := http.Get(holdersEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	holders := api.TokenHoldersResponse{}
	if err := json.Unmarshal(body, &holders); err != nil {
		log.Fatal(err)
	}
	// 4. Iterate over holders address to get current balances and check with the
	// received in the previous step. Check also if any address is a contract.
	i := 0
	contractAddresses := 0
	sameBalance := 0
	distinctBalance := 0
	noBalance := 0
	for addr, balance := range holders.Holders {
		log.Printf("%d/%d\n", i, len(holders.Holders))
		// check balance desviations
		currentBalance, err := contract.BalanceOf(nil, common.HexToAddress(addr))
		if err != nil {
			log.Fatal(err)
		}
		if cBalance := currentBalance.String(); cBalance == "0" {
			noBalance++
		} else if cBalance == balance {
			sameBalance++
		} else {
			distinctBalance++
		}
		// check if is a contract
		sourceCode, err := web3.CodeAt(ctx, common.HexToAddress(addr), lastBlockNumber)
		if err != nil {
			log.Fatal(err)
		}
		if len(sourceCode) > 2 {
			contractAddresses++
		}
		i++
	}
	// 5. Print results
	elapsed := time.Since(startTime)
	log.Println("took", elapsed)
	log.Println("same balance", sameBalance)
	log.Println("distinct balance", distinctBalance)
	log.Println("no balance", noBalance)
	log.Println("contracts", contractAddresses)
}
