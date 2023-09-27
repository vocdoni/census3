package strategyoperators

import "math/big"

// normalize returns the two balances with the same number of decimals. It also
// returns the number of decimals used to normalize these numbers. To choose the
// correct number of decimals, the function chooses the highest number of
// decimals between the two provided values.
func normalize(a, b *big.Int, aDecimals, bDecimals uint64) (*big.Int, *big.Int, uint64) {
	if aDecimals > bDecimals {
		exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(aDecimals-bDecimals)), nil)
		return a, new(big.Int).Mul(b, exp), aDecimals
	}
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(bDecimals-aDecimals)), nil)
	return new(big.Int).Mul(a, exp), b, bDecimals
}

// reduceNormalized returns the balance provided reducing the number of decimals
// of it by the number of decimals provided. It allows to fix the normalization
// of a balance after operations like multiplication or division.
func reduceNormalized(a *big.Int, aDecimals uint64) *big.Int {
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(aDecimals)), nil)
	return new(big.Int).Div(a, exp)
}

// sumBalancesCombinator returns the sum of the balances provided for each
// address in the provided map. It returns a new map with the same keys and the
// result of the sum of the balances.
func sumBalancesCombinator(balances map[string][2]*big.Int) map[string]*big.Int {
	res := make(map[string]*big.Int)
	for address, balances := range balances {
		res[address] = new(big.Int).Add(balances[0], balances[1])
	}
	return res
}

// mulBalancesCombinator returns the multiplication of the balances provided for
// each address in the provided map. If the forceNotZero flag is set to true,
// and any of the balances is zero, the address is not included in the result.
// Else if forceNotZero is set to false, and any of the balances is zero, the
// other one will be assigned to the address. The resulting balances are reduced
// by the number of decimals provided using the reduceNormalized function.
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

// membershipCombinator returns a map with the same keys as the provided map,
// and the value of each key is 1, discarding the value of the balances of the
// provided map.
func membershipCombinator(balances map[string][2]*big.Int) map[string]*big.Int {
	res := make(map[string]*big.Int)
	for address := range balances {
		res[address] = big.NewInt(1)
	}
	return res
}
