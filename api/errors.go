package api

import (
	"fmt"
	"net/http"

	"go.vocdoni.io/dvote/httprouter/apirest"
)

var (
	ErrMalformedToken = apirest.APIerror{
		Code:       4000,
		HTTPstatus: apirest.HTTPstatusBadRequest,
		Err:        fmt.Errorf("malformed token information"),
	}
	ErrMalformedCensusID = apirest.APIerror{
		Code:       4001,
		HTTPstatus: apirest.HTTPstatusBadRequest,
		Err:        fmt.Errorf("malformed census ID, it must be a integer"),
	}
	ErrMalformedStrategyID = apirest.APIerror{
		Code:       4002,
		HTTPstatus: apirest.HTTPstatusBadRequest,
		Err:        fmt.Errorf("malformed strategy ID, it must be an integer"),
	}
	ErrNotFoundToken = apirest.APIerror{
		Code:       4003,
		HTTPstatus: apirest.HTTPstatusNotFound,
		Err:        fmt.Errorf("no token found"),
	}
	ErrNotFoundTokenHolders = apirest.APIerror{
		Code:       4004,
		HTTPstatus: apirest.HTTPstatusNotFound,
		Err:        fmt.Errorf("no token holders found"),
	}
	ErrNotFoundStrategy = apirest.APIerror{
		Code:       4005,
		HTTPstatus: apirest.HTTPstatusNotFound,
		Err:        fmt.Errorf("no strategy found with the ID provided"),
	}
	ErrNotFoundCensus = apirest.APIerror{
		Code:       4006,
		HTTPstatus: apirest.HTTPstatusNotFound,
		Err:        fmt.Errorf("census not found"),
	}
	ErrNoTokens = apirest.APIerror{
		Code:       4007,
		HTTPstatus: apirest.HTTPstatusNoContent,
		Err:        fmt.Errorf("no tokens found"),
	}
	ErrNoStrategies = apirest.APIerror{
		Code:       4008,
		HTTPstatus: apirest.HTTPstatusNoContent,
		Err:        fmt.Errorf("no strategy found"),
	}
	ErrTokenAlreadyExists = apirest.APIerror{
		Code:       4009,
		HTTPstatus: http.StatusConflict,
		Err:        fmt.Errorf("token already created"),
	}
	ErrNoStrategyTokens = apirest.APIerror{
		Code:       4010,
		HTTPstatus: apirest.HTTPstatusBadRequest,
		Err:        fmt.Errorf("no tokens found for the strategy provided"),
	}
	ErrMalformedCensusQueueID = apirest.APIerror{
		Code:       4011,
		HTTPstatus: apirest.HTTPstatusBadRequest,
		Err:        fmt.Errorf("malformed queue ID"),
	}
	ErrCensusAlreadyExists = apirest.APIerror{
		Code:       4012,
		HTTPstatus: http.StatusConflict,
		Err:        fmt.Errorf("census already exists"),
	}
	ErrChainIDNotSupported = apirest.APIerror{
		Code:       4013,
		HTTPstatus: apirest.HTTPstatusBadRequest,
		Err:        fmt.Errorf("chain ID provided not supported"),
	}
	ErrMalformedPagination = apirest.APIerror{
		Code:       4014,
		HTTPstatus: apirest.HTTPstatusBadRequest,
		Err:        fmt.Errorf("malformed pagination params"),
	}
	ErrCantCreateToken = apirest.APIerror{
		Code:       5000,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("the token cannot be created"),
	}
	ErrCantCreateCensus = apirest.APIerror{
		Code:       5001,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error creating the census tree on the census database"),
	}
	ErrCantAddHoldersToCensus = apirest.APIerror{
		Code:       5002,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error adding the holders to the created census"),
	}
	ErrPruningCensus = apirest.APIerror{
		Code:       5003,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error pruning the current census tree"),
	}
	ErrCantGetToken = apirest.APIerror{
		Code:       5004,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error getting token information"),
	}
	ErrCantGetTokens = apirest.APIerror{
		Code:       5005,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error getting tokens information"),
	}
	ErrCantGetTokenHolders = apirest.APIerror{
		Code:       5006,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error getting token holders"),
	}
	ErrCantGetStrategy = apirest.APIerror{
		Code:       5007,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error getting strategy information"),
	}
	ErrCantGetStrategies = apirest.APIerror{
		Code:       5008,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error getting strategies information"),
	}
	ErrCantGetCensus = apirest.APIerror{
		Code:       5009,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error getting census information"),
	}
	ErrEncodeToken = apirest.APIerror{
		Code:       5010,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error encoding token"),
	}
	ErrEncodeTokens = apirest.APIerror{
		Code:       5011,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error encoding tokens"),
	}
	ErrEncodeTokenTypes = apirest.APIerror{
		Code:       5012,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error encoding supported tokens types"),
	}
	ErrEncodeTokenHolders = apirest.APIerror{
		Code:       5013,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error encoding token holders"),
	}
	ErrEncodeStrategyHolders = apirest.APIerror{
		Code:       5014,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error encoding strategy holders"),
	}
	ErrEncodeStrategy = apirest.APIerror{
		Code:       5015,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error encoding strategy"),
	}
	ErrEncodeStrategies = apirest.APIerror{
		Code:       5016,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error encoding strategies"),
	}
	ErrEncodeCensus = apirest.APIerror{
		Code:       5017,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error encoding census"),
	}
	ErrEncodeCensuses = apirest.APIerror{
		Code:       5018,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error encoding censuses"),
	}
	ErrInitializingWeb3 = apirest.APIerror{
		Code:       5019,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error initialising web3 client"),
	}
	ErrCantGetTokenCount = apirest.APIerror{
		Code:       5020,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error getting number of token holders"),
	}
	ErrCantGetLastBlockNumber = apirest.APIerror{
		Code:       5021,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error getting last block number from web3 endpoint"),
	}
	ErrEncodeQueueItem = apirest.APIerror{
		Code:       5022,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error encoding census queue item"),
	}
	ErrEncodeAPIInfo = apirest.APIerror{
		Code:       5023,
		HTTPstatus: apirest.HTTPstatusInternalErr,
		Err:        fmt.Errorf("error encoding API info"),
	}
)
