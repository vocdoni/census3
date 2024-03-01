package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vocdoni/census3/api"
	"go.vocdoni.io/dvote/log"
)

// GetCensus method returns a census by its ID from the API, it receives the
// censusID and returns a pointer to a GetCensusResponse and an error if something
// went wrong.
func (c *HTTPclient) GetCensus(censusID uint64) (*api.GetCensusResponse, error) {
	if censusID == 0 {
		return nil, fmt.Errorf("%w: censusID is required", ErrBadInputs)
	}
	// construct the URL to the API with the censusID
	endpoint := fmt.Sprintf(GetCensusURI, censusID)
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
		return nil, fmt.Errorf("%w: %w", ErrNoStatusOk, fmt.Errorf("%d %s", res.StatusCode, http.StatusText(res.StatusCode)))
	}
	// decode the response and return it
	response := api.GetCensusResponse{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodingResponse, err)
	}
	return &response, nil
}

// CreateCensus method creates a new census from the API, it receives a
// CreateCensusRequest and returns a queueID and an error if something went wrong.
// The queueID is used to check the status of the census creation process, it
// can be checked with the CreateCensusQueue method.
func (c *HTTPclient) CreateCensus(request *api.CreateCensusRequest) (string, error) {
	if request == nil || request.StrategyID == 0 {
		return "", fmt.Errorf("%w: strategyID is required", ErrBadInputs)
	}
	// construct the URL to the API
	url, err := c.constructURL(CreateCensusURI)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrConstructingURL, err)
	}
	// encode the input token to JSON to be sent in the request body
	requestBody, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrEncodingRequest, err)
	}
	// create the request and send it with the encoded body, if there is an
	// error or the status code is not 200, return an error
	req, err := http.NewRequest(HTTPPOST, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrCreatingRequest, err)
	}
	req.Header.Set("Content-Type", "application/json")
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
		return "", fmt.Errorf("%w: %w", ErrNoStatusOk, fmt.Errorf("%d %s", res.StatusCode, http.StatusText(res.StatusCode)))
	}
	// decode the response and return the queueID
	queueResponse := &api.QueueResponse{}
	if err := json.NewDecoder(res.Body).Decode(queueResponse); err != nil {
		return "", fmt.Errorf("%w: %w", ErrDecodingResponse, err)
	}
	return queueResponse.QueueID, nil
}

// CreateCensusQueue method checks the status of a census creation process from
// the API, it receives a queueID and returns a pointer to a CensusQueueResponse
// and an error if something went wrong.
func (c *HTTPclient) CreateCensusQueue(queueID string) (*api.CensusQueueResponse, error) {
	if queueID == "" {
		return nil, fmt.Errorf("%w: queueID is required", ErrBadInputs)
	}
	// construct the URL to the API with the queueID
	endpoint := fmt.Sprintf(CreateCensusQueueURI, queueID)
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
		return nil, fmt.Errorf("%w: %w", ErrNoStatusOk, fmt.Errorf("%d %s", res.StatusCode, http.StatusText(res.StatusCode)))
	}
	// decode the response and return it
	response := &api.CensusQueueResponse{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodingResponse, err)
	}
	return response, nil
}

// GetCensusesByStrategy method returns the censuses of a strategy from the API,
// it receives the strategyID and returns a slice of GetCensusResponse pointers
// and an error if something went wrong.
func (c *HTTPclient) GetCensusesByStrategy(strategyID uint64) ([]*api.GetCensusResponse, error) {
	if strategyID == 0 {
		return nil, fmt.Errorf("%w: strategyID is required", ErrBadInputs)
	}
	// construct the URL to the API with the strategyID
	endpoint := fmt.Sprintf(GetCensusesByStrategyURI, strategyID)
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
		return nil, fmt.Errorf("%w: %w", ErrNoStatusOk, fmt.Errorf("%d %s", res.StatusCode, http.StatusText(res.StatusCode)))
	}
	// decode the response and return the censuses
	response := &api.GetCensusesResponse{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodingResponse, err)
	}
	return response.Censuses, nil
}
