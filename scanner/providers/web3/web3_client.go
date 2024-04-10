package web3

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const defaultRetries = 3

var (
	defaultTimeout    = 2 * time.Second
	filterLogsTimeout = 3 * time.Second
	retrySleep        = 200 * time.Millisecond
)

// Client struct implements bind.ContractBackend interface for a web3 pool with
// an specific chainID. It allows to interact with the blockchain using the
// methods provided by the interface balancing the load between the available
// endpoints in the pool for the chainID.
type Client struct {
	w3p     *Web3Pool
	chainID uint64
}

// EthClient method returns the ethclient.Client for the chainID of the Client
// instance. It returns an error if the chainID is not found in the pool.
func (c *Client) EthClient() (*ethclient.Client, error) {
	endpoint, ok := c.w3p.GetEndpoint(c.chainID)
	if !ok {
		return nil, fmt.Errorf("error getting endpoint for chainID %d", c.chainID)
	}
	return endpoint.client, nil
}

// CodeAt method wraps the CodeAt method from the ethclient.Client for the
// chainID of the Client instance. It returns an error if the chainID is not
// found in the pool or if the method fails. Required by the bind.ContractBackend
// interface.
func (c *Client) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	endpoint, ok := c.w3p.GetEndpoint(c.chainID)
	if !ok {
		return nil, fmt.Errorf("error getting endpoint for chainID %d", c.chainID)
	}
	// retry the method in case of failure and get final result and error
	res, err := c.retryAndCheckErr(endpoint.URI, func() (any, error) {
		internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		return endpoint.client.CodeAt(internalCtx, account, blockNumber)
	})
	if err != nil {
		return nil, err
	}
	return res.([]byte), err
}

// CallContract method wraps the CallContract method from the ethclient.Client
// for the chainID of the Client instance. It returns an error if the chainID is
// not found in the pool or if the method fails. Required by the
// bind.ContractBackend interface.
func (c *Client) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	endpoint, ok := c.w3p.GetEndpoint(c.chainID)
	if !ok {
		return nil, fmt.Errorf("error getting endpoint for chainID %d", c.chainID)
	}
	// retry the method in case of failure and get final result and error
	res, err := c.retryAndCheckErr(endpoint.URI, func() (any, error) {
		internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		return endpoint.client.CallContract(internalCtx, call, blockNumber)
	})
	if err != nil {
		return nil, err
	}
	return res.([]byte), err
}

// EstimateGas method wraps the EstimateGas method from the ethclient.Client for
// the chainID of the Client instance. It returns an error if the chainID is not
// found in the pool or if the method fails. Required by the bind.ContractBackend
// interface.
func (c *Client) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	endpoint, ok := c.w3p.GetEndpoint(c.chainID)
	if !ok {
		return 0, fmt.Errorf("error getting endpoint for chainID %d", c.chainID)
	}
	// retry the method in case of failure and get final result and error
	res, err := c.retryAndCheckErr(endpoint.URI, func() (any, error) {
		internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		return endpoint.client.EstimateGas(internalCtx, msg)
	})
	if err != nil {
		return 0, err
	}
	return res.(uint64), err
}

// FilterLogs method wraps the FilterLogs method from the ethclient.Client for
// the chainID of the Client instance. It returns an error if the chainID is not
// found in the pool or if the method fails. Required by the bind.ContractBackend
// interface.
func (c *Client) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	endpoint, ok := c.w3p.GetEndpoint(c.chainID)
	if !ok {
		return nil, fmt.Errorf("error getting endpoint for chainID %d", c.chainID)
	}
	// retry the method in case of failure and get final result and error
	res, err := c.retryAndCheckErr(endpoint.URI, func() (any, error) {
		internalCtx, cancel := context.WithTimeout(ctx, filterLogsTimeout)
		defer cancel()
		return endpoint.client.FilterLogs(internalCtx, query)
	})
	if err != nil {
		return nil, err
	}
	return res.([]types.Log), nil
}

// HeaderByNumber method wraps the HeaderByNumber method from the ethclient.Client
// for the chainID of the Client instance. It returns an error if the chainID is
// not found in the pool or if the method fails. Required by the
// bind.ContractBackend interface.
func (c *Client) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	endpoint, ok := c.w3p.GetEndpoint(c.chainID)
	if !ok {
		return nil, fmt.Errorf("error getting endpoint for chainID %d", c.chainID)
	}
	// retry the method in case of failure and get final result and error
	res, err := c.retryAndCheckErr(endpoint.URI, func() (any, error) {
		internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		return endpoint.client.HeaderByNumber(internalCtx, number)
	})
	if err != nil {
		return nil, err
	}
	return res.(*types.Header), err
}

// PendingNonceAt method wraps the PendingNonceAt method from the
// ethclient.Client for the chainID of the Client instance. It returns an error
// if the chainID is not found in the pool or if the method fails. Required by
// the bind.ContractBackend interface.
func (c *Client) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	endpoint, ok := c.w3p.GetEndpoint(c.chainID)
	if !ok {
		return 0, fmt.Errorf("error getting endpoint for chainID %d", c.chainID)
	}
	// retry the method in case of failure and get final result and error
	res, err := c.retryAndCheckErr(endpoint.URI, func() (any, error) {
		internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		return endpoint.client.PendingNonceAt(internalCtx, account)
	})
	if err != nil {
		return 0, err
	}
	return res.(uint64), err
}

