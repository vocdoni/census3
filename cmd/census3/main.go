package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/vocdoni/census3/api"
	"github.com/vocdoni/census3/db"
	"github.com/vocdoni/census3/service"
	"github.com/vocdoni/census3/state"
	"go.vocdoni.io/dvote/log"
)

type Census3Config struct {
	dataDir, logLevel, connectKey string
	listOfWeb3Providers           []string
	port                          int
}

func main() {
	// init service config
	config := Census3Config{}
	// get home directory to create the data directory for persistent storage
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	home += "/.census3"
	// parse flags
	flag.StringVar(&config.dataDir, "dataDir", home, "data directory for persistent storage")
	flag.StringVar(&config.logLevel, "logLevel", "info", "log level (debug, info, warn, error)")
	flag.IntVar(&config.port, "port", 7788, "HTTP port for the API")
	flag.StringVar(&config.connectKey, "connectKey", "", "connect group key for IPFS connect")
	var strWeb3Providers string
	flag.StringVar(&strWeb3Providers, "web3Providers", "", "the list of URL's of available web3 providers")
	flag.Parse()
	// init viper to read config file
	pviper := viper.New()
	pviper.SetConfigName("census3")
	pviper.SetConfigType("yml")
	pviper.SetEnvPrefix("CENSUS3")
	pviper.AutomaticEnv()
	pviper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// bind flags to viper
	if err := pviper.BindPFlag("dataDir", flag.Lookup("dataDir")); err != nil {
		panic(err)
	}
	config.dataDir = pviper.GetString("dataDir")
	// read config file
	pviper.AddConfigPath(config.dataDir)
	_ = pviper.ReadInConfig()

	if err := pviper.BindPFlag("logLevel", flag.Lookup("logLevel")); err != nil {
		panic(err)
	}
	config.logLevel = pviper.GetString("logLevel")
	if err := pviper.BindPFlag("port", flag.Lookup("port")); err != nil {
		panic(err)
	}
	config.port = pviper.GetInt("port")
	if err := pviper.BindPFlag("connectKey", flag.Lookup("connectKey")); err != nil {
		panic(err)
	}
	config.connectKey = pviper.GetString("connectKey")
	if err := pviper.BindPFlag("web3Providers", flag.Lookup("web3Providers")); err != nil {
		panic(err)
	}
	config.listOfWeb3Providers = strings.Split(pviper.GetString("web3Providers"), ",")
	// init logger
	log.Init(config.logLevel, "stdout", nil)
	// check if the web3 providers are defined
	if len(config.listOfWeb3Providers) == 0 {
		log.Fatal("no web3 providers defined")
	}
	// check if the web3 providers are valid
	w3p, err := state.CheckWeb3Providers(config.listOfWeb3Providers)
	if err != nil {
		log.Fatal(err)
	}
	// init the database
	database, err := db.Init(config.dataDir)
	if err != nil {
		log.Fatal(err)
	}
	// start the holder scanner
	hc, err := service.NewHoldersScanner(database, w3p)
	if err != nil {
		log.Fatal(err)
	}
	// start the API
	err = api.Init(database, api.Census3APIConf{
		Hostname:      "0.0.0.0",
		Port:          config.port,
		DataDir:       config.dataDir,
		Web3Providers: w3p,
		GroupKey:      config.connectKey,
	})
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	go hc.Start(ctx)

	// wait for SIGTERM
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
