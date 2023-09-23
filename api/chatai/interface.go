package chatai

import "context"

// Creating interface so that is can be mocked if needed
type IChatAI interface {
	AskAIWithContext(context.Context, string) (AIAnswer, error)
}

// making sure that ChatAI satisfies this interface
var _ IChatAI = (*ChatAPI)(nil)
