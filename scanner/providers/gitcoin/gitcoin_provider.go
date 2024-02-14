package gitcoin

import (
	"bufio"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/scanner/providers"
	"go.vocdoni.io/dvote/log"
)

const (
	dateLayout      = "2006-01-02T15:04:05.999Z"
	hexAddress      = "0x000000000000000000000000000000000000006C"
	defaultCooldown = time.Hour * 6
)

type gitcoinScoreResult struct {
	Passport struct {
		Address string `json:"address"`
	} `json:"passport"`
	Score       string         `json:"score"`
	StampScores map[string]any `json:"stamp_scores"`
	Date        string         `json:"last_score_timestamp"`
}

type GitcoinPassport struct {
	// public endpoint to download the json
	apiEndpoint string
	cooldown    time.Duration
	// internal vars to manage the download
	ctx         context.Context
	cancel      context.CancelFunc
	downloading *atomic.Bool
	updated     *atomic.Bool
	// internal vars to manage the balances
	newBalances        map[common.Address]*big.Int
	newBalancesMtx     sync.RWMutex
	currentBalances    map[common.Address]*big.Int
	currentBalancesMtx sync.RWMutex
	lastUpdate         atomic.Value
}

type GitcoinPassportConf struct {
	APIEndpoint string
	Cooldown    time.Duration
}

// Init initializes the Gitcoin Passport provider with the given config. If the
// config is not of type GitcoinPassportConf, or the API endpoint is missing, it
// returns an error. If the cooldown is not set, it defaults to 6 hours.
func (g *GitcoinPassport) Init(iconf any) error {
	conf, ok := iconf.(GitcoinPassportConf)
	if !ok {
		return fmt.Errorf("invalid config type")
	}
	if conf.APIEndpoint == "" {
		return fmt.Errorf("missing API endpoint")
	}
	if conf.Cooldown == 0 {
		conf.Cooldown = defaultCooldown
	}
	g.apiEndpoint = conf.APIEndpoint
	g.cooldown = conf.Cooldown
	// init download variables
	g.ctx, g.cancel = context.WithCancel(context.Background())
	g.downloading = new(atomic.Bool)
	g.updated = new(atomic.Bool)
	g.downloading.Store(false)
	g.updated.Store(false)
	// init balances variables
	g.currentBalances = make(map[common.Address]*big.Int)
	g.newBalances = make(map[common.Address]*big.Int)

	return nil
}

// SetRef is not implemented for Gitcoin Passport.
func (g *GitcoinPassport) SetRef(_ any) error {
	return nil
}

// SetLastBlockNumber is not implemented for Gitcoin Passport.
func (g *GitcoinPassport) SetLastBlockNumber(_ uint64) {}

// SetLastBalances stores the balances of the last block (or other kind of
// reference). It is used to calculate the partial balances of the current
// block.
func (g *GitcoinPassport) SetLastBalances(_ context.Context, _ []byte,
	balances map[common.Address]*big.Int, _ uint64,
) error {
	g.currentBalancesMtx.Lock()
	defer g.currentBalancesMtx.Unlock()
	for addr, balance := range balances {
		g.currentBalances[addr] = new(big.Int).Set(balance)
	}
	log.Debugw("last balances stored", "balances", len(balances))
	return nil
}

