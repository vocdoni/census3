package gitcoin

import (
	"bufio"
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/db/annotations"
	"github.com/vocdoni/census3/scanner/providers"
	"github.com/vocdoni/census3/scanner/providers/gitcoin/db"
	queries "github.com/vocdoni/census3/scanner/providers/gitcoin/db/sqlc"
	"github.com/vocdoni/census3/scanner/providers/web3"
	"go.vocdoni.io/dvote/log"
)

const (
	dateLayout      = "2006-01-02T15:04:05.999Z"
	hexAddress      = "0x000000000000000000000000000000000000006C"
	gitcoinSymbol   = "GPS"
	gitcoinName     = "Gitcoin Passport Score"
	defaultCooldown = time.Hour * 6
	// timeouts
	symbolTimeout    = time.Second * 5
	balanceOfTimeout = time.Second * 10
	saveScoreTimeout = time.Second * 10
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
	db         *db.DB
	ctx        context.Context
	cancel     context.CancelFunc
	scoresChan chan *gitcoinScoreResult
	waiter     *sync.WaitGroup
	synced     *atomic.Bool
	// internal vars to manage the balances
	currentBalances    map[common.Address]*big.Int
	currentBalancesMtx sync.RWMutex
	// lastInsert time is used to simulate the last block number
	lastInsert *atomic.Int64
}

type GitcoinPassportConf struct {
	APIEndpoint string
	Cooldown    time.Duration
	DB          *db.DB
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
	if conf.DB == nil {
		return fmt.Errorf("missing DB")
	}
	if conf.Cooldown == 0 {
		conf.Cooldown = defaultCooldown
	}
	g.apiEndpoint = conf.APIEndpoint
	g.cooldown = conf.Cooldown
	g.db = conf.DB
	// init download variables
	g.ctx, g.cancel = context.WithCancel(context.Background())
	g.scoresChan = make(chan *gitcoinScoreResult)
	g.waiter = new(sync.WaitGroup)
	g.synced = new(atomic.Bool)
	g.synced.Store(false)
	g.currentBalances = make(map[common.Address]*big.Int)
	g.lastInsert = new(atomic.Int64)
	g.lastInsert.Store(0)

	g.startScoreUpdates()
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
	g.currentBalances = make(map[common.Address]*big.Int)
	for addr, score := range balances {
		g.currentBalances[addr] = score
	}

	log.Debugw("last balances stored", "balances", len(balances))
	return nil
}

func (g *GitcoinPassport) HoldersBalances(_ context.Context, stamp []byte, _ uint64) (
	map[common.Address]*big.Int, uint64, uint64, bool, *big.Int, error,
) {
	// get the current scores from the db, handle the case when the stamp is
	// empty and when it is not to get the scores from the db
	synced := g.synced.Load()
	totalSupply := big.NewInt(0)
	currentScores := make(map[common.Address]*big.Int)
	if len(stamp) > 0 {
		dbStampScores, err := g.db.QueriesRW.GetStampScores(g.ctx, string(stamp))
		if err != nil {
			return nil, 0, 0, false, big.NewInt(0), fmt.Errorf("error getting stamp scores: %w", err)
		}
		for _, dbStampScore := range dbStampScores {
			address := common.HexToAddress(string(dbStampScore.Address))
			score, ok := new(big.Int).SetString(string(dbStampScore.Score), 10)
			if !ok {
				return nil, 0, 0, false, big.NewInt(0), fmt.Errorf("error parsing score: %w", err)
			}
			currentScores[address] = score
			totalSupply.Add(totalSupply, score)
		}
	} else {
		dbScores, err := g.db.QueriesRW.GetScores(g.ctx)
		if err != nil {
			return nil, 0, 0, false, big.NewInt(0), fmt.Errorf("error getting scores: %w", err)
		}
		for _, dbScore := range dbScores {
			address := common.HexToAddress(string(dbScore.Address))
			score, ok := new(big.Int).SetString(string(dbScore.Score), 10)
			if !ok {
				return nil, 0, 0, false, big.NewInt(0), fmt.Errorf("error parsing score: %w", err)
			}
			currentScores[address] = score
			totalSupply.Add(totalSupply, score)
		}
	}
	// calculate the difference between the current balances and the last ones
	g.currentBalancesMtx.Lock()
	defer g.currentBalancesMtx.Unlock()
	holders := providers.CalcPartialHolders(g.currentBalances, currentScores)
	// return the balances, 1 new transfer, the current time as lastBlock, true
	// as a synced and the computed totalSupply
	return holders, 1, uint64(time.Now().Unix()), synced, totalSupply, nil
}

