package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"path"

	"github.com/ethereum/go-ethereum/common"
	"go.vocdoni.io/dvote/log"
)

const (
	// POAP_MAX_LIMIT is the maximum limit of 300 POAPs per request.
	// https://documentation.poap.tech/reference/geteventpoaps-2
	POAP_MAX_LIMIT = 300
	// POAP_URI is the endpoint to get the POAP holders for an eventID, offset
	// and limit.
	POAP_URI = "/event/%s/poaps?limit=%d&offset=%d"
)

type POAPAPIResponse struct {
	Total  int `json:"total"`
	Tokens []struct {
		Owner struct {
			ID string `json:"id"`
		} `json:"owner"`
	} `json:"tokens"`
}

// POAPSnapshot is the struct that stores the snapshot of the POAP holders for
// an event ID and from point in time.
type POAPSnapshot struct {
	from     uint64
	snapshot map[common.Address]*big.Int
}

// POAPHolderProvider is the struct that implements the HolderProvider interface
// to get the balances of the POAP token holders for an event ID. It uses the
// POAP API to get the list of POAPs for an event ID and calculate the balances
// of the token holders from the last snapshot.
type POAPHolderProvider struct {
	URI         string
	AccessToken string
	snapshots   map[string]*POAPSnapshot
}

// Init initializes the POAP external provider with the database provided.
// It returns an error if the POAP access token or api endpoint uri is not
// defined.
func (p *POAPHolderProvider) Init() error {
	if p.URI == "" {
		return fmt.Errorf("no POAP URI defined")
	}
	if p.AccessToken == "" {
		return fmt.Errorf("no POAP access token defined")
	}
	p.snapshots = make(map[string]*POAPSnapshot)
	return nil
}

// SetLastBalances sets the balances of the token holders for the given id and
// from point in time and store it in a snapshot.
func (p *POAPHolderProvider) SetLastBalances(_ context.Context, id []byte,
	balances map[common.Address]*big.Int, from uint64,
) error {
	p.snapshots[string(id)] = &POAPSnapshot{
		from:     from,
		snapshot: balances,
	}
	return nil
}

// HoldersBalances returns the balances of the token holders for the given id
// and delta point in time. It requests the list of token holders to the POAP
// API parsing every POAP holder for the event ID provided and calculate the
// balances of the token holders from the last snapshot.
func (p *POAPHolderProvider) HoldersBalances(_ context.Context, id []byte, delta uint64) (map[common.Address]*big.Int, error) {
	// parse eventID from id
	eventID := string(id)
	// get last snapshot
	newSnapshot, err := p.getLastHolders(eventID)
	if err != nil {
		return nil, err
	}
	// calculate snapshot from
	from := delta
	if snapshot, exist := p.snapshots[string(id)]; exist {
		from += snapshot.from
	}
	// save snapshot
	p.snapshots[string(id)] = &POAPSnapshot{
		from:     from,
		snapshot: newSnapshot,
	}
	// calculate and return partials from last snapshot
	return p.calcPartials(eventID, newSnapshot), nil
}

// Close method is not implemented for the POAP external provider.
func (p *POAPHolderProvider) Close() error {
	// not implemented
	return nil
}

// getLastHolders returns the holders of the POAP eventID provided. It requests the
// list of token holders to the POAP API parsing every POAP holder for the event
// ID provided. It returns a map with the address of the holder as key and the
// balance of the token holder as value. The POAP API endpoint to get the list
// of POAPs is paginated, so it requests the list of POAPs in batches of 300
// POAPs per request (maximum limit allowed by the POAP API).
func (p *POAPHolderProvider) getLastHolders(eventID string) (map[common.Address]*big.Int, error) {
	holders := make(map[common.Address]*big.Int)
	offset, total := 0, POAP_MAX_LIMIT+1
	for offset < total {
		// get holders page based on offset
		poapRes, err := p.getHoldersPage(eventID, offset)
		if err != nil {
			return nil, err
		}
		// add holders to map
		for _, poap := range poapRes.Tokens {
			addr := common.HexToAddress(poap.Owner.ID)
			holders[addr] = big.NewInt(1)
		}
		// update offset and total
		offset += POAP_MAX_LIMIT
		total = poapRes.Total
	}
	return holders, nil
}

// getHoldersPage returns the holders of the POAP eventID provided for the
// given offset. It returns a POAPAPIResponse struct with the list of POAPs
// for the eventID and the total number of POAPs for the eventID. Every POAP
// in the list contains the address of the token holder.
func (p *POAPHolderProvider) getHoldersPage(eventID string, offset int) (*POAPAPIResponse, error) {
	// init http client
	client := &http.Client{}
	// create a request to get the current page of POAPs
	endpoint := path.Join(p.URI, fmt.Sprintf(POAP_URI, eventID, POAP_MAX_LIMIT, offset))
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
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Error(err)
		}
	}()
	// parse poap from decoded response
	poapRes := POAPAPIResponse{}
	if err := json.Unmarshal(rawResults, &poapRes); err != nil {
		return nil, err
	}
	return &poapRes, nil
}

func (p *POAPHolderProvider) calcPartials(eventID string, newSnapshot map[common.Address]*big.Int) map[common.Address]*big.Int {
	// get current snapshot if exists
	currentSnapshot := make(map[common.Address]*big.Int)
	if current, exist := p.snapshots[eventID]; exist {
		currentSnapshot = current.snapshot
	}
	// the resulting partials will include:
	//  * holders from the new snapshot that are not in the current snapshot
	//    with the balance of the new snapshot
	//  * holders from the current snapshot that are not in the new snapshot
	//    but with negative balance
	//  * holders from the current snapshot that are in the new snapshot with
	//    the difference between the balances of the new and current snapshot
	partialsBalances := make(map[common.Address]*big.Int)
	for addr, balance := range newSnapshot {
		partialsBalances[addr] = balance
	}
	for addr, balance := range currentSnapshot {
		if newBalance, exist := newSnapshot[addr]; !exist {
			partialsBalances[addr] = new(big.Int).Neg(balance)
		} else {
			partialsBalances[addr] = new(big.Int).Sub(newBalance, balance)
		}
	}
	return partialsBalances
}
