package client

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/nirdosh17/go-sdk-template/test"
)

func TestRetry_Run(t *testing.T) {
	var (
		errCounter, successCounter int
	)

	errfn := func(ctx context.Context) error {
		errCounter++
		return errors.New("I will throw error!")
	}
	successfn := func(ctx context.Context) error {
		successCounter++
		return nil
	}

	t.Run("SuccessFunction", func(t *testing.T) {
		r := &Retry{
			Delay:      1 * time.Second,
			MaxRetries: 1,
		}
		err := r.Run(context.Background(), successfn)
		test.ExpectEqual(t, "function execution counter", r.MaxRetries, successCounter)
		test.ExpectNil(t, "execError", err)
	})

	t.Run("ErrorFunction", func(t *testing.T) {
		r := &Retry{
			Delay:      200 * time.Millisecond,
			MaxRetries: 3,
		}
		expectedExecDuration := r.Delay * time.Duration(r.MaxRetries)

		execStart := time.Now()
		err := r.Run(context.Background(), errfn)
		actualExecDuration := time.Since(execStart)

		if actualExecDuration < expectedExecDuration {
			t.Errorf("expected total execution time (after retries) to be greater than %v but got %v", expectedExecDuration, actualExecDuration)
		}

		test.ExpectEqual(t, "function execution counter", r.MaxRetries, errCounter)
		test.ExpectNotNil(t, "execError", err)
	})
}

func Test_DefaultRetryer(t *testing.T) {
	expected := Retry{Delay: DefaultRetryDelay, MaxRetries: DefaultMaxRetries}
	if got := DefaultRetryer(); !reflect.DeepEqual(got, &expected) {
		t.Errorf("DefaultRetryer() = %v, want %v", got, expected)
	}
}