// Close cancels the download context.
func (g *GitcoinPassport) Close() error {
	g.cancel()
	g.waiter.Wait()
	close(g.scoresChan)
	return g.db.Close()
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

// Name returns the name of the Gitcoin Passport contract. If a stamp is
// provided, check if it exists and if so, compose the name with it.
func (g *GitcoinPassport) Name(stamp []byte) (string, error) {
	if len(stamp) > 0 {
		// if stamp is provided, check that it exists
		ctx, cancel := context.WithTimeout(context.Background(), symbolTimeout)
		defer cancel()
		// if the stamp does not exists, return an error
		if exists, err := g.db.QueriesRO.ExistsStamp(ctx, string(stamp)); err != nil || !exists {
			return "", fmt.Errorf("error parsing stamp provided")
		}
		// if it exists, format token name like gitcoinName:stamp
		return fmt.Sprintf("%s %s", string(stamp), gitcoinName), nil
	}
	// if no stamp is provided, return the base gitcoin passport symbol
	return gitcoinName, nil
}

// Symbol returns the symbol of the Gitcoin Passport contract. If a stamp is
// provided, check if it exists and if so, compose the symbol with it.
func (g *GitcoinPassport) Symbol(stamp []byte) (string, error) {
	if len(stamp) > 0 {
		// if stamp is provided, check that it exists
		ctx, cancel := context.WithTimeout(context.Background(), symbolTimeout)
		defer cancel()
		// if the stamp does not exists, return an error
		if exists, err := g.db.QueriesRO.ExistsStamp(ctx, string(stamp)); err != nil || !exists {
			return "", fmt.Errorf("error parsing stamp provided")
		}
		// if it exists, format token symbol like gitcoinSymbol:stamp
		return fmt.Sprintf("%s:%s", gitcoinSymbol, string(stamp)), nil
	}
	// if no stamp is provided, return the base gitcoin passport symbol
	return gitcoinSymbol, nil
}

// Decimals is not implemented for Gitcoin Passport.
func (g *GitcoinPassport) Decimals(_ []byte) (uint64, error) {
	return 0, nil
}

// TotalSupply method returns the sum of the scores of every holder in the
// database. If a stamp is provided, the total supply is calculated with the
// sum of the holders scores for that stamp.
func (g *GitcoinPassport) TotalSupply(stamp []byte) (*big.Int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), balanceOfTimeout)
	defer cancel()

	var err error
	var totalSupplyScores []annotations.BigInt
	if len(stamp) > 0 {
		totalSupplyScores, err = g.db.QueriesRO.StampTotalSupplyScores(ctx, string(stamp))
	} else {
		totalSupplyScores, err = g.db.QueriesRO.TotalSupplyScores(ctx)
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return big.NewInt(0), nil
		}
		return nil, fmt.Errorf("error getting scores from database: %v", err)
	}

	totalSupply := big.NewInt(0)
	for _, score := range totalSupplyScores {
		bScore, ok := new(big.Int).SetString(string(score), 10)
		if !ok {
			return nil, fmt.Errorf("error parsing score from database")
		}
		totalSupply.Add(totalSupply, bScore)
	}
	return totalSupply, nil
}

