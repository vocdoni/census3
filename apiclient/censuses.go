package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vocdoni/census3/api"
)

// GetCensus makes a GET request to the /censuses/{censusID} endpoint and returns the census data.
func (c *HTTPclient) GetCensus(censusID uint64) (*api.GetCensusResponse, error) {
	endpoint := fmt.Sprintf("/censuses/%d", censusID)
	u, err := c.constructURL(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error constructing URL: %w", err)
	}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned non-OK status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	response := api.GetCensusResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &response, nil
}

// CreateCensus sends a POST request to create a new census and returns the queue ID
func (c *HTTPclient) CreateCensus(request api.CreateCensusRequest) (*api.QueueResponse, error) {
	url, err := c.constructURL("/censuses")
	if err != nil {
		return nil, fmt.Errorf("error constructing URL: %v", err)
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}

	req, err := http.NewRequest(HTTPPOST, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned non-OK status: %d", resp.StatusCode)
	}

	response := api.QueueResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &response, nil
}

// GetCensusQueue makes a GET request to the /censuses/queue/{queueID} endpoint
// and returns the current status of the census creation process.
func (c *HTTPclient) GetCensusQueue(queueID string) (*api.CensusQueueResponse, error) {
	endpoint := fmt.Sprintf("/censuses/queue/%s", queueID)
	u, err := c.constructURL(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error constructing URL: %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned non-OK status: %d", resp.StatusCode)
	}

	response := &api.CensusQueueResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return response, nil
}

// GetStrategyCensuses makes a GET request to the /censuses/strategy/{strategyID} endpoint
// and returns the censuses generated with the given strategy ID.
func (c *HTTPclient) GetStrategyCensuses(strategyID uint64) (*api.GetCensusesResponse, error) {
	endpoint := fmt.Sprintf("/censuses/strategy/%d", strategyID)
	u, err := c.constructURL(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error constructing URL: %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned non-OK status: %d", resp.StatusCode)
	}

	response := &api.GetCensusesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return response, nil
}
