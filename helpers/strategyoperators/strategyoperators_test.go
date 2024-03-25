package strategyoperators

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/census3/db"
	queries "github.com/vocdoni/census3/db/sqlc"
	"go.vocdoni.io/dvote/util"
)

type mockedTokenHolder struct {
	TokenID  []byte
	HolderID []byte
	ChainID  uint64
	Balance  string
}

var mockedTokens = [][]byte{
	util.RandomBytes(20),
	util.RandomBytes(20),
	util.RandomBytes(20),
}

var mockedTokenInformation = map[string]*TokenInformation{
	"A": {
		ID:         common.BytesToAddress(mockedTokens[0]).String(),
		Decimals:   18,
		ChainID:    1,
		ExternalID: "",
	},
	"B": {
		ID:         common.BytesToAddress(mockedTokens[1]).String(),
		Decimals:   18,
		ChainID:    1,
		ExternalID: "",
	},
	"C": {
		ID:         common.BytesToAddress(mockedTokens[2]).String(),
		Decimals:   18,
		ChainID:    1,
		MinBalance: "1",
	},
}

var mockedHolders = [][]byte{
	util.RandomBytes(20),
	util.RandomBytes(20),
	util.RandomBytes(20),
}

var mockedTokenHolders = []mockedTokenHolder{
	{
		HolderID: mockedHolders[0],
		TokenID:  mockedTokens[0],
		ChainID:  1,
		Balance:  big.NewInt(2000000000000000000).String(),
	},
	{
		HolderID: mockedHolders[0],
		TokenID:  mockedTokens[1],
		ChainID:  1,
		Balance:  big.NewInt(4000000000000000000).String(),
	},
	{
		HolderID: mockedHolders[1],
		TokenID:  mockedTokens[0],
		ChainID:  1,
		Balance:  big.NewInt(6000000000000000000).String(),
	},
	{
		HolderID: mockedHolders[2],
		TokenID:  mockedTokens[1],
		ChainID:  1,
		Balance:  big.NewInt(3000000000000000000).String(),
	},
}

func mockedStrategyOperator(dataDir string) (*StrategyOperators, error) {
	database, err := db.Init(dataDir, "census3.sql")
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	tx, err := database.RW.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback()
	}()
	qtx := database.QueriesRW.WithTx(tx)
	for _, m := range mockedTokenHolders {
		// if the token holder not exists, create it
		_, err = qtx.CreateTokenHolder(ctx, queries.CreateTokenHolderParams{
			TokenID:    m.TokenID,
			HolderID:   m.HolderID,
			BlockID:    1,
			Balance:    m.Balance,
			ChainID:    m.ChainID,
			ExternalID: "",
		})
		if err != nil {
			return nil, err
		}
	}
	return InitOperators(context.Background(), database.QueriesRO, mockedTokenInformation), tx.Commit()
}

var combinatorsTestBalances = map[string][2]*big.Int{
	"address1": {big.NewInt(100), big.NewInt(200)},
	"address2": {big.NewInt(300), big.NewInt(400)},
	"address3": {big.NewInt(300), nil},
	"address4": {new(big.Int).Set(bZero), nil},
	"address5": {nil, nil},
}

func TestStrategyOperatorMap(t *testing.T) {
	c := qt.New(t)
	// init a mocked strategy operator
	dbDataDir := filepath.Join(t.TempDir(), "db")
	defer c.Assert(os.RemoveAll(dbDataDir), qt.IsNil)
	mso, err := mockedStrategyOperator(dbDataDir)
	c.Assert(err, qt.IsNil)

	t.Run("OperatorsMap", func(t *testing.T) {
		res := mso.Map()
		for _, opTag := range ValidOperatorsTags {
			exist := false
			for _, op := range res {
				if op.Tag == opTag {
					exist = true
					break
				}
			}
			c.Assert(exist, qt.IsTrue)
		}
	})
	t.Run("tokenInfoBySymbol", func(t *testing.T) {
		for symbol, info := range mockedTokenInformation {
			address, chainID, minBalance, _, err := mso.tokenInfoBySymbol(symbol)
			c.Assert(err, qt.IsNil)
			c.Assert(address.String(), qt.Equals, info.ID)
			c.Assert(chainID, qt.Equals, info.ChainID)
			fmt.Printf("%+v, %+v\n", minBalance, info.MinBalance)
			if minBalance == nil || minBalance.Cmp(bZero) == 0 {
				c.Assert(info.MinBalance == "", qt.IsTrue)
			} else {
				c.Assert(minBalance.String(), qt.Equals, info.MinBalance)
			}
		}
		mso.tokensInfo["Z"] = &TokenInformation{
			ID:         common.BytesToAddress(util.RandomBytes(20)).String(),
			Decimals:   18,
			ChainID:    1,
			MinBalance: "wrongNumber",
		}
		_, _, _, _, err := mso.tokenInfoBySymbol("Z")
		c.Assert(err, qt.IsNotNil)
		_, _, _, _, err = mso.tokenInfoBySymbol("D")
		c.Assert(err, qt.IsNotNil)
	})
	t.Run("decimalsBySymbol", func(t *testing.T) {
		for symbol, info := range mockedTokenInformation {
			decimals, ok := mso.decimalsBySymbol(symbol)
			c.Assert(ok, qt.IsTrue)
			c.Assert(decimals, qt.Equals, info.Decimals)
		}
		mso.tokensInfo["Z"] = &TokenInformation{
			ID:         common.BytesToAddress(util.RandomBytes(20)).String(),
			Decimals:   18,
			ChainID:    1,
			MinBalance: "wrongNumber",
		}
		_, ok := mso.decimalsBySymbol("D")
		c.Assert(ok, qt.IsFalse)
	})
	t.Run("holdersBySymbol", func(t *testing.T) {
		testCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		mso.tokensInfo["Y"] = &TokenInformation{
			ID:       common.BytesToAddress(util.RandomBytes(20)).String(),
			Decimals: 18,
			ChainID:  1,
		}
		mso.tokensInfo["Z"] = &TokenInformation{
			ID:         common.BytesToAddress(util.RandomBytes(20)).String(),
			Decimals:   18,
			ChainID:    1,
			MinBalance: "wrongNumber",
		}
		_, err := mso.holdersBySymbol(testCtx, "Y")
		c.Assert(err, qt.IsNotNil)
		_, err = mso.holdersBySymbol(testCtx, "Z")
		c.Assert(err, qt.IsNotNil)
	})
}
