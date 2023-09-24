package chatai

import (
	"fmt"

	"github.com/nirdosh17/go-sdk-template/apierror"
)

var (
	// ErrInputSizeLimitExceeded represents error where user's input is larger than specified limit.
	ErrInputSizeLimitExceeded = apierror.New("INPUT_SIZE_EXCEEDED", fmt.Errorf("input size exceeded the limit of %d characters", MaxInputLength))
)
