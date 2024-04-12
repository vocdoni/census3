package poap

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/scanner/providers"
	"go.vocdoni.io/dvote/log"
)

const (
	defaultRequestTimeout = 30 * time.Second
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
	IconURI string `json:"image_url"`
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
	ctx          context.Context
	cancel       context.CancelFunc
	apiEndpoint  string
	accessToken  string
	snapshots    map[string]*POAPSnapshot
	snapshotsMtx *sync.RWMutex
}

type POAPConfig struct {
	APIEndpoint string
	AccessToken string
}

// Init initializes the POAP external provider with the database provided.
// It returns an error if the POAP access token or api endpoint uri is not
// defined.
func (p *POAPHolderProvider) Init(globalCtx context.Context, iconf any) error {
	// parse config
	conf, ok := iconf.(POAPConfig)
	if !ok {
		return fmt.Errorf("bad config type, it must be a POAPConfig struct")
	}
	if conf.APIEndpoint == "" {
		return fmt.Errorf("no POAP URI defined")
	}
	if conf.AccessToken == "" {
		return fmt.Errorf("no POAP access token defined")
	}
	p.ctx, p.cancel = context.WithCancel(globalCtx)
	p.apiEndpoint = conf.APIEndpoint
	p.accessToken = conf.AccessToken
	p.snapshots = make(map[string]*POAPSnapshot)
	p.snapshotsMtx = &sync.RWMutex{}
	return nil
}

// SetRef method is not implemented in the POAP external provider. By default it
// returns nil error.
func (p *POAPHolderProvider) SetRef(_ any) error {
	return nil
}

// SetLastBlockNumber method is not implemented in the POAP external provider.
func (p *POAPHolderProvider) SetLastBlockNumber(_ uint64) {}

// SetLastBalances sets the balances of the token holders for the given id and
// from point in time and store it in a snapshot.
func (p *POAPHolderProvider) SetLastBalances(_ context.Context, id []byte,
	balances map[common.Address]*big.Int, from uint64,
) error {
	p.snapshotsMtx.Lock()
	defer p.snapshotsMtx.Unlock()
	p.snapshots[string(id)] = &POAPSnapshot{
		from:     from,
		snapshot: balances,
	}
	log.Debugw("last balances stored", "balances", len(balances))
	return nil
}

// HoldersBalances returns the balances of the token holders for the given id
// and delta point in time. It requests the list of token holders to the POAP
// API parsing every POAP holder for the event ID provided and calculate the
// balances of the token holders from the last snapshot.
func (p *POAPHolderProvider) HoldersBalances(_ context.Context, id []byte, delta uint64) (
	map[common.Address]*big.Int, uint64, uint64, bool, *big.Int, error,
) {
	// parse eventID from id
	eventID := string(id)
	log.Infow("getting POAP holders balances", "eventID", eventID)
	// get last snapshot
	newSnapshot, err := p.lastHolders(eventID)
	if err != nil {
		return nil, 0, 0, false, big.NewInt(0), err
	}
	p.snapshotsMtx.RLock()
	defer p.snapshotsMtx.RUnlock()
	// if there is no snapshot, the final snapshot is the new snapshot, otherwise
	// calculate the partials balances from the last snapshot
	from := delta
	finalSnapshot := newSnapshot
	if currentSnapshot, exist := p.snapshots[eventID]; exist {
		finalSnapshot = providers.CalcPartialHolders(currentSnapshot.snapshot, newSnapshot)
		from += currentSnapshot.from
	}
	// store the new snapshot
	p.snapshots[string(id)] = &POAPSnapshot{
		from:     from,
		snapshot: newSnapshot,
	}
	// calculate total supply
	totalSupply := new(big.Int)
	for _, balance := range finalSnapshot {
		totalSupply.Add(totalSupply, balance)
	}
	// return the final snapshot
	return finalSnapshot, uint64(len(finalSnapshot)), from, true, totalSupply, nil
}

// Close method is not implemented in the POAP external provider. By default it
// returns nil error.
func (p *POAPHolderProvider) Close() error {
	p.cancel()
	return nil
}

// IsExternal method returns true because the POAP provider is an external
// provider.
func (p *POAPHolderProvider) IsExternal() bool {
	return true
}

// IsSynced returns true if the POAP external provider has a snapshot for the
// given id.
func (p *POAPHolderProvider) IsSynced(externalID []byte) bool {
	_, exist := p.snapshots[string(externalID)]
	return exist
}

// Address returns the address of the POAP token.
func (p *POAPHolderProvider) Address(_ []byte) common.Address {
	return common.HexToAddress(POAP_CONTRACT_ADDRESS)
}

// Type returns the type of the POAP token. By default it returns the
// CONTRACT_TYPE_POAP.
func (p *POAPHolderProvider) Type() uint64 {
	return providers.CONTRACT_TYPE_POAP
}

