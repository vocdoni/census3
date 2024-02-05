package api

import "go.vocdoni.io/dvote/types"

type Pagination struct {
	NextCursor string `json:"nextCursor"`
	PrevCursor string `json:"prevCursor"`
	PageSize   int32  `json:"pageSize"`
}

type SupportedChain struct {
	ChainID   uint64 `json:"chainID"`
	ShortName string `json:"shortName"`
	Name      string `json:"name"`
}

type APIInfo struct {
	SupportedChains []SupportedChain `json:"supportedChains"`
}

type CreateTokenRequest struct {
	ID         string `json:"ID"`
	Type       string `json:"type"`
	Tags       string `json:"tags"`
	ChainID    uint64 `json:"chainID"`
	ExternalID string `json:"externalID"`
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
	ChainAddress    string                  `json:"chainAddress"`
	ExternalID      string                  `json:"externalID,omitempty"`
	IconURI         string                  `json:"iconURI,omitempty"`
}

type GetTokensItemResponse struct {
	ID              string `json:"ID"`
	Type            string `json:"type"`
	Decimals        uint64 `json:"decimals"`
	StartBlock      uint64 `json:"startBlock"`
	Symbol          string `json:"symbol"`
	TotalSupply     string `json:"totalSupply"`
	Name            string `json:"name"`
	Synced          bool   `json:"synced"`
	DefaultStrategy uint64 `json:"defaultStrategy,omitempty"`
	Tags            string `json:"tags,omitempty"`
	ChainID         uint64 `json:"chainID"`
	ChainAddress    string `json:"chainAddress"`
	ExternalID      string `json:"externalID,omitempty"`
	IconURI         string `json:"iconURI,omitempty"`
}

type GetTokensResponse struct {
	Tokens     []GetTokensItemResponse `json:"tokens"`
	Pagination *Pagination             `json:"pagination"`
}

type TokenTypesResponse struct {
	SupportedTypes []string `json:"supportedTypes"`
}

type TokenHoldersResponse struct {
	Holders map[string]string `json:"holders"`
}

type CreateCensusRequest struct {
	StrategyID uint64 `json:"strategyID"`
	Anonymous  bool   `json:"anonymous"`
}

type QueueResponse struct {
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
	Accuracy   float64        `json:"accuracy"`
}

type GetCensusesResponse struct {
	Censuses []*GetCensusResponse `json:"censuses"`
}

type GetStrategyToken struct {
	ID         string `json:"ID"`
	Name       string `json:"name"`
	MinBalance string `json:"minBalance"`
	Method     string `json:"method"`
}

type CensusQueueResponse struct {
	Done   bool               `json:"done"`
	Error  error              `json:"error"`
	Census *GetCensusResponse `json:"census"`
}

type StrategyToken struct {
	ID           string `json:"ID"`
	ChainID      uint64 `json:"chainID"`
	MinBalance   string `json:"minBalance"`
	ChainAddress string `json:"chainAddress"`
	ExternalID   string `json:"externalID,omitempty"`
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
	URI       string                    `json:"uri,omitempty"`
	Tokens    map[string]*StrategyToken `json:"tokens"`
}

type GetStrategiesResponse struct {
	Strategies []*GetStrategyResponse `json:"strategies"`
	Pagination *Pagination            `json:"pagination"`
}

type ImportStrategyQueueResponse struct {
	Done     bool                 `json:"done"`
	Error    error                `json:"error"`
	Strategy *GetStrategyResponse `json:"strategy"`
}

type StrategyEstimation struct {
	Size               uint64  `json:"size"`
	TimeToCreateCensus uint64  `json:"timeToCreateCensus"`
	Accuracy           float64 `json:"accuracy"`
}

type GetStrategyEstimationResponse struct {
	Done       bool                `json:"done"`
	Error      error               `json:"error"`
	Estimation *StrategyEstimation `json:"estimation"`
}

type TokenHoldersAtBlock struct {
	BlockNumber uint64            `json:"blockNumber"`
	Size        int               `json:"size"`
	Holders     map[string]string `json:"holders"`
}

type GetHoldersAtLastBlockResponse struct {
	Done           bool                 `json:"done"`
	Error          error                `json:"error"`
	HoldersAtBlock *TokenHoldersAtBlock `json:"holdersAtBlock"`
}
