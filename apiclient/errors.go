package apiclient

import "fmt"

var (
	// ErrBadInputs is returned when bad inputs are provided to a method. Each
	// method will specify the inputs it expects and the conditions they must
	// meet.
	ErrBadInputs = fmt.Errorf("bad inputs provided")
	// ErrConstructingURL is returned when there is an error constructing the URL
	// to make a request to the API. The final URL is parsed and this process can
	// fail.
	ErrConstructingURL = fmt.Errorf("error constructing URL")
	// ErrEncodingRequest is returned when there is an error encoding the request
	// body to be sent to the API. The request body is encoded to JSON and this
	// process can fail.
	ErrEncodingRequest = fmt.Errorf("error encoding request")
	// ErrCreatingRequest is returned when there is an error creating the request
	// to be sent to the API. It will include the error returned by the
	// http.NewRequest
	ErrCreatingRequest = fmt.Errorf("error creating request")
	// ErrMakingRequest is returned when there is an error making the request to
	// the API. It will include the error returned by the http.Client.Do
	ErrMakingRequest = fmt.Errorf("error making request")
	// ErrNoStatusOk is returned when the server returns a non-OK status code.
	ErrNoStatusOk = fmt.Errorf("server returned non-OK status")
	// ErrDecodingResponse is returned when there is an error decoding the response
	// from the API. The response is decoded from JSON and this process can fail.
	ErrDecodingResponse = fmt.Errorf("error decoding response")
	// ErrAlreadyExists is returned when the resource already exists
	ErrAlreadyExists = fmt.Errorf("resource already exists")
)
