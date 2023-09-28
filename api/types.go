package api

import "go.vocdoni.io/dvote/types"

type SupportedChain struct {
	ChainID   uint64 `json:"chainID"`
	ShortName string `json:"shortName"`
	Name      string `json:"name"`
}

type APIInfo struct {
	SupportedChains []SupportedChain `json:"supportedChains"`
}

type CreateTokenRequest struct {
	ID      string `json:"ID"`
	Type    string `json:"type"`
	Tags    string `json:"tags"`
	ChainID uint64 `json:"chainID"`
}

type GetTokenStatusResponse struct {
	AtBlock  uint64 `json:"atBlock"`
	Synced   bool   `json:"synced"`
	Progress int    `json:"progress"`
}

type GetTokenResponse struct {
	ID              string                  `json:"ID"`
	Type            string                  `json:"type"`
	Decimals        uint64                  `json:"decimals"`
	StartBlock      uint64                  `json:"startBlock"`
	Symbol          string                  `json:"symbol"`
	TotalSupply     string                  `json:"totalSupply"`
	Name            string                  `json:"name"`
	Status          *GetTokenStatusResponse `json:"status"`
	Size            uint64                  `json:"size"`
	DefaultStrategy uint64                  `json:"defaultStrategy,omitempty"`
	Tags            string                  `json:"tags,omitempty"`
	ChainID         uint64                  `json:"chainID"`
}

type GetTokensItem struct {
	ID         string `json:"ID"`
	Type       string `json:"type"`
	StartBlock int64  `json:"startBlock"`
	Name       string `json:"name"`
	Symbol     string `json:"symbol"`
	Tags       string `json:"tags,omitempty"`
	ChainID    uint64 `json:"chainID"`
}

type GetTokensResponse struct {
	Tokens []GetTokensItem `json:"tokens"`
}

type TokenTypesResponse struct {
	SupportedTypes []string `json:"supportedTypes"`
}

type TokenHoldersResponse struct {
	Holders map[string]string `json:"holders"`
}

type CreateCensusRequest struct {
	StrategyID  uint64 `json:"strategyID"`
	BlockNumber uint64 `json:"blockNumber"`
	Anonymous   bool   `json:"anonymous"`
}

type CreateCensusResponse struct {
	QueueID string `json:"queueID"`
}

type GetCensusResponse struct {
	CensusID   uint64         `json:"ID"`
	StrategyID uint64         `json:"strategyID"`
	MerkleRoot types.HexBytes `json:"merkleRoot"`
	URI        string         `json:"uri"`
	Size       uint64         `json:"size"`
	Weight     string         `json:"weight"`
	Anonymous  bool           `json:"anonymous"`
}

type GetCensusesResponse struct {
	Censuses []*GetCensusResponse `json:"censuses"`
}

type CensusQueueResponse struct {
	Done   bool               `json:"done"`
	Error  error              `json:"error"`
	Census *GetCensusResponse `json:"census"`
}

type StrategyToken struct {
	ID         string `json:"ID"`
	ChainID    uint64 `json:"chainID"`
	MinBalance string `json:"minBalance"`
}

type CreateStrategyRequest struct {
	Alias     string                    `json:"alias"`
	Predicate string                    `json:"predicate"`
	Tokens    map[string]*StrategyToken `json:"tokens"`
}

type GetStrategyResponse struct {
	ID        uint64                    `json:"ID"`
	Alias     string                    `json:"alias"`
	Predicate string                    `json:"predicate"`
	Tokens    map[string]*StrategyToken `json:"tokens"`
}

type GetStrategiesResponse struct {
	Strategies []*GetStrategyResponse `json:"strategies"`
}
