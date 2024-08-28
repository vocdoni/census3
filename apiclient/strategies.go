package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/api"
	"go.vocdoni.io/dvote/log"
)

func (c *HTTPclient) Strategies(pageSize int, nextCursor, prevCursor string) (
	[]*api.Strategy, error,
) {
	// construct the URL to the API with the given parameters
	endpoint := fmt.Sprintf(GetStrategiesURI, pageSize, nextCursor, prevCursor)
	u, err := c.constructURL(endpoint)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrConstructingURL, err)
	}
	// create the request and send it, if there is an error or the status code
	// is not 200, return an error
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCreatingRequest, err)
	}
	res, err := c.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrMakingRequest, err)
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Errorf("error closing response body: %v", err)
		}
	}()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %s", ErrNoStatusOk,
			fmt.Errorf("%d %s", res.StatusCode, http.StatusText(res.StatusCode)))
	}
	strategiesResponse := &api.Strategies{}
	if err := json.NewDecoder(res.Body).Decode(strategiesResponse); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodingResponse, err)
	}
	return strategiesResponse.Strategies, nil
}

func (c *HTTPclient) Strategy(strategyID uint64) (*api.Strategy, error) {
	// construct the URL to the API with the given parameters
	endpoint := fmt.Sprintf(GetStrategyURI, strategyID)
	u, err := c.constructURL(endpoint)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrConstructingURL, err)
	}
	// create the request and send it, if there is an error or the status code
	// is not 200, return an error
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCreatingRequest, err)
	}
	res, err := c.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrMakingRequest, err)
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Errorf("error closing response body: %v", err)
		}
	}()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %s", ErrNoStatusOk,
			fmt.Errorf("%d %s", res.StatusCode, http.StatusText(res.StatusCode)))
	}
	strategyResponse := &api.Strategy{}
	if err := json.NewDecoder(res.Body).Decode(strategyResponse); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodingResponse, err)
	}
	return strategyResponse, nil
}

// HoldersByStrategy method queries the API for the holders of a strategy, it
// receives the strategyID and returns the queueID of the task and an error if
// something went wrong. The status of the task can be checked with the
// HoldersByStrategyQueue method.
// If truncateByDecimals is true, the amounts will be truncated by the decimals
// of the token (if > 0). Truncate only works for simple strategies (single token).
func (c *HTTPclient) HoldersByStrategy(strategyID uint64, truncateByDecimals bool) (string, error) {
	// construct the URL to the API with the given parameters
	endpoint := fmt.Sprintf(GetTokenHoldersByStrategyURI, strategyID, fmt.Sprintf("%t", truncateByDecimals))
	u, err := c.constructURL(endpoint)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrConstructingURL, err)
	}
	// create the request and send it, if there is an error or the status code
	// is not 200, return an error
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrCreatingRequest, err)
	}
	res, err := c.c.Do(req)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrMakingRequest, err)
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Errorf("error closing response body: %v", err)
		}
	}()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%w: %s", ErrNoStatusOk,
			fmt.Errorf("%d %s", res.StatusCode, http.StatusText(res.StatusCode)))
	}
	holdersResponse := &api.QueueResponse{}
	if err := json.NewDecoder(res.Body).Decode(holdersResponse); err != nil {
		return "", fmt.Errorf("%w: %w", ErrDecodingResponse, err)
	}
	return holdersResponse.QueueID, nil
}

