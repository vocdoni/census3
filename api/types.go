package api

import "math/big"

type CreateTokenRequest struct {
	Address    string `json:"address"`
	Type       string `type:"type"`
	StartBlock uint64 `type:"startBlock"`
}

type GetTokenResponse struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Decimals   uint64 `json:"decimals"`
	StartBlock uint64 `json:"startBlock"`
	Name       string `json:"name"`
}

type GetTokensResponse struct {
	Tokens []GetTokenResponse `json:"tokens"`
}

type TokenTypesResponse struct {
	SupportedTypes []string `json:"supportedTypes"`
}

type TokenHoldersResponse struct {
	Holders []string `json:"holders"`
}

type CreateCensusResquest struct {
	StrategyID  uint64 `json:"strategyId"`
	BlockNumber uint64 `json:"blockNumber"`
}

type CreateCensusResponse struct {
	CensusID uint64 `json:"censusId"`
}

type GetCensusResponse struct {
	CensusID   uint64 `json:"censusId"`
	StrategyID uint64 `json:"strategyId"`
	MerkleRoot string `json:"merkleRoot"`
	URI        string `json:"uri"`
}

type GetCensusesResponse struct {
	Censuses []uint64 `json:"censuses"`
}

type GetStrategiesResponse struct {
	Strategies []uint64 `json:"strategies"`
}

type GetStrategyToken struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	MinBalance *big.Int `json:"minBalance"`
	Method     string   `json:"method"`
}

type GetStrategyResponse struct {
	ID        uint64             `json:"id"`
	Tokens    []GetStrategyToken `json:"tokens"`
	Predicate string             `json:"strategy"`
}
