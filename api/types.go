package api

type CreateTokenRequest struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Tag     string `json:"tag"`
	ChainID uint64 `json:"chainID"`
}

type GetTokenStatusResponse struct {
	AtBlock  uint64 `json:"atBlock"`
	Synced   bool   `json:"synced"`
	Progress int    `json:"progress"`
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
	Size            uint64                  `json:"size"`
	DefaultStrategy uint64                  `json:"defaultStrategy,omitempty"`
	Tag             string                  `json:"tag,omitempty"`
	ChainID         uint64                  `json:"chainID"`
}

type GetTokensItem struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	StartBlock int64  `json:"startBlock"`
	Name       string `json:"name"`
	Symbol     string `json:"symbol"`
	Tag        string `json:"tag,omitempty"`
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
	StrategyID  uint64 `json:"strategyId"`
	BlockNumber uint64 `json:"blockNumber"`
	Anonymous   bool   `json:"anonymous"`
}

type CreateCensusResponse struct {
	QueueID string `json:"queueId"`
}

type GetCensusResponse struct {
	CensusID   uint64 `json:"censusId"`
	StrategyID uint64 `json:"strategyId"`
	MerkleRoot string `json:"merkleRoot"`
	URI        string `json:"uri"`
	Size       uint64 `json:"size"`
	Weight     string `json:"weight"`
	Anonymous  bool   `json:"anonymous"`
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

type CensusQueueResponse struct {
	Done   bool               `json:"done"`
	Error  error              `json:"error"`
	Census *GetCensusResponse `json:"census"`
}
