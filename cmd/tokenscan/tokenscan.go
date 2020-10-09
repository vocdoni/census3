package main

import (
	"context"
	"flag"
	"os"

	state "github.com/vocdoni/tokenstate/tokenstate"
	"gitlab.com/vocdoni/go-dvote/log"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	home += "/.tokenscan"
	contract := flag.String("contract", "", "token contract address")
	url := flag.String("url", "", "ethereum RPC url")
	fromblock := flag.Int64("from", 0, "from block number")
	//blocks := flag.Int64("blocks", 10000, "number of blocks to scan")
	dataDir := flag.String("dataDir", home, "data directory for persistent storage")
	flag.Parse()
	log.Init("info", "stdout")

	var ts state.TokenState
	if err = ts.Init(*dataDir, *contract); err != nil {
		log.Fatal(err)
	}
	defer ts.Close()

	var w3 state.Web3
	if err := w3.Init(context.Background(), *url, *contract); err != nil {
		log.Fatal(err)
	}

	// scan token
	if err := w3.ScanERC20Holders(&ts, uint64(*fromblock), *contract); err != nil {
		log.Fatal(err)
	}

	log.Info("Balances for block")
	totals := ts.List(0)
	for addr, amount := range totals {
		log.Infof("0x%s %s\n", addr, amount.String())
	}
	log.Infof("Total: %s\n", totals["total"].String())
	log.Infof("Balance: %s\n", totals["balance"].String())
}
