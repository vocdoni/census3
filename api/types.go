package api

import "go.vocdoni.io/dvote/types"

type Pagination struct {
	NextCursor string `json:"nextCursor"`
	PrevCursor string `json:"prevCursor"`
	PageSize   int32  `json:"pageSize"`
}

type QueueResponse struct {
	QueueID string `json:"queueID"`
}

type SupportedChain struct {
	ChainID   uint64 `json:"chainID"`
	ShortName string `json:"shortName"`
	Name      string `json:"name"`
}

type APIInfo struct {
	SupportedChains []SupportedChain `json:"supportedChains"`
}

type TokenStatus struct {
	AtBlock  uint64 `json:"atBlock"`
	Synced   bool   `json:"synced"`
	Progress int    `json:"progress"`
}

type Token struct {
	ID              string       `json:"ID"`
	Type            string       `json:"type"`
	Decimals        uint64       `json:"decimals"`
	StartBlock      uint64       `json:"startBlock"`
	Symbol          string       `json:"symbol"`
	TotalSupply     string       `json:"totalSupply"`
	Name            string       `json:"name"`
	Status          *TokenStatus `json:"status"`
	Size            uint64       `json:"size"`
	DefaultStrategy uint64       `json:"defaultStrategy,omitempty"`
	Tags            string       `json:"tags,omitempty"`
	ChainID         uint64       `json:"chainID"`
	ChainAddress    string       `json:"chainAddress"`
	ExternalID      string       `json:"externalID,omitempty"`
	IconURI         string       `json:"iconURI,omitempty"`
}

type TokenListItem struct {
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

type TokenList struct {
	Tokens     []*TokenListItem `json:"tokens"`
	Pagination *Pagination      `json:"pagination"`
}

type TokenTypes struct {
	SupportedTypes []string `json:"supportedTypes"`
}

type TokenHolders struct {
	Holders map[string]string `json:"holders"`
}

type Census struct {
	CensusID   uint64         `json:"ID"`
	StrategyID uint64         `json:"strategyID"`
	MerkleRoot types.HexBytes `json:"merkleRoot"`
	URI        string         `json:"uri"`
	Size       uint64         `json:"size"`
	Weight     string         `json:"weight"`
	Anonymous  bool           `json:"anonymous"`
	Accuracy   float64        `json:"accuracy"`
}

type Censuses struct {
	Censuses []*Census `json:"censuses"`
}

type CensusQueue struct {
	Done   bool    `json:"done"`
	Error  error   `json:"error"`
	Census *Census `json:"census"`
}

type StrategyToken struct {
	ID           string `json:"ID"`
	ChainID      uint64 `json:"chainID"`
	MinBalance   string `json:"minBalance"`
	ChainAddress string `json:"chainAddress"`
	ExternalID   string `json:"externalID,omitempty"`
	IconURI      string `json:"iconURI,omitempty"`
}

type Strategy struct {
	ID        uint64                    `json:"ID"`
	Alias     string                    `json:"alias"`
	Predicate string                    `json:"predicate"`
	URI       string                    `json:"uri,omitempty"`
	Tokens    map[string]*StrategyToken `json:"tokens"`
}

type Strategies struct {
	Strategies []*Strategy `json:"strategies"`
	Pagination *Pagination `json:"pagination"`
}

type ImportStrategyQueueResponse struct {
	Done     bool      `json:"done"`
	Error    error     `json:"error"`
	Strategy *Strategy `json:"strategy"`
}

type StrategyEstimation struct {
	Size               uint64  `json:"size"`
	TimeToCreateCensus uint64  `json:"timeToCreateCensus"`
	Accuracy           float64 `json:"accuracy"`
}

type StrategyEstimationQueue struct {
	Done       bool                `json:"done"`
	Error      error               `json:"error"`
	Estimation *StrategyEstimation `json:"estimation"`
}

type TokenHoldersAtBlock struct {
	BlockNumber uint64            `json:"blockNumber"`
	Size        int               `json:"size"`
	Holders     map[string]string `json:"holders"`
}

type TokenHolderBalance struct {
	Balance string `json:"balance"`
}

type GetHoldersAtLastBlockResponse struct {
	Done           bool                 `json:"done"`
	Error          error                `json:"error"`
	HoldersAtBlock *TokenHoldersAtBlock `json:"holdersAtBlock"`
}

type GetStrategyHoldersResponse struct {
	Holders map[string]string `json:"holders"`
}

type DeleteTokenQueueResponse struct {
	Done  bool  `json:"done"`
	Error error `json:"error"`
}
