package client

import (
	"context"
	"time"
)

const (
	DefaultRetryDelay = 2 * time.Second
	DefaultMaxRetries = 5
)

type Retryer interface {
	Run(ctx context.Context, fn func(ctx context.Context) error) error
	SetMaxRetries(n int)
}

type Retry struct {
	// Delay is time to wait before until next retry
	Delay time.Duration
	// MaxRetries is the number attempts to try running given function.
	MaxRetries int
}

func DefaultRetryer() *Retry {
	return &Retry{Delay: DefaultRetryDelay, MaxRetries: DefaultMaxRetries}
}

// Run executes given function with constant backoff strategy. It uses a fixed delay window to wait after each retry and executes for 'n' times.
//
// Example:
//
//	func testFunction(ctx context.Context) (string, error) {
//		time.Sleep(time.Second)
//		// do some work here
//		return "Hi! I am done", nil
//	}
//
//	r := NewRetryer().WithMaxRetries(5)
//	ctx:= context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
//
//	execErr := r.Run(ctx, func(ctx context.Context) error {
//	// it it not compulsory to pass context to your function
//		fResp, fErr = testFunction(ctx)
//		return fErr
//	})
func (r *Retry) Run(ctx context.Context, fn func(ctx context.Context) error) error {
	var execErr error
	for i := 1; i <= r.MaxRetries; i++ {
		execErr = fn(ctx)
		if execErr == nil {
			break
		}

		timer := time.NewTimer(r.Delay)
		if err := r.sleep(ctx, timer); err != nil {
			return err
		}
		timer.Stop()
	}
	return execErr
}

// SetMaxRetries overrides default max retries but provided value is non-zero.
func (r *Retry) SetMaxRetries(n int) {
	if n > 0 {
		r.MaxRetries = n
	}
}

func (r *Retry) sleep(ctx context.Context, timer *time.Timer) error {
	select {
	// timer sends message to channel when the time limit crosses
	case <-timer.C:
		return nil
	case <-ctx.Done():
		// context is cancelled, cleanup timer
		timer.Stop()
		return ctx.Err()
	}
}

// to enforce compile type check
var _ Retryer = (*Retry)(nil)
