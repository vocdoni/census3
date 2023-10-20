package strategyoperators

import (
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/census3/lexer"
)

func Test_basicOperators(t *testing.T) {
	c := qt.New(t)
	// init a mocked strategy operator
	dbDataDir := filepath.Join(t.TempDir(), "db")
	defer c.Assert(os.RemoveAll(dbDataDir), qt.IsNil)
	mso, err := mockedStrategyOperator(dbDataDir)
	c.Assert(err, qt.IsNil)

	t.Run("andOperator", func(t *testing.T) {
		// check basic AND operator with both holder list from database
		iter := lexer.MockIteration[*StrategyIteration]("A", "B", nil, nil)
		results, decimals, err := mso.andOperator(iter)
		c.Assert(err, qt.IsNil)
		c.Assert(decimals, qt.Equals, uint64(18))
		expectedHolder := common.BytesToAddress(mockedHolders[0])
		balances, ok := results[expectedHolder.String()]
		c.Assert(ok, qt.Equals, true)
		c.Assert(balances[0].String(), qt.Equals, "2000000000000000000")
		c.Assert(balances[1].String(), qt.Equals, "4000000000000000000")
		// declare partial results
		partialResults := &StrategyIteration{
			decimals: 18,
			Data: map[string]*big.Int{
				common.BytesToAddress(mockedHolders[0]).String(): big.NewInt(4000000000000000000),
				common.BytesToAddress(mockedHolders[1]).String(): big.NewInt(2000000000000000000),
				common.BytesToAddress(mockedHolders[2]).String(): big.NewInt(3000000000000000000),
			},
		}
		// check with partial results from previous iteration as first part of the
		// AND operator
		iter = lexer.MockIteration[*StrategyIteration]("", "B", partialResults, nil)
		results, decimals, err = mso.andOperator(iter)
		c.Assert(err, qt.IsNil)
		c.Assert(decimals, qt.Equals, uint64(18))
		expectedHolders := map[string][]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): {
				big.NewInt(4000000000000000000),
				big.NewInt(4000000000000000000),
			},
			common.BytesToAddress(mockedHolders[2]).String(): {
				big.NewInt(3000000000000000000),
				big.NewInt(3000000000000000000),
			},
		}
		c.Assert(len(results), qt.Equals, len(expectedHolders))
		for holder, balances := range results {
			expectedBalances, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balances[0].String(), qt.Equals, expectedBalances[0].String())
			c.Assert(balances[1].String(), qt.Equals, expectedBalances[1].String())
		}
		// check with partial results from previous iteration as second part of the
		// AND operator
		iter = lexer.MockIteration[*StrategyIteration]("A", "", nil, partialResults)
		results, decimals, err = mso.andOperator(iter)
		c.Assert(err, qt.IsNil)
		c.Assert(decimals, qt.Equals, uint64(18))
		expectedHolders = map[string][]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): {
				big.NewInt(2000000000000000000),
				big.NewInt(4000000000000000000),
			},
			common.BytesToAddress(mockedHolders[1]).String(): {
				big.NewInt(6000000000000000000),
				big.NewInt(2000000000000000000),
			},
		}
		c.Assert(len(results), qt.Equals, len(expectedHolders))
		for holder, balances := range results {
			expectedBalances, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balances[0].String(), qt.Equals, expectedBalances[0].String())
			c.Assert(balances[1].String(), qt.Equals, expectedBalances[1].String())
		}
	})
	t.Run("orOperator", func(t *testing.T) {
		// check basic OR operator with both holder list from database
		iter := lexer.MockIteration[*StrategyIteration]("A", "B", nil, nil)
		results, decimals, err := mso.orOperator(iter)
		c.Assert(err, qt.IsNil)
		c.Assert(decimals, qt.Equals, uint64(18))
		expectedHolders := map[string][2]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): {
				big.NewInt(2000000000000000000),
				big.NewInt(4000000000000000000),
			},
			common.BytesToAddress(mockedHolders[1]).String(): {
				big.NewInt(6000000000000000000),
				new(big.Int).Set(bZero),
			},
			common.BytesToAddress(mockedHolders[2]).String(): {
				new(big.Int).Set(bZero),
				big.NewInt(3000000000000000000),
			},
		}
		c.Assert(len(results), qt.Equals, len(expectedHolders))
		for holder, balances := range results {
			expectedBalances, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balances[0].String(), qt.Equals, expectedBalances[0].String())
			c.Assert(balances[1].String(), qt.Equals, expectedBalances[1].String())
		}

		// declare partial results
		partialResults := &StrategyIteration{
			decimals: 18,
			Data: map[string]*big.Int{
				common.BytesToAddress(mockedHolders[0]).String(): big.NewInt(4000000000000000000),
				common.BytesToAddress(mockedHolders[1]).String(): big.NewInt(2000000000000000000),
				common.BytesToAddress(mockedHolders[2]).String(): big.NewInt(3000000000000000000),
			},
		}
		// check with partial results from previous iteration as first part of the
		// OR operator
		iter = lexer.MockIteration[*StrategyIteration]("", "B", partialResults, nil)
		results, decimals, err = mso.orOperator(iter)
		c.Assert(err, qt.IsNil)
		c.Assert(decimals, qt.Equals, uint64(18))
		expectedHolders = map[string][2]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): {
				big.NewInt(4000000000000000000),
				big.NewInt(4000000000000000000),
			},
			common.BytesToAddress(mockedHolders[1]).String(): {
				big.NewInt(2000000000000000000),
				new(big.Int).Set(bZero),
			},
			common.BytesToAddress(mockedHolders[2]).String(): {
				big.NewInt(3000000000000000000),
				big.NewInt(3000000000000000000),
			},
		}
		c.Assert(len(results), qt.Equals, len(expectedHolders))
		for holder, balances := range results {
			expectedBalances, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balances[0].String(), qt.Equals, expectedBalances[0].String())
			c.Assert(balances[1].String(), qt.Equals, expectedBalances[1].String())
		}
		// check with partial results from previous iteration as second part of the
		// OR operator
		iter = lexer.MockIteration[*StrategyIteration]("A", "", nil, partialResults)
		results, decimals, err = mso.orOperator(iter)
		c.Assert(err, qt.IsNil)
		c.Assert(decimals, qt.Equals, uint64(18))
		expectedHolders = map[string][2]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): {
				big.NewInt(2000000000000000000),
				big.NewInt(4000000000000000000),
			},
			common.BytesToAddress(mockedHolders[1]).String(): {
				big.NewInt(6000000000000000000),
				big.NewInt(2000000000000000000),
			},
			common.BytesToAddress(mockedHolders[2]).String(): {
				new(big.Int).Set(bZero),
				big.NewInt(3000000000000000000),
			},
		}
		c.Assert(len(results), qt.Equals, len(expectedHolders))
		for holder, balances := range results {
			expectedBalances, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balances[0].String(), qt.Equals, expectedBalances[0].String())
			c.Assert(balances[1].String(), qt.Equals, expectedBalances[1].String())
		}
	})
}