// HoldersByStrategyQueue method checks the status of the query for the holders
// of a strategy from the API, it receives the strategyID and the queueID and
// returns a map of addresses and amounts, a boolean indicating if the queue
// task is completed and an error if something went wrong.
func (c *HTTPclient) HoldersByStrategyQueue(strategyID uint64, queueID string) (
	map[common.Address]*big.Int, bool, error,
) {
	if strategyID == 0 {
		return nil, false, fmt.Errorf("%w: strategyID is required", ErrBadInputs)
	}
	if queueID == "" {
		return nil, false, fmt.Errorf("%w: queueID is required", ErrBadInputs)
	}
	// construct the URL to the API with the given parameters
	endpoint := fmt.Sprintf(GetTokenHoldersByStrategyQueueURI, strategyID, queueID)
	u, err := c.constructURL(endpoint)
	if err != nil {
		return nil, false, fmt.Errorf("%w: %w", ErrConstructingURL, err)
	}
	// create the request and send it, if there is an error or the status code
	// is not 200, return an error
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, false, fmt.Errorf("%w: %w", ErrCreatingRequest, err)
	}
	res, err := c.c.Do(req)
	if err != nil {
		return nil, false, fmt.Errorf("%w: %w", ErrMakingRequest, err)
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Errorf("error closing response body: %v", err)
		}
	}()
	if res.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("%w: %s", ErrNoStatusOk,
			fmt.Errorf("%d %s", res.StatusCode, http.StatusText(res.StatusCode)))
	}
	// decode the queue response
	item := map[string]interface{}{}
	if err := json.NewDecoder(res.Body).Decode(&item); err != nil {
		return nil, false, fmt.Errorf("%w: %w", ErrDecodingResponse, err)
	}
	// check if the item is done and if there is an error
	if done, ok := item["done"].(bool); !ok || !done {
		return nil, false, nil
	}
	if strErr, ok := item["error"].(string); ok && strErr != "" {
		return nil, true, fmt.Errorf("error in queue item: %s", strErr)
	}
	// convert the data to a map of addresses and amounts
	rawHolders, ok := item["data"].(map[string]interface{})
	if !ok {
		return nil, true, fmt.Errorf("error getting data from queue item")
	}
	holders := make(map[common.Address]*big.Int, len(rawHolders))
	for k, v := range rawHolders {
		strBalance, ok := v.(string)
		if !ok {
			continue
		}
		addr := common.HexToAddress(k)
		amount := new(big.Int)
		if _, ok := amount.SetString(strBalance, 10); !ok {
			return nil, true, fmt.Errorf("error converting amount to big.Int")
		}
		holders[addr] = amount
	}
	return holders, true, nil
}

// AllHoldersByStrategy method queries the API for all the holders of a
// strategy, it receives the strategyID and returns a map of addresses and
// amounts and an error if something went wrong. This method is a wrapper
// around the HoldersByStrategy and HoldersByStrategyQueue methods.
// If truncateByDecimals is true, the amounts will be truncated by the decimals
// of the token (if > 0). Truncate only works for simple strategies (single token).
func (c *HTTPclient) AllHoldersByStrategy(strategyID uint64, truncateByDecimals bool) (map[common.Address]*big.Int, error) {
	queueID, err := c.HoldersByStrategy(strategyID, truncateByDecimals)
	if err != nil {
		return nil, err
	}
	for {
		holders, done, err := c.HoldersByStrategyQueue(strategyID, queueID)
		if err != nil {
			return nil, err
		}
		if done {
			return holders, nil
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// CreateStrategy method sends a request to the API to create a new strategy, it
// receives the strategy data and returns the strategyID and an error if
// something went wrong.
func (c *HTTPclient) CreateStrategy(request *api.Strategy) (uint64, error) {
	// construct the URL to the API
	u, err := c.constructURL(CreateStrategyURI)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrConstructingURL, err)
	}
	// create the request and send it, if there is an error or the status code
	// is not 200, return an error
	body, err := json.Marshal(request)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrEncodingRequest, err)
	}
	req, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(body))
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrCreatingRequest, err)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.c.Do(req)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrMakingRequest, err)
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Errorf("error closing response body: %v", err)
		}
	}()
	switch res.StatusCode {
	case http.StatusAccepted, http.StatusOK:
		break
	case http.StatusConflict:
		return 0, ErrAlreadyExists
	default:
		return 0, fmt.Errorf("%w: %s", ErrNoStatusOk,
			fmt.Errorf("%d %s", res.StatusCode, http.StatusText(res.StatusCode)))
	}
	responseStrategyID := map[string]uint64{}
	if err := json.NewDecoder(res.Body).Decode(&responseStrategyID); err != nil {
		return 0, fmt.Errorf("%w: %w", ErrDecodingResponse, err)
	}
	strategyID, ok := responseStrategyID["strategyID"]
	if !ok {
		return 0, fmt.Errorf("error getting strategyID from response")
	}
	return strategyID, nil
}
