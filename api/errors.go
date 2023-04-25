package api

import (
	"fmt"

	"go.vocdoni.io/dvote/httprouter/apirest"
)

var (
	ErrMalformedToken         = apirest.APIerror{Code: 4000, HTTPstatus: apirest.HTTPstatusBadRequest, Err: fmt.Errorf("malformed token information")}
	ErrUnknownToken           = apirest.APIerror{Code: 4001, HTTPstatus: apirest.HTTPstatusNotFound, Err: fmt.Errorf("uknown token")}
	ErrNoFoundTokenHolders    = apirest.APIerror{Code: 4002, HTTPstatus: apirest.HTTPstatusNotFound, Err: fmt.Errorf("no token holders found")}
	ErrMalformedStrategyID    = apirest.APIerror{Code: 4003, HTTPstatus: apirest.HTTPstatusBadRequest, Err: fmt.Errorf("malformed strategy ID, it must be a integer")}
	ErrNotFoundStrategy       = apirest.APIerror{Code: 4004, HTTPstatus: apirest.HTTPstatusNotFound, Err: fmt.Errorf("no strategy found with the ID provided")}
	ErrCantCreateToken        = apirest.APIerror{Code: 5000, HTTPstatus: apirest.HTTPstatusInternalErr, Err: fmt.Errorf("the token cannot be created")}
	ErrCantGetTokenHolders    = apirest.APIerror{Code: 5001, HTTPstatus: apirest.HTTPstatusInternalErr, Err: fmt.Errorf("error getting token holders")}
	ErrEncodeTokenHolders     = apirest.APIerror{Code: 5002, HTTPstatus: apirest.HTTPstatusInternalErr, Err: fmt.Errorf("error encoding token holders")}
	ErrCantGetStrategy        = apirest.APIerror{Code: 5003, HTTPstatus: apirest.HTTPstatusInternalErr, Err: fmt.Errorf("error getting strategy information")}
	ErrCantCreateCensus       = apirest.APIerror{Code: 5004, HTTPstatus: apirest.HTTPstatusInternalErr, Err: fmt.Errorf("error creating the census tree on the census database")}
	ErrCantAddHoldersToCensus = apirest.APIerror{Code: 5005, HTTPstatus: apirest.HTTPstatusInternalErr, Err: fmt.Errorf("error adding the holders to the created census")}
	ErrPruningCensus          = apirest.APIerror{Code: 5006, HTTPstatus: apirest.HTTPstatusInternalErr, Err: fmt.Errorf("error pruning the current census tree")}
	ErrEncodeStrategyHolders  = apirest.APIerror{Code: 5007, HTTPstatus: apirest.HTTPstatusInternalErr, Err: fmt.Errorf("error encoding strategy holders")}
)
