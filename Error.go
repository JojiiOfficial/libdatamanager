package libdatamanager

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidResponseHeaders error on missing or malformed response headers
	ErrInvalidResponseHeaders = errors.New("Invalid response headers")
	// ErrInvalidAuthorizationMethod error if authorization method is not implemented
	ErrInvalidAuthorizationMethod = errors.New("Invalid request authorization method")
	// ErrNonOkResponsecode if requests responsecode is non http.ok
	ErrNonOkResponsecode = errors.New("non ok response")
	// ErrResponseError response returned an error
	ErrResponseError = errors.New("Response returned an error")
)

// ResponseErr response error
type ResponseErr struct {
	Response *RestRequestResponse
	Err      error
}

func (reserr *ResponseErr) Error() string {
	if reserr.Response != nil {
		return fmt.Sprintf("HTTPCode: %d; Message: %s", reserr.Response.HTTPCode, reserr.Response.Message)
	}
	if reserr.Err != nil {
		return reserr.Err.Error()
	}

	return "Unexpected error"
}

// NewErrorFromResponse return error from response
func NewErrorFromResponse(r *RestRequestResponse, err ...error) *ResponseErr {
	var (
		responseErr ResponseErr
		e           error
	)

	// use e if err passed
	if len(err) > 0 && err[0] != nil {
		e = err[0]
	}

	// Check if http.Request was successful
	if r != nil {
		// Server throw an error
		if r.Status == ResponseError && e == nil {
			e = ErrResponseError
		}

		responseErr = ResponseErr{
			Response: r,
			Err:      e,
		}
	} else {
		// http.Request throw an error
		responseErr = ResponseErr{
			Err: e,
		}
	}

	return &responseErr
}
