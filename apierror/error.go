// Package apierror contains common api errors returned by the sdk. Errors specific to each services are inside /api/{service}/error.go files.
package apierror

import "errors"

var (
	// ErrInvalidRequestBody represents error where request parameter are invalid.
	ErrInvalidRequestBody = APIError{ErrCode: "INVALID_REQUEST_BODY", Err: errors.New("invalid input parameters")}

	// ErrRequestThrottled represents server error where request rate has exceeded the limit.
	ErrRequestThrottled = APIError{ErrCode: "TOO_MANY_REQUESTS", Err: errors.New("request rate limit exceeded")}

	// ErrInternalServer represents error where server has failed to process the request.
	ErrInternalServer = APIError{ErrCode: "INTERNAL_SERVER_ERROR", Err: errors.New("server failed")}

	// ErrResponseDeserialization represents error where SDK fails to unmarshal response as JSON.
	ErrResponseDeserialization = APIError{ErrCode: "RESPONSE_DESERIALIZATION_ERROR"}

	// ErrSDK represents local errors which occurred before making call to the server.
	ErrSDK = APIError{ErrCode: "SDK_ERROR"}

	// ErrUnhandled contains errors which are unknown and are not categorized.
	ErrUnhandled = APIError{ErrCode: "UNHANDLED_ERROR"}
)

type APIError struct {
	// ErrCode is unique identifier for each error e.g. INTERNAL_SERVER_ERROR
	ErrCode string

	// Error is full error object which can be unwrapped
	Err error
}

func New(code string, err error) *APIError {
	return &APIError{code, err}
}

// Unwrap returns the error object
func (er *APIError) UnWrap() error {
	return errors.Unwrap(er.Err)
}

// Error returns stringified error message. If multiple errors are wrapped, should return all errors as a combined string.
func (er *APIError) Error() string {
	// unwrap all errors and convert to string
	return er.ErrCode + " " + er.Err.Error()
}

// Record captures the given error object in the custom error type.
//
//	Example:
//
//	if err != nil {
//		err = apierror.ErrInvalidRequestBody.Record(err)
//	}
func (er *APIError) Record(err error) *APIError {
	er.Err = err
	return er
}

// Is checks if the error type is of certain type or not
func (xr *APIError) Is(t error) bool {
	e, ok := t.(*APIError)

	if !ok {
		return false
	}

	return xr.ErrCode == e.ErrCode
}