func TestOperators(t *testing.T) {
	c := qt.New(t)
	// init a mocked strategy operator
	dbDataDir := filepath.Join(t.TempDir(), "db")
	defer c.Assert(os.RemoveAll(dbDataDir), qt.IsNil)
	mso, err := mockedStrategyOperator(dbDataDir)
	c.Assert(err, qt.IsNil)

	t.Run("AND", func(t *testing.T) {
		// init AND operator
		iter := lexer.MockIteration[*StrategyIteration]("A", "B", nil, nil)
		op := mso.AND()
		// check AND operator with both holder list from database
		results, err := op(iter)
		c.Assert(err, qt.IsNil)
		expectedHolders := map[string]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): big.NewInt(1),
		}
		c.Assert(len(results.Data), qt.Equals, len(expectedHolders))
		for holder, balance := range results.Data {
			expectedBalance, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
		}
		// declare partial results
		partialResults := &StrategyIteration{
			decimals: 18,
			Data: map[string]*big.Int{
				common.BytesToAddress(mockedHolders[0]).String(): big.NewInt(4000000000000000000),
				common.BytesToAddress(mockedHolders[1]).String(): big.NewInt(2000000000000000000),
				common.BytesToAddress(mockedHolders[2]).String(): big.NewInt(3000000000000000000),
			},
		}
		// check with partial results from previous iteration as first part of the
		// AND operator
		iter = lexer.MockIteration[*StrategyIteration]("", "B", partialResults, nil)
		results, err = op(iter)
		c.Assert(err, qt.IsNil)
		expectedHolders = map[string]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): big.NewInt(1),
			common.BytesToAddress(mockedHolders[2]).String(): big.NewInt(1),
		}
		c.Assert(len(results.Data), qt.Equals, len(expectedHolders))
		for holder, balance := range results.Data {
			expectedBalance, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
		}
		// check with partial results from previous iteration as second part of the
		// AND operator
		iter = lexer.MockIteration[*StrategyIteration]("A", "", nil, partialResults)
		results, err = op(iter)
		c.Assert(err, qt.IsNil)
		expectedHolders = map[string]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): big.NewInt(1),
			common.BytesToAddress(mockedHolders[1]).String(): big.NewInt(1),
		}
		c.Assert(len(results.Data), qt.Equals, len(expectedHolders))
		for holder, balance := range results.Data {
			expectedBalance, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
		}
	})
	t.Run("AND_SUM", func(t *testing.T) {
		// init AND operator
		iter := lexer.MockIteration[*StrategyIteration]("A", "B", nil, nil)
		op := mso.AND_SUM()
		// check AND operator with both holder list from database
		results, err := op(iter)
		c.Assert(err, qt.IsNil)
		expectedHolders := map[string]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): big.NewInt(6000000000000000000),
		}
		c.Assert(len(results.Data), qt.Equals, len(expectedHolders))
		for holder, balance := range results.Data {
			expectedBalance, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
		}
		// declare partial results
		partialResults := &StrategyIteration{
			decimals: 18,
			Data: map[string]*big.Int{
				common.BytesToAddress(mockedHolders[0]).String(): big.NewInt(4000000000000000000),
				common.BytesToAddress(mockedHolders[1]).String(): big.NewInt(2000000000000000000),
				common.BytesToAddress(mockedHolders[2]).String(): big.NewInt(3000000000000000000),
			},
		}
		// check with partial results from previous iteration as first part of the
		// AND operator
		iter = lexer.MockIteration[*StrategyIteration]("", "B", partialResults, nil)
		results, err = op(iter)
		c.Assert(err, qt.IsNil)
		expectedHolders = map[string]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): big.NewInt(8000000000000000000),
			common.BytesToAddress(mockedHolders[2]).String(): big.NewInt(6000000000000000000),
		}
		c.Assert(len(results.Data), qt.Equals, len(expectedHolders))
		for holder, balance := range results.Data {
			expectedBalance, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
		}
		// check with partial results from previous iteration as second part of the
		// AND operator
		iter = lexer.MockIteration[*StrategyIteration]("A", "", nil, partialResults)
		results, err = op(iter)
		c.Assert(err, qt.IsNil)
		expectedHolders = map[string]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): big.NewInt(6000000000000000000),
			common.BytesToAddress(mockedHolders[1]).String(): big.NewInt(8000000000000000000),
		}
		c.Assert(len(results.Data), qt.Equals, len(expectedHolders))
		for holder, balance := range results.Data {
			expectedBalance, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
		}
	})
	t.Run("AND_MUL", func(t *testing.T) {
		// init AND operator
		iter := lexer.MockIteration[*StrategyIteration]("A", "B", nil, nil)
		op := mso.AND_MUL()
		// check AND operator with both holder list from database
		results, err := op(iter)
		c.Assert(err, qt.IsNil)
		expectedHolders := map[string]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): big.NewInt(8000000000000000000),
		}
		c.Assert(len(results.Data), qt.Equals, len(expectedHolders))
		for holder, balance := range results.Data {
			expectedBalance, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
		}
		// declare partial results
		partialResults := &StrategyIteration{
			decimals: 18,
			Data: map[string]*big.Int{
				common.BytesToAddress(mockedHolders[0]).String(): big.NewInt(4000000000000000000),
				common.BytesToAddress(mockedHolders[1]).String(): big.NewInt(2000000000000000000),
				common.BytesToAddress(mockedHolders[2]).String(): big.NewInt(3000000000000000000),
			},
		}
		// check with partial results from previous iteration as first part of the
		// AND operator
		iter = lexer.MockIteration[*StrategyIteration]("", "B", partialResults, nil)
		results, err = op(iter)
		c.Assert(err, qt.IsNil)
		expectedHolders = map[string]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): new(big.Int).SetUint64(16000000000000000000),
			common.BytesToAddress(mockedHolders[2]).String(): big.NewInt(9000000000000000000),
		}
		c.Assert(len(results.Data), qt.Equals, len(expectedHolders))
		for holder, balance := range results.Data {
			expectedBalance, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
		}
		// check with partial results from previous iteration as second part of the
		// AND operator
		iter = lexer.MockIteration[*StrategyIteration]("A", "", nil, partialResults)
		results, err = op(iter)
		c.Assert(err, qt.IsNil)
		expectedHolders = map[string]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): new(big.Int).SetUint64(8000000000000000000),
			common.BytesToAddress(mockedHolders[1]).String(): new(big.Int).SetUint64(12000000000000000000),
		}
		c.Assert(len(results.Data), qt.Equals, len(expectedHolders))
		for holder, balance := range results.Data {
			expectedBalance, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
		}
	})
	t.Run("OR", func(t *testing.T) {
		// init AND operator
		iter := lexer.MockIteration[*StrategyIteration]("A", "B", nil, nil)
		op := mso.OR()
		// check AND operator with both holder list from database
		results, err := op(iter)
		c.Assert(err, qt.IsNil)
		expectedHolders := map[string]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): big.NewInt(1),
			common.BytesToAddress(mockedHolders[1]).String(): big.NewInt(1),
			common.BytesToAddress(mockedHolders[2]).String(): big.NewInt(1),
		}
		c.Assert(len(results.Data), qt.Equals, len(expectedHolders))
		for holder, balance := range results.Data {
			expectedBalance, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
		}
		// declare partial results
		partialResults := &StrategyIteration{
			decimals: 18,
			Data: map[string]*big.Int{
				common.BytesToAddress(mockedHolders[0]).String(): new(big.Int).SetUint64(4000000000000000000),
				common.BytesToAddress(mockedHolders[1]).String(): new(big.Int).SetUint64(2000000000000000000),
				common.BytesToAddress(mockedHolders[2]).String(): new(big.Int).SetUint64(3000000000000000000),
			},
		}
		// check with partial results from previous iteration as first part of the
		// AND operator
		iter = lexer.MockIteration[*StrategyIteration]("", "B", partialResults, nil)
		results, err = op(iter)
		c.Assert(err, qt.IsNil)
		expectedHolders = map[string]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): big.NewInt(1),
			common.BytesToAddress(mockedHolders[1]).String(): big.NewInt(1),
			common.BytesToAddress(mockedHolders[2]).String(): big.NewInt(1),
		}
		c.Assert(len(results.Data), qt.Equals, len(expectedHolders))
		for holder, balance := range results.Data {
			expectedBalance, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
		}
		// check with partial results from previous iteration as second part of the
		// AND operator
		iter = lexer.MockIteration[*StrategyIteration]("A", "", nil, partialResults)
		results, err = op(iter)
		c.Assert(err, qt.IsNil)
		expectedHolders = map[string]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): big.NewInt(1),
			common.BytesToAddress(mockedHolders[1]).String(): big.NewInt(1),
			common.BytesToAddress(mockedHolders[2]).String(): big.NewInt(1),
		}
		c.Assert(len(results.Data), qt.Equals, len(expectedHolders))
		for holder, balance := range results.Data {
			expectedBalance, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
		}
	})
	t.Run("OR_SUM", func(t *testing.T) {
		// init AND operator
		iter := lexer.MockIteration[*StrategyIteration]("A", "B", nil, nil)
		op := mso.OR_SUM()
		// check AND operator with both holder list from database
		results, err := op(iter)
		c.Assert(err, qt.IsNil)
		expectedHolders := map[string]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): new(big.Int).SetUint64(6000000000000000000),
			common.BytesToAddress(mockedHolders[1]).String(): new(big.Int).SetUint64(6000000000000000000),
			common.BytesToAddress(mockedHolders[2]).String(): new(big.Int).SetUint64(3000000000000000000),
		}
		c.Assert(len(results.Data), qt.Equals, len(expectedHolders))
		for holder, balance := range results.Data {
			expectedBalance, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
		}
		// declare partial results
		partialResults := &StrategyIteration{
			decimals: 18,
			Data: map[string]*big.Int{
				common.BytesToAddress(mockedHolders[0]).String(): new(big.Int).SetUint64(4000000000000000000),
				common.BytesToAddress(mockedHolders[1]).String(): new(big.Int).SetUint64(2000000000000000000),
				common.BytesToAddress(mockedHolders[2]).String(): new(big.Int).SetUint64(3000000000000000000),
			},
		}
		// check with partial results from previous iteration as first part of the
		// AND operator
		iter = lexer.MockIteration[*StrategyIteration]("", "B", partialResults, nil)
		results, err = op(iter)
		c.Assert(err, qt.IsNil)
		expectedHolders = map[string]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): new(big.Int).SetUint64(8000000000000000000),
			common.BytesToAddress(mockedHolders[1]).String(): new(big.Int).SetUint64(2000000000000000000),
			common.BytesToAddress(mockedHolders[2]).String(): new(big.Int).SetUint64(6000000000000000000),
		}
		c.Assert(len(results.Data), qt.Equals, len(expectedHolders))
		for holder, balance := range results.Data {
			expectedBalance, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
		}
		// check with partial results from previous iteration as second part of the
		// AND operator
		iter = lexer.MockIteration[*StrategyIteration]("A", "", nil, partialResults)
		results, err = op(iter)
		c.Assert(err, qt.IsNil)
		expectedHolders = map[string]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): new(big.Int).SetUint64(6000000000000000000),
			common.BytesToAddress(mockedHolders[1]).String(): new(big.Int).SetUint64(8000000000000000000),
			common.BytesToAddress(mockedHolders[2]).String(): new(big.Int).SetUint64(3000000000000000000),
		}
		c.Assert(len(results.Data), qt.Equals, len(expectedHolders))
		for holder, balance := range results.Data {
			expectedBalance, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
		}
	})
	t.Run("OR_MUL", func(t *testing.T) {
		// init AND operator
		iter := lexer.MockIteration[*StrategyIteration]("A", "B", nil, nil)
		op := mso.OR_MUL()
		// check AND operator with both holder list from database
		results, err := op(iter)
		c.Assert(err, qt.IsNil)
		expectedHolders := map[string]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): new(big.Int).SetUint64(8000000000000000000),
			common.BytesToAddress(mockedHolders[1]).String(): new(big.Int).SetUint64(6000000000000000000),
			common.BytesToAddress(mockedHolders[2]).String(): new(big.Int).SetUint64(3000000000000000000),
		}
		c.Assert(len(results.Data), qt.Equals, len(expectedHolders))
		for holder, balance := range results.Data {
			expectedBalance, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
		}
		// declare partial results
		partialResults := &StrategyIteration{
			decimals: 18,
			Data: map[string]*big.Int{
				common.BytesToAddress(mockedHolders[0]).String(): new(big.Int).SetUint64(4000000000000000000),
				common.BytesToAddress(mockedHolders[1]).String(): new(big.Int).SetUint64(2000000000000000000),
				common.BytesToAddress(mockedHolders[2]).String(): new(big.Int).SetUint64(3000000000000000000),
			},
		}
		// check with partial results from previous iteration as first part of the
		// AND operator
		iter = lexer.MockIteration[*StrategyIteration]("", "B", partialResults, nil)
		results, err = op(iter)
		c.Assert(err, qt.IsNil)
		expectedHolders = map[string]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): new(big.Int).SetUint64(16000000000000000000),
			common.BytesToAddress(mockedHolders[1]).String(): new(big.Int).SetUint64(2000000000000000000),
			common.BytesToAddress(mockedHolders[2]).String(): new(big.Int).SetUint64(9000000000000000000),
		}
		c.Assert(len(results.Data), qt.Equals, len(expectedHolders))
		for holder, balance := range results.Data {
			expectedBalance, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
		}
		// check with partial results from previous iteration as second part of the
		// AND operator
		iter = lexer.MockIteration[*StrategyIteration]("A", "", nil, partialResults)
		results, err = op(iter)
		c.Assert(err, qt.IsNil)
		expectedHolders = map[string]*big.Int{
			common.BytesToAddress(mockedHolders[0]).String(): new(big.Int).SetUint64(8000000000000000000),
			common.BytesToAddress(mockedHolders[1]).String(): new(big.Int).SetUint64(12000000000000000000),
			common.BytesToAddress(mockedHolders[2]).String(): new(big.Int).SetUint64(3000000000000000000),
		}
		c.Assert(len(results.Data), qt.Equals, len(expectedHolders))
		for holder, balance := range results.Data {
			expectedBalance, ok := expectedHolders[holder]
			c.Assert(ok, qt.Equals, true)
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
			c.Assert(balance.String(), qt.Equals, expectedBalance.String())
		}
	})
}
