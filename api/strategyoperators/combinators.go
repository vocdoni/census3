package strategyoperators

import "math/big"

func normalize(a, b *big.Int, aDecimals, bDecimals uint64) (*big.Int, *big.Int, uint64) {
	if aDecimals > bDecimals {
		exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(aDecimals-bDecimals)), nil)
		return a, new(big.Int).Mul(b, exp), aDecimals
	}
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(bDecimals-aDecimals)), nil)
	return new(big.Int).Mul(a, exp), b, bDecimals
}

func reduceNormalized(a *big.Int, aDecimals uint64) *big.Int {
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(aDecimals)), nil)
	return new(big.Int).Div(a, exp)
}

func sumBalancesCombinator(balances map[string][2]*big.Int) map[string]*big.Int {
	res := make(map[string]*big.Int)
	for address, balances := range balances {
		res[address] = new(big.Int).Add(balances[0], balances[1])
	}
	return res
}

func mulBalancesCombinator(balances map[string][2]*big.Int, decimals uint64, forceNotZero bool) map[string]*big.Int {
	res := make(map[string]*big.Int)
	for address, balances := range balances {
		if balances[0].Cmp(big.NewInt(0)) == 0 {
			if forceNotZero {
				continue
			}
			res[address] = balances[1]
			continue
		}
		if balances[1].Cmp(big.NewInt(0)) == 0 {
			if forceNotZero {
				continue
			}
			res[address] = balances[0]
			continue
		}

		value := new(big.Int).Mul(balances[0], balances[1])
		if value.Cmp(big.NewInt(0)) == 0 {
			continue
		}
		res[address] = reduceNormalized(value, decimals)
	}
	return res
}

func membershipCombinator(balances map[string][2]*big.Int) map[string]*big.Int {
	res := make(map[string]*big.Int)
	for address := range balances {
		res[address] = big.NewInt(1)
	}
	return res
}
