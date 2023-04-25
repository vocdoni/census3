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

type CreateCensusResponse struct {
	CensusID string `json:"censusId"`
}
