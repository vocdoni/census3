package apiclient

import (
	"fmt"
	"net/url"
	"strings"
)

const queryParamsSeparator = "?"

// constructURL constructs a URL from the base URL, endpoint and additional path elements
func (c *HTTPclient) constructURL(endpoint string) (string, error) {
	if endpoint == "" {
		return c.addr.String(), nil
	}
	// decode endpoint in paths and query parameters
	parts := strings.SplitN(endpoint, queryParamsSeparator, 2)
	finalEndpoint := c.addr.JoinPath(parts[0])
	if len(parts) < 2 {
		return finalEndpoint.String(), nil
	}
	queryParams, err := url.ParseQuery(parts[1])
	if err != nil {
		return "", fmt.Errorf("error parsing query parameters: %w", err)
	}
	finalEndpoint.RawQuery = queryParams.Encode()
	return finalEndpoint.String(), nil
}
