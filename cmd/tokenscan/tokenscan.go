package main

import (
	"context"
	"flag"
	"os"

	state "github.com/vocdoni/tokenstate"
	"github.com/vocdoni/tokenstate/entitybridge"
	"gitlab.com/vocdoni/go-dvote/log"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	home += "/.tokenscan"
	contract := flag.String("contract", "0x106c8eBaD6D9A71c962Da4088721221de9BD4fB7", "token contract address")
	url := flag.String("url", "https://rpc.xdaichain.com/", "ethereum RPC url")
	gwUrl := flag.String("gwUrl", "", "gateway api endpoint")
	signer := flag.String("signer", "", "ethereum sign keys private key")
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

	b := entitybridge.NewEntityBridgeService()
	if err := b.Init(context.Background(), *url, *gwUrl, *contract, *signer); err != nil {
		log.Infof("service initialization error: %s\n", err)
		return
	}

	// create token entity
	res, err := b.CreateEntityMetadata()
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("ipfs file URL: %s", res)

	// scan token
	if err := b.TokenState.ScanERC20Holders(&ts, uint64(*fromblock), *contract); err != nil {
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
