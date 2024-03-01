package apiclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/vocdoni/census3/api"
)

// GetTokens method returns a list of tokens from the API, it accepts a
// pageSize, nextCursor and prevCursor. If the pageSize is -1 and cursors are
// empty, it will return all the tokens. If somethign goes wrong, it will return
// an error.
func (c *HTTPclient) GetTokens(pageSize int, nextCursor, prevCursor string) ([]*api.GetTokensItemResponse, error) {
	// construct the URL to the API with the pageSize, nextCursor and prevCursor
	endpoint := fmt.Sprintf(GetTokensURI, pageSize, nextCursor, prevCursor)
	u, err := c.constructURL(endpoint)
	if err != nil {
		return nil, errors.Join(ErrConstructingURL, err)
	}
	// create the request and send it, if there is an error or the status code
	// is not 200, return an error
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, errors.Join(ErrCreatingRequest, err)
	}
	res, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Join(ErrMakingRequest, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.Join(ErrNoStatusOk, fmt.Errorf("%d %s", res.StatusCode, http.StatusText(res.StatusCode)))
	}
	// decode the response and return the tokens
	tokensResponse := &api.GetTokensResponse{}
	if err := json.NewDecoder(res.Body).Decode(&tokensResponse); err != nil {
		return nil, errors.Join(ErrDecodingResponse, err)
	}
	return tokensResponse.Tokens, nil
}

// GetToken method returns a token from the API, it accepts the tokenID, chainID
// and externalID. If something goes wrong, it will return an error.
func (c *HTTPclient) GetToken(tokenID string, chainID uint64, externalID string) (*api.GetTokenResponse, error) {
	if tokenID == "" || chainID == 0 {
		return nil, fmt.Errorf("tokenID and chainID are required")
	}
	// construct the URL to the API with the tokenID, chainID and externalID
	// provided
	endpoint := fmt.Sprintf(GetAndDeleteTokenURI, tokenID, chainID, externalID)
	u, err := c.constructURL(endpoint)
	if err != nil {
		return nil, errors.Join(ErrConstructingURL, err)
	}
	// create the request and send it, if there is an error or the status code
	// is not 200, return an error
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, errors.Join(ErrCreatingRequest, err)
	}
	res, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Join(ErrMakingRequest, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.Join(ErrNoStatusOk, fmt.Errorf("%d %s", res.StatusCode, http.StatusText(res.StatusCode)))
	}
	// decode the response and return the token
	tokenResponse := &api.GetTokenResponse{}
	if err := json.NewDecoder(res.Body).Decode(&tokenResponse); err != nil {
		return nil, errors.Join(ErrDecodingResponse, err)
	}
	return tokenResponse, nil
}

// CreateToken method creates a token in the API, it accepts a token to be
// created. The minimun required token attributes are: ID, type and chainID.
// Optional attributes: externalID, tags. If something goes wrong, it will
// return an error.
func (c *HTTPclient) CreateToken(token *api.CreateTokenRequest) error {
	if token.ID == "" || token.Type == "" || token.ChainID == 0 {
		return fmt.Errorf("ID, Type and ChainID are required")
	}
	// construct the URL to the API
	u, err := c.constructURL(CreateTokensURI)
	if err != nil {
		return errors.Join(ErrConstructingURL, err)
	}
	// encode the input token to JSON to be sent in the request body
	reqBody, err := json.Marshal(token)
	if err != nil {
		return errors.Join(ErrEncodingRequest, err)
	}
	// create the request and send it with the encoded body, if there is an
	// error or the status code is not 201, return an error
	req, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(reqBody))
	if err != nil {
		return errors.Join(ErrCreatingRequest, err)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.c.Do(req)
	if err != nil {
		return errors.Join(ErrMakingRequest, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		return errors.Join(ErrNoStatusOk, fmt.Errorf("%d %s", res.StatusCode, http.StatusText(res.StatusCode)))
	}
	return nil
}

// DeleteToken method deletes a token in the API, it accepts the tokenID,
// chainID and externalID. If something goes wrong, it will return an error. It
// returns the queueID, that can be used to check the status of the deletion
// process using the DeleteTokenQueue method.
func (c *HTTPclient) DeleteToken(tokenID string, chainID uint64, externalID string) (string, error) {
	if tokenID == "" || chainID == 0 {
		return "", fmt.Errorf("tokenID and chainID are required")
	}
	// construct the URL to the API with the tokenID, chainID and externalID
	// provided
	endpoint := fmt.Sprintf(GetAndDeleteTokenURI, tokenID, chainID, externalID)
	u, err := c.constructURL(endpoint)
	if err != nil {
		return "", errors.Join(ErrConstructingURL, err)
	}
	// create the request and send it, if there is an error or the status code
	// is not 200, return an error
	req, err := http.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return "", errors.Join(ErrCreatingRequest, err)
	}
	res, err := c.c.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making DELETE request: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", errors.Join(ErrNoStatusOk, fmt.Errorf("%d %s", res.StatusCode, http.StatusText(res.StatusCode)))
	}
	// decode the response and return the queue response
	queueResponse := &api.QueueResponse{}
	if err := json.NewDecoder(res.Body).Decode(&queueResponse); err != nil {
		return "", errors.Join(ErrDecodingResponse, err)
	}
	return queueResponse.QueueID, nil
}

