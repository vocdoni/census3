package strategyoperators

import (
	"math/big"
	"testing"
)

func Test_normalize(t *testing.T) {
	a, _ := new(big.Int).SetString("6000000000000000000", 10)
	b, _ := new(big.Int).SetString("6", 10)
	aDecimals := uint64(18)
	bDecimals := uint64(0)

	expectedA, expectedB := new(big.Int).Set(a), new(big.Int).Set(a)
	resultA, resultB, commaPlaces := normalize(a, b, aDecimals, bDecimals)
	if expectedA.Cmp(resultA) != 0 || expectedB.Cmp(resultB) != 0 {
		t.Fatal(expectedA, resultA, expectedB, resultB)
	}
	if commaPlaces != aDecimals {
		t.Fatal(commaPlaces, aDecimals)
	}

	b, _ = new(big.Int).SetString("6000", 10)
	bDecimals = uint64(3)
	expectedA, expectedB = new(big.Int).Set(a), new(big.Int).Set(a)
	resultA, resultB, commaPlaces = normalize(a, b, aDecimals, bDecimals)
	if expectedA.Cmp(resultA) != 0 || expectedB.Cmp(resultB) != 0 {
		t.Fatal(expectedA, resultA, expectedB, resultB)
	}
	if commaPlaces != aDecimals {
		t.Fatal(commaPlaces, aDecimals)
	}
}

func Test_reduceNormalized(t *testing.T) {
	a, _ := new(big.Int).SetString("6000000000000000000", 10)
	aDecimals := uint64(18)
	expected := new(big.Int).SetInt64(6)
	result := reduceNormalized(a, aDecimals)
	if expected.Cmp(result) != 0 {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
	if a.Cmp(reduceNormalized(a, 0)) != 0 {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func Test_sumBalancesCombinator(t *testing.T) {
	expected := map[string]*big.Int{
		"address1": big.NewInt(300),
		"address2": big.NewInt(700),
		"address3": big.NewInt(300),
		"address4": new(big.Int).Set(bZero),
		"address5": new(big.Int).Set(bZero),
	}
	result := sumBalancesCombinator(combinatorsTestBalances)
	for address, expectedBalance := range expected {
		if result[address].Cmp(expectedBalance) != 0 {
			t.Errorf("Expected balance for %s to be %v, but got %v", address, expectedBalance, result[address])
		}
	}
}

func Test_mulBalancesCombinator(t *testing.T) {
	expected1 := map[string]*big.Int{
		"address1": big.NewInt(200),
		"address2": big.NewInt(1200),
	}
	results1 := mulBalancesCombinator(combinatorsTestBalances, 2, true)
	for address, balance := range expected1 {
		if results1[address].Cmp(balance) != 0 {
			t.Errorf("Expected balance for %s to be %v, but got %v", address, balance, results1[address])
		}
	}
	expected2 := map[string]*big.Int{
		"address1": big.NewInt(200),
		"address2": big.NewInt(1200),
		"address3": big.NewInt(300),
		"address4": big.NewInt(0),
		"address5": big.NewInt(0),
	}
	results2 := mulBalancesCombinator(combinatorsTestBalances, 2, false)
	for address, balance := range expected2 {
		if results2[address].Cmp(balance) != 0 {
			t.Errorf("Expected balance for %s to be %v, but got %v", address, balance, results2[address])
		}
	}
}

func Test_membershipCombinator(t *testing.T) {
	expected := map[string]*big.Int{
		"address1": big.NewInt(1),
		"address2": big.NewInt(1),
		"address3": big.NewInt(1),
	}

	results := membershipCombinator(combinatorsTestBalances)
	if len(expected) != len(results) {
		t.Fatalf("Expected %d results, but got %d", len(expected), len(results))
	}

	for address, balance := range results {
		if eBalance := expected[address]; balance.Cmp(eBalance) != 0 {
			t.Errorf("Expected balance for %s to be %v, but got %v", address, eBalance, balance)
		}
	}
}
