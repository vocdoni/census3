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
	"unsafe"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/scanner/providers"
	"go.vocdoni.io/dvote/log"
)

const (
	dateLayout = "2006-01-02T15:04:05.999Z"
	hexAddress = "0x000000000000000000000000000000000000006C"
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

func (g *GitcoinPassport) Init(iconf any) error {
	conf, ok := iconf.(GitcoinPassportConf)
	if !ok {
		return fmt.Errorf("invalid config type")
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
	g.currentBalancesMtx = sync.RWMutex{}
	g.newBalances = make(map[common.Address]*big.Int)
	g.newBalancesMtx = sync.RWMutex{}
	g.lastUpdate.Store(time.Time{})
	return nil
}

func (g *GitcoinPassport) SetRef(_ any) error {
	return nil
}

func (g *GitcoinPassport) SetLastBalances(_ context.Context, _ []byte,
	balances map[common.Address]*big.Int, _ uint64,
) error {
	log.Infof("setting last balances for %d addresses", len(balances))
	g.currentBalancesMtx.Lock()
	defer g.currentBalancesMtx.Unlock()
	for addr, balance := range balances {
		g.currentBalances[addr] = new(big.Int).Set(balance)
	}
	return nil
}

func (g *GitcoinPassport) HoldersBalances(_ context.Context, _ []byte, _ uint64) (
	map[common.Address]*big.Int, uint64, uint64, bool, error,
) {
	lastUpdate, ok := g.lastUpdate.Load().(time.Time)
	if !ok {
		return nil, 1, 0, false, fmt.Errorf("error getting last update")
	}
	if time.Since(lastUpdate) > g.cooldown && !g.downloading.Load() {
		log.Info("downloading Gitcoin Passport balances")
		go func() {
			g.downloading.Store(true)
			defer g.downloading.Store(false)

			if err := g.updateBalances(); err != nil {
				fmt.Println(err)
				log.Warnw("Error updating Gitcoin Passport balances", "err", err)
				return
			}
		}()
	}
	lastUpdateID := uint64(lastUpdate.Unix())
	if g.updated.Load() {
		log.Info("retrieving last Gitcoin Passport balances")
		g.updated.Store(false)

		g.currentBalancesMtx.RLock()
		g.newBalancesMtx.RLock()
		defer g.currentBalancesMtx.RUnlock()
		defer g.newBalancesMtx.RUnlock()
		return providers.CalcPartialHolders(g.currentBalances, g.newBalances),
			1, lastUpdateID, true, nil
	}
	log.Infof("no changes in Gitcoin Passport balances from last %s", g.cooldown)
	return nil, 1, lastUpdateID, true, nil
}

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
	log.Infof("downloading json from %s (%d bytes)...", g.apiEndpoint, res.ContentLength)
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
			log.Infow("still downloading Gitcoin Passport balances...",
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
				}
			}
		}
	}
	balancesSize := unsafe.Sizeof(balances)
	log.Infow("Gitcoin Passport balances download finished",
		"elapsed", elapsed,
		"holders", len(balances),
		"size", balancesSize)
	// remove duplicated addresses keeping the last one
	// calculate partial balances and store them
	g.updated.Store(true)
	g.newBalancesMtx.Lock()
	defer g.newBalancesMtx.Unlock()
	g.newBalances = balances
	g.lastUpdate.Store(time.Now())
	return nil
}

func (g *GitcoinPassport) Close() error {
	g.cancel()
	return nil
}

func (g *GitcoinPassport) IsExternal() bool {
	return true
}

func (g *GitcoinPassport) IsSynced(_ []byte) bool {
	g.currentBalancesMtx.RLock()
	defer g.currentBalancesMtx.RUnlock()
	return len(g.currentBalances) > 0
}

func (g *GitcoinPassport) Address() common.Address {
	return common.HexToAddress(hexAddress)
}

func (g *GitcoinPassport) Type() uint64 {
	return providers.CONTRACT_TYPE_GITCOIN
}

func (g *GitcoinPassport) ChainID() uint64 {
	return 1
}

func (g *GitcoinPassport) Name(_ []byte) (string, error) {
	return "Gitcoin Passport Score", nil
}

func (g *GitcoinPassport) Symbol(_ []byte) (string, error) {
	return "GPS", nil
}

func (g *GitcoinPassport) Decimals(_ []byte) (uint64, error) {
	return 0, nil
}

func (g *GitcoinPassport) TotalSupply(_ []byte) (*big.Int, error) {
	return big.NewInt(0), nil
}

func (g *GitcoinPassport) BalanceOf(_ common.Address, _ []byte) (*big.Int, error) {
	return big.NewInt(0), nil
}

func (g *GitcoinPassport) BalanceAt(_ context.Context, _ common.Address, _ []byte, _ uint64) (*big.Int, error) {
	return big.NewInt(0), nil
}

func (g *GitcoinPassport) BlockTimestamp(_ context.Context, _ uint64) (string, error) {
	return fmt.Sprint(time.Now()), nil
}

func (g *GitcoinPassport) BlockRootHash(_ context.Context, _ uint64) ([]byte, error) {
	lastUpdate, ok := g.lastUpdate.Load().(time.Time)
	if !ok {
		return nil, fmt.Errorf("error getting last update")
	}
	timeHash := md5.Sum([]byte(lastUpdate.Format(time.RFC3339)))
	return timeHash[:], nil
}

func (g *GitcoinPassport) LatestBlockNumber(_ context.Context, _ []byte) (uint64, error) {
	return uint64(time.Now().Unix() / 60), nil
}

func (g *GitcoinPassport) CreationBlock(_ context.Context, _ []byte) (uint64, error) {
	return 1, nil
}

func (g *GitcoinPassport) IconURI(_ []byte) (string, error) {
	return "", nil
}
