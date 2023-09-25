package model

// AIAnswer is the response from chatai service.
type AIAnswer struct {
	Answer          string  `json:"answer"`
	ConfidenceScore float32 `json:"confidenceScore"`
}
