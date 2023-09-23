// Package apierror contains high level api errors and interfaces.
package apierror

// API errors
var (
	// Check if your parameters are correct.
	ErrInvalidRequestBody = XCloudErr{ErrCode: "INVALID_REQUEST_BODY", Message: "invalid input parameters, try again after fixing the values"}

	// Request rate has exceeded. Try making fewer requests.
	ErrRequestThrottled = XCloudErr{ErrCode: "TOO_MANY_REQUESTS", Message: "rate limit triggered, try lowering the number of requests"}

	// ErrInternalServer is for 5XX errors.
	ErrInternalServer = XCloudErr{ErrCode: "INTERNAL_SERVER_ERROR", Message: "api call to remove server failed"}

	// ErrSDK consists of errors which cannot be categorized properly e.g. unhandled cases, response parsing failures e.t.c.
	ErrSDK = XCloudErr{ErrCode: "SDK_ERROR"}

	ErrUnhandled = XCloudErr{ErrCode: "UNHANDLED_ERROR"}
)

type XCloudErr struct {
	// StatusCode is http status code received from XCloud API e.g. 200, 400
	// StatusCode int

	// ErrCode is unique identifier for each error e.g. INTERNAL_SERVER_ERROR
	ErrCode string

	// Message is a human readable actionable message regarding the error.
	Message string

	// Error is full error object which can be unwrapped
	Err error
}

// Unwrap returns the error object
func (er *XCloudErr) UnWrap() error {
	return er.Err
}

// Error returns stringified error message. If multiple errors are wrapped, should return all errors as a combined string.
func (er *XCloudErr) Error() string {
	return er.Message
}

// Record captures context for an error e.g. detailed message and error object.
//
//	Example:
//
//	if err != nil {
//		err = xcloud.ErrInvalidRequestBody.Record("", err)
//	}
func (er *XCloudErr) Record(msg string, err error) *XCloudErr {
	if msg != "" {
		er.Message = msg
	}
	er.Err = err
	return er
}

// Is checks if the error type is of certain type or not
func (xr *XCloudErr) Is(t error) bool {
	e, ok := t.(*XCloudErr)

	if !ok {
		return false
	}

	return xr.ErrCode == e.ErrCode
}
