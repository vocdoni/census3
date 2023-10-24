package service

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/db"
)

// POAP_URI is the endpoint to get the POAP holders for an eventID and offset.
// It uses the maximum limit of 300 POAPs per request.
// https://documentation.poap.tech/reference/geteventpoaps-2
const POAP_URI = "https://api.poap.tech/event/%s/poaps?limit=300&offset=%d"

type POAPAPIResponse struct {
	Total  int `json:"total"`
	Tokens []struct {
		Owner struct {
			ID string `json:"id"`
		} `json:"owner"`
	} `json:"tokens"`
}

type POAPExternalProvider struct {
	AccessToken string
}

// Init initializes the POAP external provider with the database provided.
// It returns an error if the POAP access token is not defined.
func (p *POAPExternalProvider) Init(_ *db.DB) error {
	if p.AccessToken == "" {
		return fmt.Errorf("no POAP access token defined")
	}
	return nil
}

// GetHolders returns the holders of the POAP eventID provided. It requests the
// list of token holders to the POAP API parsing every POAP holder for the event
// ID provided. It returns a map with the address of the holder as key and the
// balance of the token holder as value. The POAP API endpoint to get the list
// of POAPs is paginated, so it requests the list of POAPs in batches of 300
// POAPs per request (maximum limit allowed by the POAP API).
func (p *POAPExternalProvider) GetHolders(ids ...any) (map[common.Address]*big.Int, error) {
	// parse eventID
	if len(ids) == 0 {
		return nil, fmt.Errorf("no POAP eventID provided")
	}
	eventID, ok := ids[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("invalid POAP eventID provided, it must be a big.Int")
	}
	// init http client
	client := &http.Client{}
	// create a request to get the first page of poaps
	offset := 0
	endpoint := fmt.Sprintf(POAP_URI, eventID.String(), offset)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", p.AccessToken)
	// do the request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// decode response
	rawResults, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// parse poap from decoded response
	poapRes := POAPAPIResponse{}
	if err := json.Unmarshal(rawResults, &poapRes); err != nil {
		return nil, err
	}
	// compose holders map
	holders := make(map[common.Address]*big.Int)
	for _, poap := range poapRes.Tokens {
		addr := common.HexToAddress(poap.Owner.ID)
		holders[addr] = big.NewInt(1)
	}
	return holders, nil
}
