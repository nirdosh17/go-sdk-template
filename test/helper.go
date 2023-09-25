// Package test contains helper methods for tests.
package test

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func ExpectEqual(t *testing.T, field string, expected, received interface{}) {
	if received != expected {
		t.Errorf("expected %v to be %v but received %v", field, expected, received)
	}
}

func ExpectNil(t *testing.T, field string, value interface{}) {
	if value != nil {
		t.Errorf("expected %v to be nil but received %v", field, value)
	}
}

func ExpectNotNil(t *testing.T, field string, value interface{}) {
	if value == nil {
		t.Errorf("expected %v to be not nil but received %v", field, value)
	}
}

func ExpectSameType(t *testing.T, field string, expected, received interface{}) {
	rt := reflect.TypeOf(received)
	et := reflect.TypeOf(expected)
	if rt != et {
		t.Errorf("expected %v to be type of %v but received %v", field, et, rt)
	}
}

// MockHTTPClient mocks http client for testing. It satisfies client.HTTPClient interface.
type MockHTTPClient struct {
	// send body which you want to receive in the response
	JSONBody *string
	// response status code
	StatusCode int
	// returns error if provided
	Err error
}

func (c *MockHTTPClient) Do(r *http.Request) (*http.Response, error) {
	var body io.ReadCloser
	if c.JSONBody != nil {
		body = io.NopCloser(bytes.NewReader([]byte(*c.JSONBody)))
	}

	return &http.Response{
		StatusCode: c.StatusCode,
		Body:       body,
	}, c.Err
}