// TypeName returns the type name of the POAP token. By default it returns the
// string "POAP".
func (p *POAPHolderProvider) TypeName() string {
	return providers.TokenTypeName(providers.CONTRACT_TYPE_POAP)
}

// ChainID method is not implemented in the POAP external provider. By default 1.
func (p *POAPHolderProvider) ChainID() uint64 {
	return 1
}

// Name returns the name of the POAP token. It makes a request to the POAP API
// endpoint to get the event info for the eventID provided and returns the name
// of the event.
func (p *POAPHolderProvider) Name(id []byte) (string, error) {
	info, err := p.getEventInfo(string(id))
	if err != nil {
		return "", err
	}
	return info.Name, nil
}

// Symbol returns the symbol of the POAP token. It makes a request to the POAP
// API endpoint to get the event info for the eventID provided and returns the
// symbol of the event, which is composed by the prefix POAP_SYMBOL_PREFIX and
// the fancyID of the event.
func (p *POAPHolderProvider) Symbol(id []byte) (string, error) {
	info, err := p.getEventInfo(string(id))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", POAP_SYMBOL_PREFIX, info.FancyID), nil
}

// Decimals method is not implemented in the POAP external provider. By default
// it returns 0 and nil error.
func (p *POAPHolderProvider) Decimals(_ []byte) (uint64, error) {
	return 0, nil
}

// TotalSupply method is not implemented in the POAP external provider. By
// default it returns 0 and nil error.
func (p *POAPHolderProvider) TotalSupply(id []byte) (*big.Int, error) {
	p.snapshotsMtx.RLock()
	defer p.snapshotsMtx.RUnlock()
	totalSupply := new(big.Int)
	if snapshot, exist := p.snapshots[string(id)]; exist {
		for _, balance := range snapshot.snapshot {
			totalSupply.Add(totalSupply, balance)
		}
		return totalSupply, nil
	}
	return big.NewInt(0), nil
}

func (p *POAPHolderProvider) BalanceOf(addr common.Address, id []byte) (*big.Int, error) {
	// parse eventID from id
	eventID := string(id)
	// get the last stored snapshot
	p.snapshotsMtx.RLock()
	defer p.snapshotsMtx.RUnlock()
	if snapshot, exist := p.snapshots[eventID]; exist {
		if balance, exist := snapshot.snapshot[addr]; exist {
			return balance, nil
		}
		return nil, fmt.Errorf("no balance found for address %s", addr.String())
	}
	return nil, fmt.Errorf("no snapshot found for eventID %s", eventID)
}

func (p *POAPHolderProvider) BalanceAt(_ context.Context, _ common.Address, _ []byte, _ uint64) (*big.Int, error) {
	return big.NewInt(0), fmt.Errorf("not implemented")
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

func (p *POAPHolderProvider) LatestBlockNumber(_ context.Context, id []byte) (uint64, error) {
	// parse eventID from id
	eventID := string(id)
	// get the last stored snapshot
	p.snapshotsMtx.RLock()
	defer p.snapshotsMtx.RUnlock()
	if snapshot, exist := p.snapshots[eventID]; exist {
		return snapshot.from, nil
	}
	return 0, fmt.Errorf("no snapshot found for eventID %s", eventID)
}

// CreationBlock method is not implemented in the POAP external provider. By
// default it returns 0 and nil error.
func (p *POAPHolderProvider) CreationBlock(_ context.Context, _ []byte) (uint64, error) {
	return 0, nil
}

// IconURI returns the icon uri for the given id. It makes a request to the POAP
// API endpoint to get the event info for the eventID provided and returns the
// icon uri.
func (p *POAPHolderProvider) IconURI(id []byte) (string, error) {
	info, err := p.getEventInfo(string(id))
	if err != nil {
		return "", err
	}
	return info.IconURI, nil
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
	strURL, err := url.JoinPath(p.apiEndpoint, fmt.Sprintf(POAP_URI, eventID))
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
	internalCtx, cancel := context.WithTimeout(p.ctx, defaultRequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(internalCtx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", p.accessToken)
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
	strURL, err := url.JoinPath(p.apiEndpoint, fmt.Sprintf(EVENT_URI, eventID))
	if err != nil {
		return nil, err
	}
	endpoint, err := url.Parse(strURL)
	if err != nil {
		return nil, err
	}
	// create request and add headers
	internalCtx, cancel := context.WithTimeout(p.ctx, defaultRequestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(internalCtx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", p.accessToken)
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

// CensusKeys method returns the holders and balances provided transformed. The
// POAP provider does not need to transform the holders and balances, so it
// returns the data as is.
func (p *POAPHolderProvider) CensusKeys(data map[common.Address]*big.Int) (map[common.Address]*big.Int, error) {
	return data, nil
}
