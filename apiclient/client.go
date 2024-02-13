package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/vocdoni/census3/api"
	"go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

const (
	// HTTPGET is the method string used for calling Request()
	HTTPGET = http.MethodGet
	// HTTPPOST is the method string used for calling Request()
	HTTPPOST = http.MethodPost
	// HTTPDELETE is the method string used for calling
	HTTPDELETE = http.MethodDelete
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// HTTPclient is the Census3 API HTTP client.
type HTTPclient struct {
	c               Doer
	token           *uuid.UUID
	addr            *url.URL
	supportedChains []string
}

// NewHTTPclient creates a new HTTP(s) API Census3 client.
func NewHTTPclient(addr *url.URL, bearerToken *uuid.UUID) (*HTTPclient, error) {
	tr := &http.Transport{
		IdleConnTimeout:    10 * time.Second,
		DisableCompression: false,
		WriteBufferSize:    1 * 1024 * 1024, // 1 MiB
		ReadBufferSize:     1 * 1024 * 1024, // 1 MiB
	}
	c := &HTTPclient{
		c:     &http.Client{Transport: tr, Timeout: time.Second * 8},
		token: bearerToken,
		addr:  addr,
	}
	data, status, err := c.Request(HTTPGET, nil, "info")
	if err != nil {
		return nil, err
	}
	if status != apirest.HTTPstatusOK {
		log.Warnw("cannot get info from API server", "status", status, "data", data)
		return c, nil
	}
	info := &api.APIInfo{}
	if err := json.Unmarshal(data, info); err != nil {
		return nil, fmt.Errorf("cannot get API info from API server")
	}
	for _, chain := range info.SupportedChains {
		c.supportedChains = append(c.supportedChains, chain.Name)
	}
	return c, nil
}

// SetAuthToken configures the bearer authentication token.
func (c *HTTPclient) SetAuthToken(token *uuid.UUID) {
	c.token = token
}

// BearerToken returns the current bearer authentication token.
func (c *HTTPclient) BearerToken() *uuid.UUID {
	return c.token
}

// SetHostAddr configures the host address of the API server.
func (c *HTTPclient) SetHostAddr(addr *url.URL) error {
	c.addr = addr
	data, status, err := c.Request(HTTPGET, nil, "chain", "info")
	if err != nil {
		return err
	}
	if status != apirest.HTTPstatusOK {
		return fmt.Errorf("API error: %d (%s)", status, data)
	}
	return nil
}

// Request performs a `method` type raw request to the endpoint specified in urlPath parameter.
// Method is either GET or POST. If POST, a JSON struct should be attached.  Returns the response,
// the status code and an error.
func (c *HTTPclient) Request(method string, jsonBody any, urlPath ...string) ([]byte, int, error) {
	body, err := json.Marshal(jsonBody)
	if err != nil {
		return nil, 0, err
	}
	u, err := url.Parse(c.addr.String())
	if err != nil {
		return nil, 0, err
	}
	u.Path = path.Join(u.Path, path.Join(urlPath...))
	headers := http.Header{}
	if c.token != nil {
		headers = http.Header{
			"Authorization": []string{"Bearer " + c.token.String()},
			"User-Agent":    []string{"Census3 API client / 1.0"},
			"Content-Type":  []string{"application/json"},
		}
	}

	log.Debugw("http request", "type", method, "path", u.Path, "body", jsonBody)
	var resp *http.Response

	resp, err = c.c.Do(&http.Request{
		Method: method,
		URL:    u,
		Header: headers,
		Body: func() io.ReadCloser {
			if jsonBody == nil {
				return nil
			}
			return io.NopCloser(bytes.NewBuffer(body))
		}(),
	})

	if err != nil {
		return nil, 0, err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	return data, resp.StatusCode, nil
}
