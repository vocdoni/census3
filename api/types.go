package api

// CreateTokenRequest struct defines the expected request by create token handler
type CreateTokenRequest struct {
	Address    string `json:"address"`
	Type       string `type:"type"`
	StartBlock uint64 `type:"startBlock"`
}

type TokenResponse struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Decimals   int    `json:"decimals"`
	StartBlock uint64 `json:"startBlock"`
	Name       string `json:"name"`
}

type TokenTypesResponse struct {
	SupportedTypes []string `json:"supportedTypes"`
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
