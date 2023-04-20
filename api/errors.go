package api

import (
	"fmt"

	"go.vocdoni.io/dvote/httprouter/apirest"
)

var (
	ErrTokenMalformed      = apirest.APIerror{Code: 4000, HTTPstatus: apirest.HTTPstatusBadRequest, Err: fmt.Errorf("malformed token information")}
	ErrUnknownToken        = apirest.APIerror{Code: 4001, HTTPstatus: apirest.HTTPstatusNotFound, Err: fmt.Errorf("uknown token")}
	ErrNoFoundTokenHolders = apirest.APIerror{Code: 4002, HTTPstatus: apirest.HTTPstatusNotFound, Err: fmt.Errorf("no token holders found")}
	ErrCantCreateToken     = apirest.APIerror{Code: 5000, HTTPstatus: apirest.HTTPstatusInternalErr, Err: fmt.Errorf("the token cannot be created")}
	ErrCantGetTokenHolders = apirest.APIerror{Code: 5001, HTTPstatus: apirest.HTTPstatusInternalErr, Err: fmt.Errorf("error getting token holders")}
	ErrEncodeTokenHolders  = apirest.APIerror{Code: 5002, HTTPstatus: apirest.HTTPstatusInternalErr, Err: fmt.Errorf("error encoding token holders")}
)
