package strategyoperators

import (
	"log"
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
		log.Fatal(expectedA, resultA, expectedB, resultB)
	}
	if commaPlaces != aDecimals {
		log.Fatal(commaPlaces, aDecimals)
	}

	b, _ = new(big.Int).SetString("6000", 10)
	bDecimals = uint64(3)
	expectedA, expectedB = new(big.Int).Set(a), new(big.Int).Set(a)
	resultA, resultB, commaPlaces = normalize(a, b, aDecimals, bDecimals)
	if expectedA.Cmp(resultA) != 0 || expectedB.Cmp(resultB) != 0 {
		log.Fatal(expectedA, resultA, expectedB, resultB)
	}
	if commaPlaces != aDecimals {
		log.Fatal(commaPlaces, aDecimals)
	}
}
