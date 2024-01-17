package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/VictoriaMetrics/metrics"
	"github.com/google/uuid"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/vocdoni/census3/api"
	"github.com/vocdoni/census3/db"
	"github.com/vocdoni/census3/internal"
	"github.com/vocdoni/census3/service"
	"github.com/vocdoni/census3/service/providers"
	"github.com/vocdoni/census3/service/providers/poap"
	"github.com/vocdoni/census3/service/providers/web3"
	"go.vocdoni.io/dvote/log"
)

type Census3Config struct {
	dataDir, logLevel, connectKey  string
	listOfWeb3Providers            []string
	port                           int
	poapAPIEndpoint, poapAuthToken string
	scannerCoolDown                time.Duration
	adminToken                     string
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
	flag.StringVar(&config.poapAPIEndpoint, "poapAPIEndpoint", "", "POAP API access token")
	flag.StringVar(&config.poapAuthToken, "poapAuthToken", "", "POAP API access token")
	var strWeb3Providers string
	flag.StringVar(&strWeb3Providers, "web3Providers", "", "the list of URL's of available web3 providers")
	flag.DurationVar(&config.scannerCoolDown, "scannerCoolDown", 120*time.Second, "the time to wait before next scanner iteration")
	flag.StringVar(&config.adminToken, "adminToken", "", "the admin UUID token for the API")
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
	if err := pviper.BindPFlag("poapAPIEndpoint", flag.Lookup("poapAPIEndpoint")); err != nil {
		panic(err)
	}
	config.poapAPIEndpoint = pviper.GetString("poapAPIEndpoint")
	if err := pviper.BindPFlag("poapAuthToken", flag.Lookup("poapAuthToken")); err != nil {
		panic(err)
	}
	config.poapAuthToken = pviper.GetString("poapAuthToken")
	if err := pviper.BindPFlag("web3Providers", flag.Lookup("web3Providers")); err != nil {
		panic(err)
	}
	config.listOfWeb3Providers = strings.Split(pviper.GetString("web3Providers"), ",")
	if err := pviper.BindPFlag("scannerCoolDown", flag.Lookup("scannerCoolDown")); err != nil {
		panic(err)
	}
	config.scannerCoolDown = pviper.GetDuration("scannerCoolDown")
	if err := pviper.BindPFlag("adminToken", flag.Lookup("adminToken")); err != nil {
		panic(err)
	}
	config.adminToken = pviper.GetString("adminToken")
	// init logger
	log.Init(config.logLevel, "stdout", nil)
	// check if the web3 providers are defined
	if len(config.listOfWeb3Providers) == 0 {
		log.Fatal("no web3 providers defined")
	}
	// check if the web3 providers are valid
	w3p, err := web3.InitNetworkEndpoints(config.listOfWeb3Providers)
	if err != nil {
		log.Fatal(err)
	}
	// init the database
	database, err := db.Init(config.dataDir)
	if err != nil {
		log.Fatal(err)
	}
	// init the ERC20 token providers
	erc20Provider := new(web3.ERC20HolderProvider)
	if err := erc20Provider.Init(web3.Web3ProviderConfig{Endpoints: w3p}); err != nil {
		log.Fatal(err)
	}
	// init POAP external provider
	poapProvider := new(poap.POAPHolderProvider)
	if err := poapProvider.Init(poap.POAPConfig{
		URI:         config.poapAPIEndpoint,
		AccessToken: config.poapAuthToken,
	}); err != nil {
		log.Fatal(err)
	}
	holderProviders := map[uint64]providers.HolderProvider{
		providers.CONTRACT_TYPE_ERC20: erc20Provider,
		providers.CONTRACT_TYPE_POAP:  poapProvider,
	}
	// start the holder scanner with the database
	hc, err := service.NewHoldersScanner(database, w3p, holderProviders, config.scannerCoolDown)
	if err != nil {
		log.Fatal(err)
	}
	// if the admin token is not defined, generate a random one
	if config.adminToken != "" {
		if _, err := uuid.Parse(config.adminToken); err != nil {
			log.Fatal("bad admin token format, it must be a valid UUID")
		}
	} else {
		config.adminToken = uuid.New().String()
		log.Infof("no admin token defined, using a random one: %s", config.adminToken)
	}
	// Start the API
	apiService, err := api.Init(database, api.Census3APIConf{
		Hostname:        "0.0.0.0",
		Port:            config.port,
		DataDir:         config.dataDir,
		Web3Providers:   w3p,
		GroupKey:        config.connectKey,
		HolderProviders: holderProviders,
		AdminToken:      config.adminToken,
	})
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	go hc.Start(ctx)

	metrics.NewCounter(fmt.Sprintf("census3_info{version=%q,chains=%q}",
		internal.Version, w3p.String())).Set(1)

	// wait for SIGTERM
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Warnf("received SIGTERM, exiting at %s", time.Now().Format(time.RFC850))
	cancel()
	log.Infof("waiting for routines to end gracefully...")
	// closing database
	go func() {
		if err := apiService.Stop(); err != nil {
			log.Fatal(err)
		}
		if err := database.Close(); err != nil {
			log.Fatal(err)
		}
		for _, provider := range holderProviders {
			if err := provider.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	os.Exit(0)
}
