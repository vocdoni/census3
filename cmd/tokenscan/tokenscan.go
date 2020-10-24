package main

import (
	"flag"
	"os"

	"github.com/vocdoni/tokenstate/api"
	"github.com/vocdoni/tokenstate/service"
	"gitlab.com/vocdoni/go-dvote/crypto/ethereum"
	"gitlab.com/vocdoni/go-dvote/log"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	home += "/.tokenscan"
	//contract := flag.String("contract", "", "token contract address")
	url := flag.String("url", "", "ethereum web3 url")
	//fromblock := flag.Int64("from", 0, "from block number")
	//blocks := flag.Int64("blocks", 10000, "number of blocks to scan")
	dataDir := flag.String("dataDir", home, "data directory for persistent storage")
	flag.Parse()
	log.Init("debug", "stdout")

	/*	var ts state.TokenState
		if err = ts.Init(*dataDir, *contract); err != nil {
			log.Fatal(err)
		}
		defer ts.Close()
	*/
	signer := ethereum.SignKeys{}
	if err := signer.Generate(); err != nil {
		log.Fatal(err)
	}

	sc, err := service.NewScanner(*dataDir + "/tokens")
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()
	api.Init("0.0.0.0", 7788, &signer, sc)
	sc.Start(*url)

	/*	var w3 state.Web3
		if err := w3.Init(context.Background(), *url, *contract); err != nil {
			log.Fatal(err)
		}

		// scan token
		if _, err := w3.ScanERC20Holders(&ts, uint64(*fromblock), *contract); err != nil {
			log.Fatal(err)
		}

		log.Info("Balances for block")
		totals := ts.List(0)
		for addr, amount := range totals {
			log.Infof("0x%s %s\n", addr, amount.String())
		}
		log.Infof("Total: %s\n", totals["total"].String())
		log.Infof("Balance: %s\n", totals["balance"].String())
	*/
}
