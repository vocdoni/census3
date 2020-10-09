package entitybridge

// Config wraps the entity bridge configs
type Config struct {
	DataDir,
	TokenContract,
	RegistryContract,
	ResolverContract,
	Web3HomeEndpoint,
	Web3ForeignEndpoint,
	GatewayURL,
	LogLevel,
	LogOutput,
	LogErrorFile,
	EthSigner string
	SameWeb3,
	SaveConfig bool
}

// NewConfig creates a new Config instance
func NewConfig() *Config {
	return &Config{}
}