// BalanceOf method returns the current score of the address provided from the
// database. If any stamp name is provided, the score returned is about it and
// not the global Gitcoin Passport score.
func (g *GitcoinPassport) BalanceOf(addr common.Address, stamp []byte) (*big.Int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), balanceOfTimeout)
	defer cancel()

	var err error
	var score annotations.BigInt
	if len(stamp) > 0 {
		score, err = g.db.QueriesRO.GetStampScoreForAddress(ctx, queries.GetStampScoreForAddressParams{
			Address: annotations.Address(addr.String()),
			Stamp:   string(stamp),
		})
	} else {
		score, err = g.db.QueriesRO.GetScore(ctx, annotations.Address(addr.String()))
	}
	if err != nil {
		return nil, fmt.Errorf("error getting balance of '%s': %v", addr.String(), err)
	}
	if balance, ok := new(big.Int).SetString(string(score), 10); ok {
		return balance, nil
	}
	return nil, fmt.Errorf("error parsing holder balance")
}

// BalanceAt is not implemented for Gitcoin Passport.
func (g *GitcoinPassport) BalanceAt(_ context.Context, _ common.Address, _ []byte, _ uint64) (*big.Int, error) {
	return big.NewInt(0), nil
}

// BlockTimestamp method returns the timestamp for the provided block number, in
// this case, the block number is the time in unix seconds of the last insert in
// the gitcoin database, so the transformation is direct.
func (g *GitcoinPassport) BlockTimestamp(_ context.Context, insertTime uint64) (string, error) {
	return time.Unix(int64(insertTime), 0).Format(web3.TimeLayout), nil
}

// BlockRootHash method returns the block root hash for the provided block
// number, in this case, the unix seconds when the last score was inserted. The
// block root hash is emulated hashing the string representation of the block
// number (or last insert time) with sha256.
func (g *GitcoinPassport) BlockRootHash(_ context.Context, insertTime uint64) ([]byte, error) {
	hash := sha256.Sum256([]byte(fmt.Sprint(insertTime)))
	return hash[:], nil
}

// LatestBlockNumber method returns the number of the last block of the network,
// in this case, the last block number is emulated by the last time where an
// score was updated or inserted in the database.
func (g *GitcoinPassport) LatestBlockNumber(_ context.Context, _ []byte) (uint64, error) {
	return uint64(g.lastInsert.Load()), nil
}

// CreationBlock is not implemented for Gitcoin Passport.
func (g *GitcoinPassport) CreationBlock(_ context.Context, _ []byte) (uint64, error) {
	return 1, nil
}

// IconURI is not implemented for Gitcoin Passport.
func (g *GitcoinPassport) IconURI(_ []byte) (string, error) {
	return "", nil
}

func (g *GitcoinPassport) startScoreUpdates() {
	log.Debug("starting Gitcoin Passport score updates")
	g.waiter.Add(1)
	go func() {
		defer g.waiter.Done()
		ticker := time.NewTicker(g.cooldown)
		for {
			select {
			case <-g.ctx.Done():
				return
			default:
				if err := g.updateScores(); err != nil {
					log.Warnw("error updating Gitcoin Passport scores", "err", err)
				}
				<-ticker.C
			}
		}
	}()
	g.waiter.Add(1)
	go func() {
		defer g.waiter.Done()
		for score := range g.scoresChan {
			if err := g.saveScore(score); err != nil && !errors.Is(err, context.Canceled) {
				log.Warnw("error saving score", "err", err)
			}
		}
	}()
}

func (g *GitcoinPassport) updateScores() error {
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
	validBalances := 0
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
			if strings.Contains(err.Error(), "unexpected end of JSON input") {
				return fmt.Errorf("%v: if the process has been stopped manually, ignore this error", err)
			}
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
					g.scoresChan <- score
					validBalances++
					lastBalancesUpdates[addr] = date
				}
			}
		}
	}
	g.synced.Store(true)
	log.Infow("Gitcoin Passport balances download finished",
		"elapsed", elapsed,
		"holders", validBalances)
	return nil
}

