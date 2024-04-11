package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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

// HoldersByStrategy returns the holders of a strategy
func (c *HTTPclient) HoldersByStrategy(strategyID uint64) (*api.GetStrategyHoldersResponse, error) {
	// construct the URL to the API with the given parameters
	endpoint := fmt.Sprintf(GetTokenHoldersByStrategyURI, strategyID)
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
	holdersResponse := &api.GetStrategyHoldersResponse{}
	if err := json.NewDecoder(res.Body).Decode(holdersResponse); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodingResponse, err)
	}
	return holdersResponse, nil
}

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
	if res.StatusCode != http.StatusAccepted {
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
