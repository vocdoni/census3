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
	"github.com/vocdoni/census3/helpers/web3"
	"github.com/vocdoni/census3/internal"
	"github.com/vocdoni/census3/scanner"
	"github.com/vocdoni/census3/scanner/providers/farcaster"
	"github.com/vocdoni/census3/scanner/providers/gitcoin"
	gitcoinDB "github.com/vocdoni/census3/scanner/providers/gitcoin/db"
	"github.com/vocdoni/census3/scanner/providers/manager"
	"github.com/vocdoni/census3/scanner/providers/poap"
	web3provider "github.com/vocdoni/census3/scanner/providers/web3"
	"go.vocdoni.io/dvote/log"
)

type Census3Config struct {
	dataDir, logLevel, connectKey  string
	listOfWeb3Providers            []string
	port                           int
	poapAPIEndpoint, poapAuthToken string
	gitcoinEndpoint                string
	gitcoinCooldown                time.Duration
	scannerConcurrentTokens        int
	scannerCoolDown                time.Duration
	adminToken                     string
	initialTokens                  string
	farcaster                      bool
	filtersPath                    string
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
	flag.StringVar(&config.poapAPIEndpoint, "poapAPIEndpoint", "", "POAP API endpoint")
	flag.StringVar(&config.poapAuthToken, "poapAuthToken", "", "POAP API access token")
	flag.StringVar(&config.gitcoinEndpoint, "gitcoinEndpoint", "", "Gitcoin Passport API access token")
	flag.DurationVar(&config.gitcoinCooldown, "gitcoinCooldown", 6*time.Hour, "Gitcoin Passport API cooldown")
	var strWeb3Providers string
	flag.StringVar(&strWeb3Providers, "web3Providers", "", "the list of URL's of available web3 providers")
	flag.DurationVar(&config.scannerCoolDown, "scannerCoolDown", 120*time.Second, "the time to wait before next scanner iteration")
	flag.IntVar(&config.scannerConcurrentTokens, "scannerConcurrentTokens", 5, "the number of tokens to scan concurrently")
	flag.StringVar(&config.adminToken, "adminToken", "", "the admin UUID token for the API")
	flag.StringVar(&config.initialTokens, "initialTokens", "", "path of the initial tokens json file")
	flag.BoolVar(&config.farcaster, "farcaster", false, "enables farcaster support")
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
	if err := pviper.BindPFlag("gitcoinEndpoint", flag.Lookup("gitcoinEndpoint")); err != nil {
		panic(err)
	}
	config.gitcoinEndpoint = pviper.GetString("gitcoinEndpoint")
	if err := pviper.BindPFlag("gitcoinCooldown", flag.Lookup("gitcoinCooldown")); err != nil {
		panic(err)
	}
	config.gitcoinCooldown = pviper.GetDuration("gitcoinCooldown")
	if err := pviper.BindPFlag("web3Providers", flag.Lookup("web3Providers")); err != nil {
		panic(err)
	}
	config.listOfWeb3Providers = strings.Split(pviper.GetString("web3Providers"), ",")
	if err := pviper.BindPFlag("scannerConcurrentTokens", flag.Lookup("scannerConcurrentTokens")); err != nil {
		panic(err)
	}
	config.scannerConcurrentTokens = pviper.GetInt("scannerConcurrentTokens")
	if err := pviper.BindPFlag("scannerCoolDown", flag.Lookup("scannerCoolDown")); err != nil {
		panic(err)
	}
	config.scannerCoolDown = pviper.GetDuration("scannerCoolDown")
	if err := pviper.BindPFlag("adminToken", flag.Lookup("adminToken")); err != nil {
		panic(err)
	}
	config.adminToken = pviper.GetString("adminToken")
	if err := pviper.BindPFlag("initialTokens", flag.Lookup("initialTokens")); err != nil {
		panic(err)
	}
	config.initialTokens = pviper.GetString("initialTokens")
	if err := pviper.BindPFlag("farcaster", flag.Lookup("farcaster")); err != nil {
		panic(err)
	}
	config.farcaster = pviper.GetBool("farcaster")
	// set the filters path into the config, create the folder if it does not
	// exitst yet
	config.filtersPath = config.dataDir + "/filters"
	if err := os.MkdirAll(config.filtersPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	// init logger
	log.Init(config.logLevel, "stdout", nil)
	// check if the web3 providers are defined
	if len(config.listOfWeb3Providers) == 0 {
		log.Fatal("no web3 providers defined")
	}
	// check if the web3 providers are valid
	w3p, err := web3.NewWeb3Pool()
	if err != nil {
		log.Fatal(err)
	}
	for _, uri := range config.listOfWeb3Providers {
		if err := w3p.AddEndpoint(uri); err != nil {
			log.Warnf("error adding web3 provider (%s): %v", uri, err)
		}
	}
	// init the database
	database, err := db.Init(config.dataDir, "census3.sql")
	if err != nil {
		log.Fatal(err)
	}
	// init the provider manager
	pm := manager.NewProviderManager()
	// init the web3 token providers
	web3ProviderConf := web3provider.Web3ProviderConfig{Endpoints: w3p}
	pm.AddProvider(new(web3provider.ERC20HolderProvider).Type(), web3ProviderConf)
	pm.AddProvider(new(web3provider.ERC721HolderProvider).Type(), web3ProviderConf)
	pm.AddProvider(new(web3provider.ERC777HolderProvider).Type(), web3ProviderConf)
	// init POAP external provider
	if config.poapAPIEndpoint != "" {
		pm.AddProvider(new(poap.POAPHolderProvider).Type(), poap.POAPConfig{
			APIEndpoint: config.poapAPIEndpoint,
			AccessToken: config.poapAuthToken,
		})
	}
	if config.gitcoinEndpoint != "" {
		gitcoinDatabase, err := gitcoinDB.Init(config.dataDir, "gitcoinpassport.sql")
		if err != nil {
			log.Fatal(err)
		}
		pm.AddProvider(new(gitcoin.GitcoinPassport).Type(), gitcoin.GitcoinPassportConf{
			APIEndpoint: config.gitcoinEndpoint,
			Cooldown:    config.gitcoinCooldown,
			DB:          gitcoinDatabase,
		})
	}
	// if farcaster is enabled, init the farcaster database and the provider
	var farcasterDB *farcaster.DB
	if config.farcaster {
		log.Debugf("farcaster support enabled")
		farcasterDB, err = farcaster.InitDB(config.dataDir, "farcaster.sql")
		if err != nil {
			log.Fatal(err)
		}
		pm.AddProvider(new(farcaster.FarcasterProvider).Type(), farcaster.FarcasterProviderConf{
			Endpoints: w3p,
			DB:        farcasterDB,
		})
	}
	// start the token updater with the database and the provider manager
	updater := scanner.NewUpdater(database, w3p, pm, config.filtersPath)
	// start the holder scanner with the database and the provider manager
	hc := scanner.NewScanner(database, updater, w3p, pm, config.scannerCoolDown, config.filtersPath)
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
	ctx, cancel := context.WithCancel(context.Background())
	apiService, err := api.Init(database, api.Census3APIConf{
		MainCtx:         ctx,
		Hostname:        "0.0.0.0",
		Port:            config.port,
		DataDir:         config.dataDir,
		Web3Providers:   w3p,
		GroupKey:        config.connectKey,
		HolderProviders: pm.Providers(ctx),
		AdminToken:      config.adminToken,
		TokenUpdater:    updater,
	})
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		// create initial tokens in background
		if err := apiService.CreateInitialTokens(config.initialTokens); err != nil {
			log.Warnf("error creating initial tokens: %s", err)
		}
		log.Info("initial tokens created, or at least tried to")
	}()
	// start the holder scanner
	go hc.Start(ctx)
	go updater.Start(ctx, config.scannerConcurrentTokens)

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
		hc.Stop()
		updater.Stop()
		if err := apiService.Stop(); err != nil {
			log.Fatal(err)
		}
		if err := database.Close(); err != nil {
			log.Fatal(err)
		}
		// if farcaster is enabled, close the farcaster database
		if config.farcaster {
			if err := farcasterDB.CloseDB(); err != nil {
				log.Fatal(err)
			}
		}
		log.Infof("all routines ended")
	}()
	time.Sleep(5 * time.Second)
	os.Exit(0)
}
