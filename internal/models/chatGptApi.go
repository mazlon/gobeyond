package models

import "github.com/google/uuid"

type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type GptRequestBody struct {
	Messages []Message `json:"messages"`
	Model    string    `json:"model"`
}

type Choice struct {
	Index        int         `json:"index"`
	Message      Message     `json:"message"`
	Logprobs     interface{} `json:"logprobs"` // null or other type, hence interface{}
	FinishReason string      `json:"finish_reason"`
}

// Define the Usage struct for the "usage" object
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// Define the ApiResponse struct for the top-level response object
type ApiResponse struct {
	Id                string      `json:"id"`
	QuestionId        uuid.UUID   `json:"question_id"`
	Object            string      `json:"object"`
	Created           int         `json:"created"`
	Model             string      `json:"model"`
	Choices           []Choice    `json:"choices"`
	Usage             Usage       `json:"usage"`
	SystemFingerprint interface{} `json:"system_fingerprint"` // can be null, hence interface{}
}
