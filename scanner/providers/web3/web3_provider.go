package web3

type Web3ProviderRef struct {
	HexAddress string
	ChainID    uint64
}

type Web3ProviderConfig struct {
	Web3ProviderRef
	Endpoints NetworkEndpoints
}
