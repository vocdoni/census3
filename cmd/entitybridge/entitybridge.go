package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/vocdoni/tokenstate/entitybridge"
	"gitlab.com/vocdoni/go-dvote/config"
	"gitlab.com/vocdoni/go-dvote/crypto/ethereum"
	"gitlab.com/vocdoni/go-dvote/log"
)

const defaultDir = "/.entitybridge"

func newConfig() (*entitybridge.Config, config.Error) {
	var err error
	var cfgError config.Error

	// get home dir
	cfg := entitybridge.NewConfig()
	home, err := os.UserHomeDir()
	if err != nil {
		cfgError = config.Error{
			Critical: true,
			Message:  fmt.Sprintf("cannot get user home directory with error: %s", err),
		}
		return nil, cfgError
	}

	// flags
	flag.StringVar(&cfg.DataDir, "dataDir", home+defaultDir, "data directory for persistent storage")

	cfg.TokenContract = *flag.String("tokenContract", "", "token contract address")
	cfg.RegistryContract = *flag.String("registryContract", "", "registry contract address")
	cfg.ResolverContract = *flag.String("resolverContract", "", "resolver contract address")
	cfg.Web3HomeEndpoint = *flag.String("web3Home", "", "web3 endpoint pointing to the network used for fetching the tokens info")
	cfg.Web3ForeignEndpoint = *flag.String("web3Foreign", "", "web3 endpoint pointing to the network used for creating the entities")
	cfg.SameWeb3 = *flag.Bool("sameWeb3", false, "if true the bridge will act as a simple transactor on the same network")
	cfg.GatewayURL = *flag.String("gatewayURL", "", "gateway api endpoint that gives access to IPFS")
	cfg.EthSigner = *flag.String("ethSigner", "", "ethereum sign keys private key")
	cfg.LogLevel = *flag.String("logLevel", "info", "Log level (debug, info, warn, error, fatal)")
	cfg.LogOutput = *flag.String("logOutput", "stdout", "Log output (stdout, stderr or filepath)")
	cfg.LogErrorFile = *flag.String("logErrorFile", "", "Log errors and warnings to a file")
	cfg.SaveConfig = *flag.Bool("saveConfig", false, "overwrites an existing config file with the CLI provided flags")
	flag.CommandLine.SortFlags = false

	flag.Parse()

	// setting up viper
	viper := viper.New()
	viper.SetConfigName("entityridge")
	viper.SetConfigType("yml")
	viper.SetEnvPrefix("EBRIDGE")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// set FlagVars first
	viper.BindPFlag("dataDir", flag.Lookup("dataDir"))
	cfg.DataDir = viper.GetString("dataDir")

	// Add viper config path (now we know it)
	viper.AddConfigPath(cfg.DataDir)

	// binding flags
	viper.BindPFlag(("tokenContract"), flag.Lookup("tokenContract"))
	viper.BindPFlag(("registryContract"), flag.Lookup("registryContract"))
	viper.BindPFlag(("resolverContract"), flag.Lookup("resolverContract"))
	viper.BindPFlag(("web3HomeEndpoint"), flag.Lookup("web3Home"))
	viper.BindPFlag(("web3ForeignEndpoint"), flag.Lookup("web3Foreign"))
	viper.BindPFlag(("sameWeb3"), flag.Lookup("sameWeb3"))
	viper.BindPFlag(("gatewayURL"), flag.Lookup("gatewayURL"))
	viper.BindPFlag(("ethSigner"), flag.Lookup("ethSigner"))
	viper.BindPFlag("logLevel", flag.Lookup("logLevel"))
	viper.BindPFlag("logErrorFile", flag.Lookup("logErrorFile"))
	viper.BindPFlag("logOutput", flag.Lookup("logOutput"))
	viper.BindPFlag("saveConfig", flag.Lookup("saveConfig"))

	// check if config file exists
	_, err = os.Stat(cfg.DataDir + "/entitybridge.yml")
	if os.IsNotExist(err) {
		cfgError = config.Error{
			Message: fmt.Sprintf("creating new config file in %s", cfg.DataDir),
		}
		// creting config folder if not exists
		err = os.MkdirAll(cfg.DataDir, os.ModePerm)
		if err != nil {
			cfgError = config.Error{
				Message: fmt.Sprintf("cannot create data directory: %s", err),
			}
		}
		// create config file if not exists
		if err := viper.SafeWriteConfig(); err != nil {
			cfgError = config.Error{
				Message: fmt.Sprintf("cannot write config file into config dir: %s", err),
			}
		}
	} else {
		// read config file
		err = viper.ReadInConfig()
		if err != nil {
			cfgError = config.Error{
				Message: fmt.Sprintf("cannot read loaded config file in %s: %s", cfg.DataDir, err),
			}
		}
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		cfgError = config.Error{
			Message: fmt.Sprintf("cannot unmarshal loaded config file: %s", err),
		}
	}

	if len(cfg.EthSigner) < 32 {
		fmt.Println("no signing key, generating one...")
		signer := ethereum.NewSignKeys()
		err = signer.Generate()
		if err != nil {
			cfgError = config.Error{
				Message: fmt.Sprintf("cannot generate signing key: %s", err),
			}
			return cfg, cfgError
		}
		_, priv := signer.HexString()
		viper.Set("ethSigner", priv)
		cfg.EthSigner = priv
		cfg.SaveConfig = true
	}

	if cfg.SaveConfig {
		viper.Set("saveConfig", false)
		if err := viper.WriteConfig(); err != nil {
			cfgError = config.Error{
				Message: fmt.Sprintf("cannot overwrite config file into config dir: %s", err),
			}
		}
	}

	return cfg, cfgError
}

func main() {

	// setup config
	// creating config and init logger
	cfg, cfgErr := newConfig()
	if cfg == nil {
		log.Fatal("cannot read configuration")
	}
	log.Init(cfg.LogLevel, cfg.LogOutput)
	if path := cfg.LogErrorFile; path != "" {
		if err := log.SetFileErrorLog(path); err != nil {
			log.Fatal(err)
		}
	}
	log.Debugf("initializing config %+v", *cfg)

	// check if errors during config creation and determine if Critical
	if cfgErr.Critical && cfgErr.Message != "" {
		log.Fatalf("critical error loading config: %s", cfgErr.Message)
	} else if !cfgErr.Critical && cfgErr.Message != "" {
		log.Warnf("non Critical error loading config: %s", cfgErr.Message)
	} else if !cfgErr.Critical && cfgErr.Message == "" {
		log.Infof("config file loaded successfully, remember CLI flags have preference")
	}

	var signer *ethereum.SignKeys
	signer = ethereum.NewSignKeys()
	// add signing private key if exist in configuration or flags
	if len(cfg.EthSigner) != 32 {
		log.Infof("adding custom signing key")
		err := signer.AddHexKey(cfg.EthSigner)
		if err != nil {
			log.Fatalf("error adding hex key: (%s)", err)
		}
		pub, _ := signer.HexString()
		log.Infof("using custom pubKey %s", pub)
	} else {
		log.Fatal("no private key or wrong key (size != 16 bytes)")
	}

	b := entitybridge.NewEntityBridgeService()
	if err := b.Init(context.Background(), cfg, signer); err != nil {
		log.Warnf("entity bridge service initialization error: %s", err)
		return
	}

	// create token entity
	res, err := b.CreateEntityMetadata()
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("ipfs file URL: %s", res)
}
