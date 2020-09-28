package main

import (
	"flag"
	"fmt"
	"os"

	state "github.com/vocdoni/tokenstate"
	"gitlab.com/vocdoni/go-dvote/log"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	home += "/.tokenscan"
	contract := flag.String("contract", "", "token contract address")
	url := flag.String("url", "http://127.0.0.1:8545", "ethereum RPC url")
	fromblock := flag.Int64("from", 0, "from block number")
	//blocks := flag.Int64("blocks", 10000, "number of blocks to scan")
	decimals := flag.Int("decimals", 18, "numer of decimals for token")
	dataDir := flag.String("dataDir", home, "data directory for persistent storage")
	flag.Parse()
	log.Init("info", "stdout")

	var ts state.TokenState
	if err = ts.Init(*dataDir, *contract); err != nil {
		log.Fatal(err)
	}
	defer ts.Close()

	if err := state.ScanERC20(&ts, *url, uint64(*fromblock), *contract, *decimals); err != nil {
		log.Fatal(err)
	}

	log.Info("Balances for block")
	totals := ts.List(0)
	for addr, amount := range totals {
		fmt.Printf("0x%s %s\n", addr, amount.String())
	}
	fmt.Printf("Total: %s\n", totals["total"].String())
	fmt.Printf("Balance: %s\n", totals["balance"].String())
}
