package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"go.vocdoni.io/dvote/log"
)

const (
	// POAP_SYMBOL_PREFIX is the prefix of the POAP token symbol to be used in
	// with the eventID to compose the token symbol.
	POAP_SYMBOL_PREFIX = "POAP"
	// POAP_MAX_LIMIT is the maximum limit of 300 POAPs per request.
	// https://documentation.poap.tech/reference/geteventpoaps-2
	POAP_MAX_LIMIT = 300
	// EVENT_URI is the endpoint to get the event info for an eventID.
	EVENT_URI = "/events/id/%s"
	// POAP_URI is the endpoint to get the POAP holders for an eventID.
	POAP_URI = "/event/%s/poaps"
	// POAP_CONTRACT_ADDRESS is the address of the POAP contract.
	POAP_CONTRACT_ADDRESS = "0x22c1f6050e56d2876009903609a2cc3fef83b415"
)

// EventAPIResponse is the struct that stores the response of the POAP API
// endpoint to get the event info for an event ID.
type EventAPIResponse struct {
	FancyID string `json:"fancy_id"`
	Name    string `json:"name"`
}

// POAPAPIResponse is the struct that stores the response of the POAP API
// endpoint to get the list of POAPs for an event ID.
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

// Decimals method is not implemented in the POAP external provider. By default
// it returns 0 and nil error.
func (p *POAPHolderProvider) Decimals(_ context.Context, _ []byte) (uint64, error) {
	return 0, nil
}

// TotalSupply method is not implemented in the POAP external provider. By
// default it returns 0 and nil error.
func (p *POAPHolderProvider) TotalSupply(_ context.Context, _ []byte) (*big.Int, error) {
	return big.NewInt(0), nil
}

// BlockTimestamp method is not implemented in the POAP external provider. By
// default it returns an empty string and nil error.
func (p *POAPHolderProvider) BlockTimestamp(_ context.Context, _ uint64) (string, error) {
	return "", nil
}

// BlockRootHash method is not implemented in the POAP external provider. By
// default it returns an empty bytes slice and nil error.
func (p *POAPHolderProvider) BlockRootHash(_ context.Context, _ uint64) ([]byte, error) {
	return []byte{}, nil
}

// CreationBlock method is not implemented in the POAP external provider. By
// default it returns 0 and nil error.
func (p *POAPHolderProvider) CreationBlock(_ context.Context, _ []byte) (uint64, error) {
	return 0, nil
}

// Close method is not implemented in the POAP external provider. By default it
// returns nil error.
func (p *POAPHolderProvider) Close() error {
	return nil
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
	newSnapshot, err := p.lastHolders(eventID)
	if err != nil {
		return nil, err
	}
	// calculate snapshot from
	from := delta
	if snapshot, exist := p.snapshots[string(id)]; exist {
		from += snapshot.from
	}
	// calculate partials balances
	partialBalances := p.calcPartials(eventID, newSnapshot)
	// save snapshot
	p.snapshots[string(id)] = &POAPSnapshot{
		from:     from,
		snapshot: newSnapshot,
	}
	// return partials from last snapshot
	return partialBalances, nil
}

func (p *POAPHolderProvider) Address(_ context.Context, _ []byte) (common.Address, error) {
	return common.HexToAddress(POAP_CONTRACT_ADDRESS), nil
}

func (p *POAPHolderProvider) Name(_ context.Context, id []byte) (string, error) {
	info, err := p.getEventInfo(string(id))
	if err != nil {
		return "", err
	}
	return info.Name, nil
}

func (p *POAPHolderProvider) Symbol(_ context.Context, id []byte) (string, error) {
	info, err := p.getEventInfo(string(id))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", POAP_SYMBOL_PREFIX, info.FancyID), nil
}

func (p *POAPHolderProvider) BalanceOf(_ context.Context, id []byte, addr common.Address) (*big.Int, error) {
	// parse eventID from id
	eventID := string(id)
	// get the last stored snapshot
	if snapshot, exist := p.snapshots[eventID]; exist {
		if balance, exist := snapshot.snapshot[addr]; exist {
			return balance, nil
		}
		return nil, fmt.Errorf("no balance found for address %s", addr.String())
	}
	return nil, fmt.Errorf("no snapshot found for eventID %s", eventID)
}

