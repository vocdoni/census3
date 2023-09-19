package state

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func CheckWeb3Providers(providersURIs []string) (map[uint64]string, error) {
	if len(providersURIs) == 0 {
		return nil, fmt.Errorf("no URIs provided")
	}

	providers := make(map[uint64]string)
	for _, uri := range providersURIs {
		cli, err := ethclient.Dial(uri)
		if err != nil {
			return nil, fmt.Errorf("error dialing web3 provider uri '%s': %w", uri, err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		// get the chainID from the web3 endpoint
		chainID, err := cli.ChainID(ctx)
		if err != nil {
			return nil, fmt.Errorf("error getting the chainID from the web3 provider '%s': %w", uri, err)
		}
		providers[uint64(chainID.Int64())] = uri
	}
	return providers, nil
}
