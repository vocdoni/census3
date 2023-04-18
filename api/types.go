package api

type CreateTokenRequest struct {
	Address    string `json:"address"`
	Type       string `type:"type"`
	StartBlock uint64 `type:"startBlock"`
}

type TokenHoldersResponse struct {
	Holders []string `json:"holders"`
}
