package api

// CreateTokenRequest struct defines the expected request by create token handler
type CreateTokenRequest struct {
	Address    string `json:"address"`
	Type       string `type:"type"`
	StartBlock uint64 `type:"startBlock"`
}

// TokenHoldersResponse struct defines the response of token holders handler
type TokenHoldersResponse struct {
	Holders []string `json:"holders"`
}

type CreateCensusResquest struct {
	StrategyID  uint64 `json:"strategyId"`
	BlockNumber uint64 `json:"blockNumber"`
}

type CreateCensusResponse struct {
	CensusID string `json:"censusId"`
}

type GetCensusResponse struct {
	CensusID   string `json:"censusId"`
	StrategyID string `json:"strategyId"`
	MerkleRoot string `json:"merkleRoot"`
	URI        string `json:"uri"`
}
