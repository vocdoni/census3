package api

type CreateTokenRequest struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Tag  string `json:"tag"`
}

type GetTokenStatusResponse struct {
	AtBlock  uint64 `json:"atBlock"`
	Synced   bool   `json:"synced"`
	Progress uint64 `json:"progress"`
}

type GetTokenResponse struct {
	ID              string                  `json:"id"`
	Type            string                  `json:"type"`
	Decimals        uint64                  `json:"decimals"`
	StartBlock      uint64                  `json:"startBlock"`
	Symbol          string                  `json:"symbol"`
	TotalSupply     string                  `json:"totalSupply"`
	Name            string                  `json:"name"`
	Status          *GetTokenStatusResponse `json:"status"`
	Size            uint32                  `json:"size"`
	DefaultStrategy uint64                  `json:"defaultStrategy,omitempty"`
	Tag             string                  `json:"tag,omitempty"`
}

type GetTokensItem struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	StartBlock uint64 `json:"startBlock"`
	Name       string `json:"name"`
	Symbol     string `json:"symbol"`
	Tag        string `json:"tag,omitempty"`
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

type CreateCensusResquest struct {
	StrategyID  uint64 `json:"strategyId"`
	BlockNumber uint64 `json:"blockNumber"`
	Anonymous   bool   `json:"anonymous"`
}

type CreateCensusResponse struct {
	CensusID uint64 `json:"censusId"`
}

type GetCensusResponse struct {
	CensusID   uint64 `json:"censusId"`
	StrategyID uint64 `json:"strategyId"`
	MerkleRoot string `json:"merkleRoot"`
	URI        string `json:"uri"`
	Size       int32  `json:"size"`
	Weight     string `json:"weight"`
	ChainID    uint64 `json:"chainId"`
	Anonymous  bool   `json:"anonymous"`
	Published  bool   `json:"published"`
}

type GetCensusesResponse struct {
	Censuses []uint64 `json:"censuses"`
}

type GetStrategiesResponse struct {
	Strategies []uint64 `json:"strategies"`
}

type GetStrategyToken struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	MinBalance string `json:"minBalance"`
	Method     string `json:"method"`
}

type GetStrategyResponse struct {
	ID        uint64             `json:"id"`
	Tokens    []GetStrategyToken `json:"tokens"`
	Predicate string             `json:"strategy"`
}
