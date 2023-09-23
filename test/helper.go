package test

import (
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
