package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	flag "github.com/spf13/pflag"
	"github.com/vocdoni/tokenstate/api"
	"github.com/vocdoni/tokenstate/service"
	"go.vocdoni.io/dvote/crypto/ethereum"
	"go.vocdoni.io/dvote/log"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	home += "/.tokenscan"
	url := flag.String("url", "", "ethereum web3 url")
	dataDir := flag.String("dataDir", home, "data directory for persistent storage")
	logLevel := flag.String("logLevel", "info", "log level (debug, info, warn, error)")
	port := flag.Int32("port", 7788, "HTTP port for the API")
	op := flag.String("op", "", "operation to perform (erc20, nation3)")
	flag.Parse()
	log.Init(*logLevel, "stdout")

	signer := ethereum.SignKeys{}
	if err := signer.Generate(); err != nil {
		log.Fatal(err)
	}

	if op == nil {
		log.Fatal("no operation specified")
	}
	if *op == "" {
		log.Fatal("no operation specified")
	}

	switch *op {
	case "erc20":
		sc, err := service.NewScanner(*dataDir, *url)
		if err != nil {
			log.Fatal(err)
		}
		api.Init("0.0.0.0", *port, &signer, sc)
		ctx, close := context.WithCancel(context.Background())
		defer close()
		go sc.Start(ctx)
	case "nation3":
		sc, err := service.NewNation3Scanner(*dataDir, *url)
		if err != nil {
			log.Fatal(err)
		}
		ctx, close := context.WithCancel(context.Background())
		defer close()
		go sc.Start(ctx)
	}

	// Wait for SIGTERM
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Warnf("received SIGTERM, exiting at %s", time.Now().Format(time.RFC850))
	log.Infof("waiting for routines to end gracefuly...")
	time.Sleep(12 * time.Second)
	os.Exit(0)
}