func (p *POAPHolderProvider) LatestBlockNumber(_ context.Context, id []byte) (uint64, error) {
	// parse eventID from id
	eventID := string(id)
	// get the last stored snapshot
	if snapshot, exist := p.snapshots[eventID]; exist {
		return snapshot.from, nil
	}
	return 0, fmt.Errorf("no snapshot found for eventID %s", eventID)
}

// lastHolders returns the holders of the POAP eventID provided. It requests the
// list of token holders to the POAP API parsing every POAP holder for the event
// ID provided. It returns a map with the address of the holder as key and the
// balance of the token holder as value. The POAP API endpoint to get the list
// of POAPs is paginated, so it requests the list of POAPs in batches of 300
// POAPs per request (maximum limit allowed by the POAP API).
func (p *POAPHolderProvider) lastHolders(eventID string) (map[common.Address]*big.Int, error) {
	holders := make(map[common.Address]*big.Int)
	offset, total := 0, POAP_MAX_LIMIT+1
	for offset < total {
		// get holders page based on offset
		poapRes, err := p.holdersPage(eventID, offset)
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

// holdersPage returns the holders of the POAP eventID provided for the given
// offset. It returns a POAPAPIResponse struct with the list of POAPs for the
// eventID and the total number of POAPs for the eventID. Every POAP in the
// list contains the address of the token holder.
func (p *POAPHolderProvider) holdersPage(eventID string, offset int) (*POAPAPIResponse, error) {
	// compose the endpoint for the request
	strURL, err := url.JoinPath(p.URI, fmt.Sprintf(POAP_URI, eventID))
	if err != nil {
		return nil, err
	}
	endpoint, err := url.Parse(strURL)
	if err != nil {
		return nil, err
	}
	q := endpoint.Query()
	q.Add("limit", fmt.Sprint(POAP_MAX_LIMIT))
	q.Add("offset", fmt.Sprint(offset))
	endpoint.RawQuery = q.Encode()
	// create request and add headers
	req, err := http.NewRequest("GET", endpoint.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", p.AccessToken)
	// do the request
	res, err := http.DefaultClient.Do(req)
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

// getEventInfo returns the event info for the given eventID. It makes a request
// to the POAP API endpoint to get the event info for the eventID provided and
// returns an EventAPIResponse struct with the event info.
func (p *POAPHolderProvider) getEventInfo(eventID string) (*EventAPIResponse, error) {
	// compose the endpoint for the request
	strURL, err := url.JoinPath(p.URI, fmt.Sprintf(EVENT_URI, eventID))
	if err != nil {
		return nil, err
	}
	endpoint, err := url.Parse(strURL)
	if err != nil {
		return nil, err
	}
	// create request and add headers
	req, err := http.NewRequest("GET", endpoint.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", p.AccessToken)
	// do the request
	res, err := http.DefaultClient.Do(req)
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
	eventRes := EventAPIResponse{}
	if err := json.Unmarshal(rawResults, &eventRes); err != nil {
		return nil, err
	}
	return &eventRes, nil
}

// calcPartials calculates the partials balances of the token holders for the
// given eventID and new snapshot. It returns a map with the address of the
// holder as key and the balance of the token holder as value. The partials
// balances will include:
//   - holders from the new snapshot that are not in the current snapshot with
//     the balance of the new snapshot
//   - holders from the current snapshot that are not in the new snapshot but
//     with zero balance
//   - holders from the current snapshot that are in the new snapshot with the
//     balance of the new snapshot if the balance has changed
func (p *POAPHolderProvider) calcPartials(eventID string, newSnapshot map[common.Address]*big.Int) map[common.Address]*big.Int {
	// get current snapshot if exists
	currentSnapshot := make(map[common.Address]*big.Int)
	if current, exist := p.snapshots[eventID]; exist {
		currentSnapshot = current.snapshot
	}
	// calculate partials balances from current and new snapshots
	partialsBalances := make(map[common.Address]*big.Int)
	for addr, balance := range newSnapshot {
		if currentBalance, exist := currentSnapshot[addr]; !exist || currentBalance.Cmp(balance) != 0 {
			partialsBalances[addr] = balance
		}
	}
	// add zero balances for holders in current snapshot but not in new snapshot
	for addr, currentBalance := range currentSnapshot {
		if _, exist := newSnapshot[addr]; !exist {
			partialsBalances[addr] = new(big.Int).Neg(currentBalance)
		}
	}
	return partialsBalances
}
