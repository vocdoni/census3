package lexer

import (
	"fmt"
	"slices"
	"sort"
)

var data = map[string][]string{
	"MATIC (staked)": {
		"0x63ca58fdcfa0bab3e743", // result 1
		"0x75df00c550e4d5867269", // result 2
		"0xbb0574ee8d15b5869fef", // result 3
		"0xb9e3e451eebbbdcaed5a",
		"0x9be089b89bf7bbddd5ea",
	},
	"USDC": {
		"0x63ca58fdcfa0bab3e743", // result 1
		"0xcd82100c3bc9411ed352",
		"0xbb0574ee8d15b5869fef", // result 3
		"0x6a78167e9a02d3693a6e",
		"0xbb54d600f7dde07440ae",
	},
	"BTC": {
		"0x995129124181c3d67974",
		"0x75df00c550e4d5867269", // result 2
		"0xcf62ab575da9daece6c2",
		"0xd7427cf2e3e766e0e9de",
		"0x774d1ca6845e2f30ba2d",
	},
	"ETH": {
		"0x2479badef3c5f9d10bc7",
		"0x1650fd8d9aa883fef994",
		"0x75df00c550e4d5867269", // result 2
		"0x61ff45597f7ccb7cad1b",
		"0xaa5ba44ce301c784926d",
	},
	"wANT": {
		"0xbb0574ee8d15b5869fef", // result 3
		"0xc18d1a2c1687b6711832",
		"0x21dc25c4650fffbbff76",
		"0xb014c34d7e57a0da5284",
		"0xa88f67b54395697c6810",
	},
	"ANT": {
		"0x63ca58fdcfa0bab3e743", // result 1
		"0xbee2a6d71759d55238bf",
		"0x4dbf5af682b609c7ae41",
		"0x8ddf4e74cdaef6b44ffd",
		"0x75df00c550e4d5867269", // result 2
	},
}

func AND(iter *Iteration[[]string]) ([]string, error) {
	tagA, dataA := iter.A()
	if dataA == nil {
		dataA = data[tagA]
	}
	tagB, dataB := iter.B()
	if dataB == nil {
		dataB = data[tagB]
	}

	res := []string{}
	for _, a := range dataA {
		if slices.Contains(dataB, a) {
			res = append(res, a)
		}
	}
	return res, nil
}

func OR(iter *Iteration[[]string]) ([]string, error) {
	tagA, dataA := iter.A()
	if dataA == nil {
		dataA = data[tagA]
	}
	tagB, dataB := iter.B()
	if dataB == nil {
		dataB = data[tagB]
	}

	res := append([]string{}, dataB...)
	for _, a := range dataA {
		exist := false
		for _, b := range dataB {
			if a == b {
				exist = true
				break
			}
		}
		if !exist {
			res = append(res, a)
		}
	}
	return res, nil
}

func XOR(iter *Iteration[[]string]) ([]string, error) {
	tagA, dataA := iter.A()
	if dataA == nil {
		dataA = data[tagA]
	}
	tagB, dataB := iter.B()
	if dataB == nil {
		dataB = data[tagB]
	}

	res := []string{}
	for _, a := range dataA {
		exist := false
		for _, b := range dataB {
			if a == b {
				exist = true
				break
			}
		}
		if !exist {
			res = append(res, a)
		}
	}
	for _, b := range dataB {
		exist := false
		for _, a := range dataA {
			if b == a {
				exist = true
				break
			}
		}
		if !exist {
			res = append(res, b)
		}
	}
	return res, nil
}

func Example() {
	predicate := "MATIC\\ \\(staked\\) AND ((USDC OR (BTC AND ETH)) AND (wANT XOR ANT))"
	operators := []*Operator[[]string]{
		{Tag: "AND", Fn: AND},
		{Tag: "OR", Fn: OR},
		{Tag: "XOR", Fn: XOR},
	}
	operatorsTags := []string{"AND", "OR", "XOR"}
	// create lexer and parse predicate
	rootToken, err := NewLexer(operatorsTags).Parse(predicate)
	if err != nil {
		panic(err)
	}
	// create evaluator and eval with operators
	res, err := NewEval(operators).EvalToken(rootToken)
	if err != nil {
		panic(err)
	}
	sort.Strings(res)
	fmt.Println(res)
	// Output: [0x63ca58fdcfa0bab3e743 0x75df00c550e4d5867269 0xbb0574ee8d15b5869fef]
}
