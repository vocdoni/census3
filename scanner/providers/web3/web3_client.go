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

var (
	defaultTimeout    = 2 * time.Second
	filterLogsTimeout = 3 * time.Second
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
	internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	// check if the method fails, if it does, disable the endpoint
	res, err := endpoint.client.CodeAt(internalCtx, account, blockNumber)
	c.checkErr(err, endpoint.URI)
	return res, err
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
	internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	// check if the method fails, if it does, disable the endpoint
	res, err := endpoint.client.CallContract(internalCtx, call, blockNumber)
	c.checkErr(err, endpoint.URI)
	return res, err
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
	internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	// check if the method fails, if it does, disable the endpoint
	res, err := endpoint.client.EstimateGas(internalCtx, msg)
	c.checkErr(err, endpoint.URI)
	return res, err
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
	internalCtx, cancel := context.WithTimeout(ctx, filterLogsTimeout)
	defer cancel()
	// check if the method fails, if it does, disable the endpoint
	res, err := endpoint.client.FilterLogs(internalCtx, query)
	c.checkErr(err, endpoint.URI)
	return res, err
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
	internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	// check if the method fails, if it does, disable the endpoint
	res, err := endpoint.client.HeaderByNumber(internalCtx, number)
	c.checkErr(err, endpoint.URI)
	return res, err
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
	internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	// check if the method fails, if it does, disable the endpoint
	res, err := endpoint.client.PendingNonceAt(internalCtx, account)
	c.checkErr(err, endpoint.URI)
	return res, err
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
	internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	// check if the method fails, if it does, disable the endpoint
	res, err := endpoint.client.SuggestGasPrice(internalCtx)
	c.checkErr(err, endpoint.URI)
	return res, err
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
	internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	// check if the method fails, if it does, disable the endpoint
	err := endpoint.client.SendTransaction(internalCtx, tx)
	c.checkErr(err, endpoint.URI)
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
	internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	// check if the method fails, if it does, disable the endpoint
	res, err := endpoint.client.PendingCodeAt(internalCtx, account)
	c.checkErr(err, endpoint.URI)
	return res, err
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
	internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	// check if the method fails, if it does, disable the endpoint
	res, err := endpoint.client.SubscribeFilterLogs(internalCtx, query, ch)
	c.checkErr(err, endpoint.URI)
	return res, err
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
	internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	// check if the method fails, if it does, disable the endpoint
	res, err := endpoint.client.SuggestGasTipCap(internalCtx)
	c.checkErr(err, endpoint.URI)
	return res, err
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
	internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	// check if the method fails, if it does, disable the endpoint
	res, err := endpoint.client.BalanceAt(internalCtx, account, blockNumber)
	c.checkErr(err, endpoint.URI)
	return res, err
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
	internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	// check if the method fails, if it does, disable the endpoint
	res, err := endpoint.client.BlockNumber(internalCtx)
	c.checkErr(err, endpoint.URI)
	return res, err
}

// checkErr method disables the endpoint if the error is not nil.
func (c *Client) checkErr(err error, uri string) {
	if err != nil {
		c.w3p.DisableEndpoint(c.chainID, uri)
	}
}
