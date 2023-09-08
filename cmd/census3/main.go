package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	flag "github.com/spf13/pflag"
	"github.com/vocdoni/census3/api"
	"github.com/vocdoni/census3/db"
	"github.com/vocdoni/census3/service"
	"github.com/vocdoni/census3/state"
	"go.vocdoni.io/dvote/log"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	home += "/.census3"
	dataDir := flag.String("dataDir", home, "data directory for persistent storage")
	logLevel := flag.String("logLevel", "info", "log level (debug, info, warn, error)")
	port := flag.Int("port", 7788, "HTTP port for the API")
	connectKey := flag.String("connectKey", "", "connect group key for IPFS connect")
	listOfWeb3Providers := flag.String("web3Providers", "", "the list of URL's of available web3 providers (separated with commas)")
	flag.Parse()
	log.Init(*logLevel, "stdout", nil)

	database, err := db.Init(*dataDir)
	if err != nil {
		log.Fatal(err)
	}

	web3Providers := strings.Split(*listOfWeb3Providers, ",")
	w3p, err := state.CheckWeb3Providers(web3Providers)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(w3p)

	// Start the holder scanner
	hc, err := service.NewHoldersScanner(database, w3p)
	if err != nil {
		log.Fatal(err)
	}

	// Start the API
	err = api.Init(database, api.Census3APIConf{
		Hostname:      "0.0.0.0",
		Port:          *port,
		DataDir:       *dataDir,
		Web3Providers: w3p,
		GroupKey:      *connectKey,
	})
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	go hc.Start(ctx)

	// Wait for SIGTERM
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Warnf("received SIGTERM, exiting at %s", time.Now().Format(time.RFC850))
	cancel()
	log.Infof("waiting for routines to end gracefully...")
	// closing database
	go func() {
		if err := database.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	time.Sleep(5 * time.Second)
	os.Exit(0)
}
