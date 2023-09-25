// Package client package provides interface to perform api call to external services.

package client

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/nirdosh17/go-sdk-template/apierror"
	"github.com/nirdosh17/go-sdk-template/logger"
	"github.com/nirdosh17/go-sdk-template/model"
	"github.com/nirdosh17/go-sdk-template/test"
)

func TestDefaultClient(t *testing.T) {
	c := &http.Client{Timeout: DefaultHTTPTimeout}
	if got := DefaultClient(); !reflect.DeepEqual(got, c) {
		t.Errorf("DefaultClient() = %v, want %v", got, c)
	}
}

func TestRequest_Perform(t *testing.T) {
	json := `{"answer": "answer from AI", "confidenceScore": 73}`
	mock := test.MockHTTPClient{
		StatusCode: 200,
		JSONBody:   &json,
		Err:        nil,
	}
	r := Request{Client: &mock, Logger: logger.NewDefaultLogger(), Debug: true}

	t.Run("success", func(t *testing.T) {
		var a model.AIAnswer
		type reqBody struct {
			query string
		}
		err := r.Perform(context.Background(), "http://api.doesnotmatter.com", "POST", &reqBody{query: "some question"}, &a)
		test.ExpectNil(t, "Request.Perform", err)
		test.ExpectEqual(t, "Request.Perform", "answer from AI 73", fmt.Sprintf("%v %v", a.Answer, a.ConfidenceScore))
	})

	t.Run("server failure", func(t *testing.T) {
		var a model.AIAnswer
		mock.StatusCode = 500
		err := r.Perform(context.Background(), "http://api.doesnotmatter.com", "POST", nil, &a)
		apiErr := err.(*apierror.APIError)
		test.ExpectEqual(t, "Request.Perform", "INTERNAL_SERVER_ERROR", apiErr.ErrCode)
	})

	t.Run("req validation error", func(t *testing.T) {
		var a model.AIAnswer
		mock.StatusCode = 422
		err := r.Perform(context.Background(), "http://api.doesnotmatter.com", "POST", nil, &a)
		apiErr := err.(*apierror.APIError)
		test.ExpectEqual(t, "Request.Perform", "INVALID_REQUEST_BODY", apiErr.ErrCode)
	})

	t.Run("300 unhandled error", func(t *testing.T) {
		var a model.AIAnswer
		mock.StatusCode = 301
		err := r.Perform(context.Background(), "http://api.doesnotmatter.com", "POST", nil, &a)
		apiErr := err.(*apierror.APIError)
		test.ExpectEqual(t, "Request.Perform", "UNHANDLED_ERROR", apiErr.ErrCode)
	})

	t.Run("response deserialization error", func(t *testing.T) {
		var a model.AIAnswer
		invalidJSON := `{"key: value}`
		mock.JSONBody = &invalidJSON
		err := r.Perform(context.Background(), "http://api.doesnotmatter.com", "POST", nil, &a)
		apiErr := err.(*apierror.APIError)
		test.ExpectEqual(t, "Request.Perform", "RESPONSE_DESERIALIZATION_ERROR", apiErr.ErrCode)
	})
}