// DeleteTokenQueue method checks the status of the deletion process of a token
// in the API, it accepts the tokenID, queueID, chainID and externalID. If
// something goes wrong, it will return an error. It returns the delete queue
// response.
func (c *HTTPclient) DeleteTokenQueue(tokenID string, chainID uint64, externalID, queueID string) (*api.DeleteTokenQueueResponse, error) {
	// construct the URL to the API with the tokenID, queueID, chainID and
	// externalID provided
	endpoint := fmt.Sprintf(DeleteTokenQueueURI, tokenID, queueID, chainID, externalID)
	u, err := c.constructURL(endpoint)
	if err != nil {
		return nil, errors.Join(ErrConstructingURL, err)
	}
	// create the request and send it, if there is an error or the status code
	// is not 200, return an error
	req, err := http.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, errors.Join(ErrCreatingRequest, err)
	}
	res, err := c.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making DELETE request: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.Join(ErrNoStatusOk, fmt.Errorf("%d %s", res.StatusCode, http.StatusText(res.StatusCode)))
	}
	// decode the response and return the queue response
	queueResponse := &api.DeleteTokenQueueResponse{}
	if err := json.NewDecoder(res.Body).Decode(&queueResponse); err != nil {
		return nil, errors.Join(ErrDecodingResponse, err)
	}
	return queueResponse, nil
}

// GetTokenHolder method returns the balance of a token holder from the API, it
// accepts the tokenID, chainID, externalID and holderID. If something goes
// wrong, it will return an error.
func (c *HTTPclient) GetTokenHolder(tokenID string, chainID uint64, externalID, holderID string) (*big.Int, error) {
	// construct the URL to the API with the tokenID, holderID, chainID and
	// externalID provided
	endpoint := fmt.Sprintf(GetTokenHolderURI, tokenID, holderID, chainID, externalID)
	u, err := c.constructURL(endpoint)
	if err != nil {
		return nil, errors.Join(ErrConstructingURL, err)
	}
	// create the request and send it, if there is an error or the status code
	// is not 200, return an error
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, errors.Join(ErrCreatingRequest, err)
	}
	res, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Join(ErrMakingRequest, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.Join(ErrNoStatusOk, fmt.Errorf("%d %s", res.StatusCode, http.StatusText(res.StatusCode)))
	}
	// decode the response and return the token holder response
	holderResponse := &api.GetTokenHolderResponse{}
	if err := json.NewDecoder(res.Body).Decode(&holderResponse); err != nil {
		return nil, errors.Join(ErrDecodingResponse, err)
	}
	bBalance, ok := new(big.Int).SetString(holderResponse.Balance, 10)
	if !ok {
		return nil, fmt.Errorf("error parsing balance: %v", holderResponse.Balance)
	}
	return bBalance, nil
}

// GetTokenTypes method returns the supported token types from the API. If
// something goes wrong, it will return an error. It returns the supported token
// types as a slice of strings.
func (c *HTTPclient) GetTokenTypes() ([]string, error) {
	// construct the URL to the API
	u, err := c.constructURL(GetTokenTypes)
	if err != nil {
		return nil, errors.Join(ErrConstructingURL, err)
	}
	// create the request and send it, if there is an error or the status code
	// is not 200, return an error
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, errors.Join(ErrCreatingRequest, err)
	}
	res, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Join(ErrMakingRequest, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.Join(ErrNoStatusOk, fmt.Errorf("%d %s", res.StatusCode, http.StatusText(res.StatusCode)))
	}
	// decode the response and return the token types
	tokenTypesResponse := &api.TokenTypesResponse{}
	if err := json.NewDecoder(res.Body).Decode(&tokenTypesResponse); err != nil {
		return nil, errors.Join(ErrDecodingResponse, err)
	}
	return tokenTypesResponse.SupportedTypes, nil
}
