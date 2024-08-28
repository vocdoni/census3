package apiclient

const (
	// Token endpoints:

	// GetTokensURI is the URI for getting tokens, it accepts a pageSize,
	// nextCursor and prevCursor
	GetTokensURI = "/tokens?pageSize=%d&nextCursor=%s&prevCursor=%s"
	// CreateTokenURI is the URI for creating a token, it not accepts any
	// parameters
	CreateTokensURI = "/tokens"
	// GetAndDeleteTokenURI is the URI for getting and launchin a token
	// deletion, it accepts the tokenID, chainID and externalID
	GetAndDeleteTokenURI = "/tokens/%s?chainID=%d&externalID=%s"
	// DeleteTokenQueueURI is the URI for deleting a token from a queue, it
	// accepts the tokenID, queueID, chainID and externalID
	DeleteTokenQueueURI = "/tokens/%s/queue/%s?chainID=%d&externalID=%s"
	// GetTokenHolderURI is the URI for getting token holders, it accepts
	// the tokenID, holderID, chainID and externalID
	GetTokenHolderURI = "/tokens/%s/holders/%s?chainID=%d&externalID=%s"
	// GetTokenTypes is the URI for getting token types, it accepts no
	// parameters
	GetTokenTypes = "/tokens/types"

	// Strategies endpoints:

	// GetStrategiesURI is the URI for getting strategies, it accepts a pageSize,
	// nextCursor and prevCursor
	GetStrategiesURI = "/strategies?pageSize=%d&nextCursor=%s&prevCursor=%s"
	// GetStrategyURI is the URI for getting a strategy, it accepts the strategyID
	GetStrategyURI = "/strategies/%d"
	// CreateStrategyURI is the URI for creating a strategy, it accepts no
	// parameters
	CreateStrategyURI = "/strategies"
	// GetTokenHoldersByStrategyURI is the URI for getting token holders of a given strategy
	GetTokenHoldersByStrategyURI = "/strategies/%d/holders?truncateByDecimals=%s"
	// GetTokenHoldersByStrategyURI is the URI for getting token holders of a given strategy
	GetTokenHoldersByStrategyQueueURI = "/strategies/%d/holders/queue/%s"

	// Censuses endpoints:

	// GetCensusURI is the URI for getting a census, it accepts the censusID
	GetCensusURI = "/censuses/%d"
	// CreateCensusURI is the URI for creating a census, it accepts no
	// parameters
	CreateCensusURI = "/censuses"
	// CreateCensusQueueURI is the URI for creating a census queue, it
	// accepts the queueID
	CreateCensusQueueURI = "/censuses/queue/%s"
	// GetCensusesByStrategyURI is the URI for getting the censuses of a
	// strategy, it accepts the strategyID
	GetCensusesByStrategyURI = "/censuses/strategy/%d"
)
