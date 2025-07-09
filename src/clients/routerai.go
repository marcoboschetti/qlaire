package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/marcoboschetti/qlaire/src/entities"
)

const (
	openRouterURL      = "https://openrouter.ai/api/v1/chat/completions"
	model              = "deepseek/deepseek-chat-v3-0324:free"
	seedPromptTemplate = `You are an AI assistant tasked with finding the single most culturally relevant reference for a user’s ad campaign.

User’s ad platform: %s  
Product description: “%s”  
Product title: “%s”

Requirements:
1. Choose **one** seed entity (movie, videogame, book, artist, brand, etc.) **most closely aligned** with the product’s genre, tone, and target audience. It must be a well‑known example that directly reflects the user’s description.
2. Return only the exact name of that entity.
3. Also return its Qloo entity type from this list:

- urn:entity:actor  
- urn:entity:album  
- urn:entity:artist  
- urn:entity:author  
- urn:entity:book  
- urn:entity:brand  
- urn:entity:destination  
- urn:entity:director  
- urn:entity:locality  
- urn:entity:movie  
- urn:entity:person  
- urn:entity:place  
- urn:entity:podcast  
- urn:entity:tv_show  
- urn:entity:videogame  

**Output format** (raw JSON only, no markdown):

{
  "query": "seed entity name here",
  "type": "urn:entity:..."
}`
)

var openRouterToken string

// GenerateQlooSeed calls the OpenRouter LLM to generate a Qloo search seed and entity type
func LLMGenerateQlooSeed(inputs entities.JobInputs) (*entities.GeneratedSeed, error) {
	if openRouterToken == "" {
		openRouterToken = os.Getenv("OPENROUTER_AI_API_TOKEN")
	}

	prompt := buildSeedPrompt(inputs)

	body, err := callOpenRouter(prompt)
	if err != nil {
		return nil, err
	}

	rawContent, err := extractContent(body)
	if err != nil {
		return nil, err
	}

	cleanJSON := cleanMarkdown(rawContent)

	var result entities.GeneratedSeed
	if err := json.Unmarshal([]byte(cleanJSON), &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cleaned JSON: %w", err)
	}

	return &result, nil
}

func buildSeedPrompt(inputs entities.JobInputs) string {
	return fmt.Sprintf(seedPromptTemplate,
		inputs.TargetPlatform,
		inputs.Product,
		inputs.Title,
	)
}

func callOpenRouter(prompt string) ([]byte, error) {
	reqBody := map[string]interface{}{
		"model": model,
		"messages": []map[string]string{{
			"role":    "user",
			"content": prompt,
		}},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", openRouterURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openRouterToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("LLM request failed: %w", err)
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func extractContent(body []byte) (string, error) {
	var wrapper struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.Unmarshal(body, &wrapper); err != nil {
		return "", fmt.Errorf("failed to unmarshal OpenRouter response: %w", err)
	}
	if len(wrapper.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}
	return wrapper.Choices[0].Message.Content, nil
}

func cleanMarkdown(raw string) string {
	clean := strings.TrimSpace(raw)
	clean = strings.TrimPrefix(clean, "```json")
	clean = strings.TrimPrefix(clean, "```")
	clean = strings.TrimSuffix(clean, "```")
	return strings.TrimSpace(clean)
}