func (g *GitcoinPassport) saveScore(score *gitcoinScoreResult) error {
	internalCtx, cancel := context.WithTimeout(g.ctx, saveScoreTimeout)
	defer cancel()
	// create a db tx to store the score
	tx, err := g.db.RW.BeginTx(internalCtx, nil)
	if err != nil {
		return fmt.Errorf("error creating tx: %w", err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(sql.ErrTxDone, err) {
			log.Warnw("error rolling back tx", "err", err)
		}
	}()
	qtx := g.db.QueriesRW.WithTx(tx)
	// parse address and balance
	address := common.HexToAddress(score.Passport.Address)
	fBalance, err := strconv.ParseFloat(score.Score, 64)
	if err != nil {
		return fmt.Errorf("error parsing score: %w", err)
	}
	dbAddress := annotations.Address(address.String())
	balance := big.NewInt(int64(fBalance))
	// get the current score, if it does not exist create it and its stamps
	currentScore, err := qtx.GetScore(internalCtx, dbAddress)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("error getting score: %w", err)
		}
		// create the score and the stamps
		if _, err := qtx.NewScore(internalCtx, queries.NewScoreParams{
			Address: dbAddress,
			Score:   annotations.BigInt(balance.String()),
		}); err != nil {
			return fmt.Errorf("error creating score: %w", err)
		}
		for name, stampScore := range score.StampScores {
			fStampScore, err := parseFloat(stampScore)
			if err != nil {
				log.Warnw("error parsing stamp score",
					"error", err.Error(),
					"score", stampScore,
					"address", address.String(),
					"name", name)
				continue
			}
			if stampBalance := big.NewInt(int64(fStampScore)); stampBalance.Cmp(big.NewInt(0)) == 1 {
				if _, err := qtx.NewStampScore(internalCtx, queries.NewStampScoreParams{
					Address: dbAddress,
					Name:    name,
					Score:   annotations.BigInt(stampBalance.String()),
				}); err != nil {
					return fmt.Errorf("error creating stamp: %w", err)
				}
			}
		}
		g.lastInsert.Store(time.Now().Unix())
		return tx.Commit()
	}
	// if the score exists and its score is different, update it
	if string(currentScore) != balance.String() {
		if _, err := qtx.UpdateScore(internalCtx, queries.UpdateScoreParams{
			Address: dbAddress,
			Score:   annotations.BigInt(balance.String()),
		}); err != nil {
			return fmt.Errorf("error updating score: %w", err)
		}
	}
	// remove current stamps before adding the current ones
	if _, err := qtx.DeleteStampForAddress(internalCtx, dbAddress); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("error deleting stamps: %w", err)
	}
	// add the current stamps
	for name, stampScore := range score.StampScores {
		fStampScore, err := parseFloat(stampScore)
		if err != nil {
			log.Warnw("error parsing stamp score",
				"error", err.Error(),
				"score", stampScore,
				"address", address.String(),
				"name", name)
			continue
		}
		if stampBalance := big.NewInt(int64(fStampScore)); stampBalance.Cmp(big.NewInt(0)) == 1 {
			if _, err := qtx.NewStampScore(internalCtx, queries.NewStampScoreParams{
				Address: dbAddress,
				Name:    name,
				Score:   annotations.BigInt(fmt.Sprint(int64(fStampScore))),
			}); err != nil {
				return fmt.Errorf("error creating stamp: %w", err)
			}
		}
	}
	g.lastInsert.Store(time.Now().Unix())
	return tx.Commit()
}

func parseFloat(input any) (float64, error) {
	switch value := input.(type) {
	case string:
		return strconv.ParseFloat(value, 64)
	case float64:
		return value, nil
	default:
		return 0, fmt.Errorf("invalid type")
	}
}
