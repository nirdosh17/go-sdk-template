// this file contains errors related to chatai service only

package chatai

import "github.com/nirdosh17/go-sdk-template/apierror"

var (
	// ErrInputSizeLimitExceeded represents error code InputSizeLimitExceeded.
	//
	// Input string character size has exceeded. Try a smaller input string.
	ErrInputSizeLimitExceeded = apierror.XCloudErr{ErrCode: "INPUT_SIZE_EXCEEDED", Message: "max input size for question field is 200 characters"}
)
