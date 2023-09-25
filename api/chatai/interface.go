package chatai

import (
	"context"

	"github.com/nirdosh17/go-sdk-template/model"
)

// Creating interface so that is can be mocked if needed
type IChatAI interface {
	AskAIWithContext(context.Context, string) (model.AIAnswer, error)
}

// making sure that ChatAI satisfies this interface
var _ IChatAI = (*ChatAPI)(nil)
