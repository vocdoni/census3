package apiclient

import (
	"fmt"
	"net/url"
	"path"
)

// constructURL constructs a URL from the base URL, endpoint and additional path elements
func (c *HTTPclient) constructURL(endpoint string, args ...string) (string, error) {
	u, err := url.Parse(c.addr.String())
	if err != nil {
		return "", fmt.Errorf("error parsing base URL: %v", err)
	}
	allPaths := append([]string{endpoint}, args...)
	u.Path = path.Join(append([]string{u.Path}, allPaths...)...)
	return u.String(), nil
}
