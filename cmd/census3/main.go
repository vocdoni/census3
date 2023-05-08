package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	flag "github.com/spf13/pflag"
	"github.com/vocdoni/census3/api"
	"github.com/vocdoni/census3/db"
	"github.com/vocdoni/census3/service"
	"go.vocdoni.io/dvote/log"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	home += "/.census3"
	url := flag.String("url", "", "ethereum web3 url")
	dataDir := flag.String("dataDir", home, "data directory for persistent storage")
	logLevel := flag.String("logLevel", "info", "log level (debug, info, warn, error)")
	port := flag.Int("port", 7788, "HTTP port for the API")
	flag.Parse()
	log.Init(*logLevel, "stdout")

	db, q, err := db.Init(*dataDir)
	if err != nil {
		log.Fatal(err)
	}

	// Start the holder scanner
	hc, err := service.NewHoldersScanner(db, q, *url)
	if err != nil {
		log.Fatal(err)
	}

	// Start the API
	err = api.Init(db, q, api.Census3APIConf{
		Hostname: "0.0.0.0",
		Port:     *port,
		DataDir:  *dataDir,
		Web3URI:  *url,
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
	time.Sleep(5 * time.Second)
	os.Exit(0)
}
