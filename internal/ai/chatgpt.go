package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mazlon/gobeyond/internal/config"
	"github.com/mazlon/gobeyond/internal/models"
)

func AskFromChatGptSingleQuestion(question *models.Questions) (models.ApiResponse, error) {
	apiKey := config.GetTheEnv("GPT_API_KEY")
	url := config.ChatGptQuestionsEndpoint
	var gptRequestBody models.GptRequestBody
	var gptApiResponse models.ApiResponse
	message := models.Message{
		Content: question.Question,
		Role:    "system",
	}
	gptRequestBody.Messages = append(gptRequestBody.Messages, message)
	gptRequestBody.Model = config.ChatGptModel
	// Construct the request payload
	// var payload = map[string]interface{}{
	// 	"model":      config.ChatGptModel,
	// 	"prompt":     question.Question,
	// 	"max_tokens": 60,
	// }
	payloadBytes, _ := json.Marshal(gptRequestBody)

	// Create a new request using http
	req, err := http.NewRequest("POST", url, bytes.NewReader(payloadBytes))
	if err != nil {
		return gptApiResponse, err
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return gptApiResponse, err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return gptApiResponse, err
	}
	err = json.Unmarshal(responseBody, &gptApiResponse)
	if err != nil {
		return gptApiResponse, err
	}
	gptApiResponse.QuestionId = question.ID
	return gptApiResponse, nil
}
