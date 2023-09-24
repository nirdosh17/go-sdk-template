// Package client package provides interface to perform api call to external services.

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/nirdosh17/go-sdk-template/apierror"
)

const (
	DefaultHTTPTimeout = 30 * time.Second
)

type Requester interface {
	Perform(ctx context.Context, url string, method string, requestBody interface{}, target interface{}) error
}

type Request struct {
	Client *http.Client
}

// TODO: rename to default request
func DefaultClient() *http.Client {
	return &http.Client{Timeout: DefaultHTTPTimeout}
}

// MakeRequest performs api call to given url with supplied arguments.
// It will include "requestBody" in the request if it is non-nil.
// Response from server will be deserialized to "target" interface.
func (r *Request) Perform(ctx context.Context, url string, method string, requestBody interface{}, target interface{}) error {
	var (
		toSend    *bytes.Buffer
		respBytes []byte
	)

	if requestBody != nil {
		b, err := json.Marshal(requestBody)
		if err != nil {
			return apierror.ErrInvalidRequestBody.Record(fmt.Errorf("serialization failure: %v", err))
		}
		toSend = bytes.NewBuffer(b)
	}

	request, err := http.NewRequestWithContext(ctx, method, url, toSend)
	if err != nil {
		return apierror.ErrInvalidRequestBody.Record(err)
	}

	resp, err := r.Client.Do(request)
	if err != nil {
		return apierror.ErrSDK.Record(fmt.Errorf("api request failure: %w", err))
	}
	defer resp.Body.Close()

	status := resp.StatusCode
	if status >= 500 {
		return &apierror.ErrInternalServer
	}

	respBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return apierror.ErrSDK.Record(fmt.Errorf("failed reading response body: %w", err))
	}

	if target != nil {
		err = json.Unmarshal(respBytes, target)
		if err != nil {
			return apierror.ErrSDK.Record(fmt.Errorf("response de-serialization err: %w", err))
		}
	}

	switch {
	case status >= 200 && status < 300:
		return nil
	case status >= 400 && status < 500:
		return apierror.ErrInvalidRequestBody.Record(fmt.Errorf("server response: %v", string(respBytes)))
	default:
		// 3XX not handled
		return apierror.ErrUnhandled.Record(fmt.Errorf("server error %d", status))
	}
}