// HoldersBalances returns the balances of the Gitcoin Passport holders. If the
// cooldown time has passed, it starts a new download. If the download is
// finished, it returns the partial balances of the current block and the last
// block scanned. If the download is not finished or the cooldown time has not
// passed, it returns nil balances (no changes).
func (g *GitcoinPassport) HoldersBalances(_ context.Context, _ []byte, _ uint64) (
	map[common.Address]*big.Int, uint64, uint64, bool, *big.Int, error,
) {
	// if there is no last update, set it to zero
	lastUpdate, ok := g.lastUpdate.Load().(time.Time)
	if !ok {
		lastUpdate = time.Time{}
	}
	if time.Since(lastUpdate) > g.cooldown && !g.downloading.Load() {
		log.Info("downloading Gitcoin Passport balances")
		go func() {
			g.downloading.Store(true)
			defer g.downloading.Store(false)

			if err := g.updateBalances(); err != nil {
				log.Warnw("Error updating Gitcoin Passport balances", "err", err)
				return
			}
		}()
	}
	lastUpdateID := uint64(lastUpdate.Unix())
	if g.updated.Load() {
		log.Info("retrieving last Gitcoin Passport balances")
		g.updated.Store(false)

		g.newBalancesMtx.RLock()
		defer g.newBalancesMtx.RUnlock()
		// calculate total supply
		totalSupply := big.NewInt(0)
		for _, balance := range g.newBalances {
			totalSupply.Add(totalSupply, balance)
		}
		g.currentBalancesMtx.RLock()
		defer g.currentBalancesMtx.RUnlock()
		return providers.CalcPartialHolders(g.currentBalances, g.newBalances),
			1, lastUpdateID, true, totalSupply, nil
	}
	log.Infof("no changes in Gitcoin Passport balances from last %s", g.cooldown)
	return nil, 1, lastUpdateID, true, big.NewInt(0), nil
}

// updateBalances downloads the json from the API endpoint and stores the
// balances in the newBalances variable. It also stores the last update time.
func (g *GitcoinPassport) updateBalances() error {
	// download de json from API endpoint
	req, err := http.NewRequestWithContext(g.ctx, http.MethodGet, g.apiEndpoint, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error downloading json: %w", err)
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Warn("error closing response body")
		}
	}()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error downloading json: %s", res.Status)
	}
	log.Debugw("downloading json from gitcoin endpoint",
		"endpoin", g.apiEndpoint,
		"size", res.ContentLength)
	// some vars to track progress
	bytesRead := 0
	iterations := 0
	elapsed := time.Now()
	// read the json line by line
	balances := map[common.Address]*big.Int{}
	lastBalancesUpdates := map[common.Address]time.Time{}
	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		// update progress
		bytesRead += len(scanner.Bytes())
		if iterations++; iterations%10000 == 0 {
			progress := float64(bytesRead) / float64(res.ContentLength) * 100
			log.Debugw("still downloading Gitcoin Passport balances...",
				"progress", fmt.Sprintf("%.2f", progress),
				"elapsed", time.Since(elapsed).Seconds())
		}
		// parse the line
		score := &gitcoinScoreResult{}
		if err := json.Unmarshal(scanner.Bytes(), score); err != nil {
			return fmt.Errorf("error parsing json: %w", err)
		}
		// if the score is greater than 0 store it
		if score.Score != "" && score.Score != "0E-9" {
			addr := common.HexToAddress(score.Passport.Address)
			fBalance, err := strconv.ParseFloat(score.Score, 64)
			if err != nil {
				return fmt.Errorf("error parsing score: %w", err)
			}
			if fBalance != 0 {
				date, err := time.Parse(dateLayout, score.Date)
				if err != nil {
					return fmt.Errorf("error parsing date: %w", err)
				}
				if lastUpdate, exists := lastBalancesUpdates[addr]; !exists || date.After(lastUpdate) {
					balances[addr] = big.NewInt(int64(fBalance))
					lastBalancesUpdates[addr] = date
				}
			}
		}
	}
	log.Infow("Gitcoin Passport balances download finished",
		"elapsed", elapsed,
		"holders", len(balances))
	// remove duplicated addresses keeping the last one
	// calculate partial balances and store them
	g.updated.Store(true)
	g.newBalancesMtx.Lock()
	defer g.newBalancesMtx.Unlock()
	g.newBalances = balances
	g.lastUpdate.Store(time.Now())
	return nil
}

// Close cancels the download context.
func (g *GitcoinPassport) Close() error {
	g.cancel()
	return nil
}

// IsExternal returns true because Gitcoin Passport is an external provider.
func (g *GitcoinPassport) IsExternal() bool {
	return true
}

