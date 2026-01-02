package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	aiAPIURL  = "https://svceai.site/api/chat"
	aiTimeout = 30 * time.Second
)

type aiRequest struct {
	Message string   `json:"message"`
	History []string `json:"history"`
}

// callAI sends a prompt to the AI API and returns the response.
func callAI(prompt string) (string, error) {
	reqBody, err := json.Marshal(aiRequest{Message: prompt, History: []string{}})
	if err != nil {
		return "", fmt.Errorf("failed to marshal AI request: %w", err)
	}

	client := &http.Client{Timeout: aiTimeout}
	req, err := http.NewRequest("POST", aiAPIURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("failed to create AI request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to call AI API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("AI API returned non-OK status: %s", resp.Status)
	}

	var respBody bytes.Buffer
	if _, err := respBody.ReadFrom(resp.Body); err != nil {
		return "", fmt.Errorf("failed to read AI response body: %w", err)
	}

	return respBody.String(), nil
}