// SuggestGasPrice method wraps the SuggestGasPrice method from the
// ethclient.Client for the chainID of the Client instance. It returns an error
// if the chainID is not found in the pool or if the method fails. Required by
// the bind.ContractBackend interface.
func (c *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	endpoint, ok := c.w3p.GetEndpoint(c.chainID)
	if !ok {
		return nil, fmt.Errorf("error getting endpoint for chainID %d", c.chainID)
	}
	// retry the method in case of failure and get final result and error
	res, err := c.retryAndCheckErr(endpoint.URI, func() (any, error) {
		internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		return endpoint.client.SuggestGasPrice(internalCtx)
	})
	if err != nil {
		return nil, err
	}
	return res.(*big.Int), err
}

// SendTransaction method wraps the SendTransaction method from the ethclient.Client
// for the chainID of the Client instance. It returns an error if the chainID is
// not found in the pool or if the method fails. Required by the
// bind.ContractBackend interface.
func (c *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	endpoint, ok := c.w3p.GetEndpoint(c.chainID)
	if !ok {
		return fmt.Errorf("error getting endpoint for chainID %d", c.chainID)
	}
	// retry the method in case of failure and get final result and error
	_, err := c.retryAndCheckErr(endpoint.URI, func() (any, error) {
		internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		return nil, endpoint.client.SendTransaction(internalCtx, tx)
	})
	return err
}

// PendingCodeAt method wraps the PendingCodeAt method from the ethclient.Client
// for the chainID of the Client instance. It returns an error if the chainID is
// not found in the pool or if the method fails. Required by the
// bind.ContractBackend interface.
func (c *Client) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	endpoint, ok := c.w3p.GetEndpoint(c.chainID)
	if !ok {
		return nil, fmt.Errorf("error getting endpoint for chainID %d", c.chainID)
	}
	// retry the method in case of failure and get final result and error
	res, err := c.retryAndCheckErr(endpoint.URI, func() (any, error) {
		internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		return endpoint.client.PendingCodeAt(internalCtx, account)
	})
	if err != nil {
		return nil, err
	}
	return res.([]byte), err
}

// SubscribeFilterLogs method wraps the SubscribeFilterLogs method from the
// ethclient.Client for the chainID of the Client instance. It returns an error
// if the chainID is not found in the pool or if the method fails. Required by
// the bind.ContractBackend interface.
func (c *Client) SubscribeFilterLogs(ctx context.Context,
	query ethereum.FilterQuery, ch chan<- types.Log,
) (ethereum.Subscription, error) {
	endpoint, ok := c.w3p.GetEndpoint(c.chainID)
	if !ok {
		return nil, fmt.Errorf("error getting endpoint for chainID %d", c.chainID)
	}
	// retry the method in case of failure and get final result and error
	res, err := c.retryAndCheckErr(endpoint.URI, func() (any, error) {
		internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		return endpoint.client.SubscribeFilterLogs(internalCtx, query, ch)
	})
	if err != nil {
		return nil, err
	}
	return res.(ethereum.Subscription), err
}

// SuggestGasTipCap method wraps the SuggestGasTipCap method from the
// ethclient.Client for the chainID of the Client instance. It returns an error
// if the chainID is not found in the pool or if the method fails. Required by
// the bind.ContractBackend interface.
func (c *Client) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	endpoint, ok := c.w3p.GetEndpoint(c.chainID)
	if !ok {
		return nil, fmt.Errorf("error getting endpoint for chainID %d", c.chainID)
	}
	// retry the method in case of failure and get final result and error
	res, err := c.retryAndCheckErr(endpoint.URI, func() (any, error) {
		internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		return endpoint.client.SuggestGasTipCap(internalCtx)
	})
	if err != nil {
		return nil, err
	}
	return res.(*big.Int), err
}

// BalanceAt method wraps the BalanceAt method from the ethclient.Client for the
// chainID of the Client instance. It returns an error if the chainID is not
// found in the pool or if the method fails. This method is required by internal
// logic, it is not required by the bind.ContractBackend interface.
func (c *Client) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	endpoint, ok := c.w3p.GetEndpoint(c.chainID)
	if !ok {
		return nil, fmt.Errorf("error getting endpoint for chainID %d", c.chainID)
	}
	// retry the method in case of failure and get final result and error
	res, err := c.retryAndCheckErr(endpoint.URI, func() (any, error) {
		internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		return endpoint.client.BalanceAt(internalCtx, account, blockNumber)
	})
	if err != nil {
		return nil, err
	}
	return res.(*big.Int), err
}

// BlockNumber method wraps the BlockNumber method from the ethclient.Client for
// the chainID of the Client instance. It returns an error if the chainID is not
// found in the pool or if the method fails. This method is required by internal
// logic, it is not required by the bind.ContractBackend interface.
func (c *Client) BlockNumber(ctx context.Context) (uint64, error) {
	endpoint, ok := c.w3p.GetEndpoint(c.chainID)
	if !ok {
		return 0, fmt.Errorf("error getting endpoint for chainID %d", c.chainID)
	}
	// retry the method in case of failure and get final result and error
	res, err := c.retryAndCheckErr(endpoint.URI, func() (any, error) {
		internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		return endpoint.client.BlockNumber(internalCtx)
	})
	if err != nil {
		return 0, err
	}
	return res.(uint64), err
}

// retryAndCheckErr method retries a function call in case of error and checks
// the error after the retries. It returns the result of the function call and
// the error if the retries are exhausted. It is used to retry the methods of
// the ethclient.Client in case of failure. If the error is not nil after the
// retries, the endpoint is disabled in the pool and the error is returned.
func (c *Client) retryAndCheckErr(uri string, fn func() (any, error)) (any, error) {
	var res any
	var err error
	for i := 0; i < defaultRetries; i++ {
		res, err = fn()
		if err == nil {
			return res, nil
		}
		time.Sleep(retrySleep)
	}
	c.w3p.DisableEndpoint(c.chainID, uri)
	return nil, fmt.Errorf("error after %d retries: %w", defaultRetries, err)
}
