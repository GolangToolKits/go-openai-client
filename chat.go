package goopenaiclient

import "time"

// Message Message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest ChatRequest
type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// LogprobContent LogprobContent
type LogprobContent struct {
	Token       string           `json:"token"`
	Logprob     float64          `json:"logprob"`
	Bytes       []byte           `json:"bytes"`
	TopLogprobs []LogprobContent `json:"top_logprobs"`
}

// Logprob Logprob
type Logprob struct {
	Content []LogprobContent `json:"content"`
}

// Choice Choice
type Choice struct {
	Index        int64   `json:"index"`
	Message      Message `json:"message"`
	Logprobs     Logprob `json:"logprobs"`
	FinishReason string  `json:"finish_reason"`
}

// ChatUsage ChatUsage
type ChatUsage struct {
	CompletionTokens int64 `json:"completion_tokens"`
	PromptTokens     int64 `json:"prompt_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

// ChatResponse ChatResponse
type ChatResponse struct {
	ID                string    `json:"id"`
	Object            string    `json:"object"`
	Created           time.Time `json:"created"`
	Model             string    `json:"model"`
	Choices           []Choice  `json:"choices"`
	Usage             ChatUsage `json:"usage"`
	SystemFingerprint string    `json:"system_fingerprint"`
}