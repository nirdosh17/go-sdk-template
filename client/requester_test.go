// Package client package provides interface to perform api call to external services.

package client

import (
	"net/http"
	"reflect"
	"testing"
)

func TestDefaultClient(t *testing.T) {
	c := &http.Client{Timeout: DefaultHTTPTimeout}
	if got := DefaultClient(); !reflect.DeepEqual(got, c) {
		t.Errorf("DefaultClient() = %v, want %v", got, c)
	}
}
