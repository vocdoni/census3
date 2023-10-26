package strategyoperators

import "math/big"

// bZero is a big.Int with value 0 to be used as a default value for nil or to
// compare with balances.
var bZero = big.NewInt(0)

// normalize returns the two balances with the same number of decimals. It also
// returns the number of decimals used to normalize these numbers. To choose the
// correct number of decimals, the function chooses the highest number of
// decimals between the two provided values.
func normalize(a, b *big.Int, aDecimals, bDecimals uint64) (*big.Int, *big.Int, uint64) {
	// prevent nil pointer exceptions by assigning zero to nil balances
	if a == nil {
		a = new(big.Int).Set(bZero)
	}
	if b == nil {
		b = new(big.Int).Set(bZero)
	}
	// if a have more decimals than b, calculate the exponent by multiplying b
	// by 10^aDecimals-bDecimals
	if aDecimals > bDecimals {
		exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(aDecimals-bDecimals)), nil)
		// return the original a, b multiplied by the exponent, and the number
		// of decimals of a
		return a, new(big.Int).Mul(b, exp), aDecimals
	}
	// else if b have more decimals than a, calculate the exponent by
	// multiplying a by 10^bDecimals-aDecimals
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(bDecimals-aDecimals)), nil)
	// return a multiplied by the exponent, the original a and the number
	// of decimals of b
	return new(big.Int).Mul(a, exp), b, bDecimals
}

// reduceNormalized returns the balance provided reducing the number of decimals
// of it by the number of decimals provided. It allows to fix the normalization
// of a balance after operations like multiplication or division.
func reduceNormalized(a *big.Int, aDecimals uint64) *big.Int {
	// prevent nil pointer exceptions by assigning zero to nil balance
	if a == nil {
		a = new(big.Int).Set(bZero)
	}
	// calculate the exponent to reduce the number of decimals by powering 10
	// by the number of decimals to reduce
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(aDecimals)), nil)
	// return the balance divided by the exponent to reduce the number of
	// decimals to the desired one
	return new(big.Int).Div(a, exp)
}

// sumBalancesCombinator returns the sum of the balances provided for each
// address in the provided map. It returns a new map with the same keys and the
// result of the sum of the balances.
func sumBalancesCombinator(balances map[string][2]*big.Int) map[string]*big.Int {
	res := make(map[string]*big.Int)
	for address, balances := range balances {
		// prevent nil pointer exceptions by assigning zero to nil balances
		a := balances[0]
		if a == nil {
			a = new(big.Int).Set(bZero)
		}
		b := balances[1]
		if b == nil {
			b = new(big.Int).Set(bZero)
		}
		// sum the balances and assign the result to the address, it does not
		// matter if any of the balances is zero
		res[address] = new(big.Int).Add(a, b)
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
	for address, balance := range balances {
		// prevent nil pointer exceptions by assigning zero to nil balances
		a, b := balance[0], balance[1]
		if a == nil {
			a = new(big.Int).Set(bZero)
		}
		if b == nil {
			b = new(big.Int).Set(bZero)
		}
		// check if any both balances are zero, if so and forceNotZero is set
		// to true, continue to the next address, else assign zero to the
		// address
		if a.Cmp(bZero) == 0 && b.Cmp(bZero) == 0 {
			if !forceNotZero {
				res[address] = new(big.Int).Set(bZero)
			}
			continue
		}
		// check if any of the balances is zero, if so and forceNotZero is true,
		// continue to the next address, else assign the other one to the
		// address
		if a.Cmp(bZero) == 0 {
			if forceNotZero {
				continue
			}
			res[address] = b
			continue
		}
		if b.Cmp(bZero) == 0 {
			if forceNotZero {
				continue
			}
			res[address] = a
			continue
		}
		// if none of the balances is zero, multiply them, and if the result is
		// zero and forceNotZero is true, continue to the next, else assign the
		// result (zero) to the address
		value := new(big.Int).Mul(a, b)
		if value.Cmp(bZero) == 0 {
			if !forceNotZero {
				res[address] = value
			}
			continue
		}
		// if the result is not zero, reduce the number of decimals of the
		// original one and assign it to the address
		res[address] = reduceNormalized(value, decimals)
	}
	return res
}

// membershipCombinator returns a map with the same keys as the provided map,
// and the value of each key is 1, discarding the value of the balances of the
// provided map.
func membershipCombinator(balances map[string][2]*big.Int) map[string]*big.Int {
	res := make(map[string]*big.Int)
	for address, balance := range balances {
		// prevent nil pointer exceptions by assigning zero to nil balances
		a, b := balance[0], balance[1]
		if a == nil {
			a = new(big.Int).Set(bZero)
		}
		if b == nil {
			b = new(big.Int).Set(bZero)
		}
		// if both balances are zero, continue to the next address
		if a.Cmp(bZero) == 0 && b.Cmp(bZero) == 0 {
			continue
		}
		// else assign 1 to the address to indicate membership
		res[address] = big.NewInt(1)
	}
	return res
}