// IsSynced returns true if the balances are not empty.
func (g *GitcoinPassport) IsSynced(_ []byte) bool {
	g.currentBalancesMtx.RLock()
	defer g.currentBalancesMtx.RUnlock()
	return len(g.currentBalances) > 0
}

// Address returns the address of the Gitcoin Passport contract.
func (g *GitcoinPassport) Address(_ []byte) common.Address {
	return common.HexToAddress(hexAddress)
}

// Type returns the type of the Gitcoin Passport contract.
func (g *GitcoinPassport) Type() uint64 {
	return providers.CONTRACT_TYPE_GITCOIN
}

// TypeName returns the type name of the Gitcoin Passport contract.
func (g *GitcoinPassport) TypeName() string {
	return providers.TokenTypeName(providers.CONTRACT_TYPE_GITCOIN)
}

// ChainID returns the chain ID of the Gitcoin Passport contract.
func (g *GitcoinPassport) ChainID() uint64 {
	return 1
}

// Name returns the name of the Gitcoin Passport contract.
func (g *GitcoinPassport) Name(_ []byte) (string, error) {
	return "Gitcoin Passport Score", nil
}

// Symbol returns the symbol of the Gitcoin Passport contract.
func (g *GitcoinPassport) Symbol(_ []byte) (string, error) {
	return "GPS", nil
}

// Decimals is not implemented for Gitcoin Passport.
func (g *GitcoinPassport) Decimals(_ []byte) (uint64, error) {
	return 0, nil
}

// TotalSupply is not implemented for Gitcoin Passport.
func (g *GitcoinPassport) TotalSupply(_ []byte) (*big.Int, error) {
	return big.NewInt(0), nil
}

// BalanceOf is not implemented for Gitcoin Passport.
func (g *GitcoinPassport) BalanceOf(_ common.Address, _ []byte) (*big.Int, error) {
	return big.NewInt(0), nil
}

// BalanceAt is not implemented for Gitcoin Passport.
func (g *GitcoinPassport) BalanceAt(_ context.Context, _ common.Address, _ []byte, _ uint64) (*big.Int, error) {
	return big.NewInt(0), nil
}

// BlockTimestamp returns the timestamp of the last update of the balances.
func (g *GitcoinPassport) BlockTimestamp(_ context.Context, _ uint64) (string, error) {
	lastUpdate, ok := g.lastUpdate.Load().(time.Time)
	if !ok {
		return "", fmt.Errorf("error getting last update")
	}
	return fmt.Sprint(lastUpdate), nil
}

// BlockNumber returns the block number of the last update of the balances.
func (g *GitcoinPassport) BlockRootHash(_ context.Context, _ uint64) ([]byte, error) {
	lastUpdate, ok := g.lastUpdate.Load().(time.Time)
	if !ok {
		return nil, fmt.Errorf("error getting last update")
	}
	timeHash := md5.Sum([]byte(lastUpdate.Format(time.RFC3339)))
	return timeHash[:], nil
}

// BlockNumber returns the block number of the last update of the balances.
func (g *GitcoinPassport) LatestBlockNumber(_ context.Context, _ []byte) (uint64, error) {
	lastUpdate, ok := g.lastUpdate.Load().(time.Time)
	if !ok {
		return 0, fmt.Errorf("error getting last update")
	}
	return uint64(lastUpdate.Unix() / 60), nil
}

// CreationBlock is not implemented for Gitcoin Passport.
func (g *GitcoinPassport) CreationBlock(_ context.Context, _ []byte) (uint64, error) {
	return 1, nil
}

// IconURI is not implemented for Gitcoin Passport.
func (g *GitcoinPassport) IconURI(_ []byte) (string, error) {
	return "", nil
}

// CensusKeys method returns the holders and balances provided transformed.
// The Gitcoin Passport provider does not need to transform the holders and
// balances, so it returns the data as is.
func (p *GitcoinPassport) CensusKeys(data map[common.Address]*big.Int) (map[common.Address]*big.Int, error) {
	return data, nil
}
