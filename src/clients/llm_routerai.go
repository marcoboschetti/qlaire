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
	openRouterURL = "https://openrouter.ai/api/v1/chat/completions"
	model         = "deepseek/deepseek-chat-v3-0324:free"
)

type routerAIClient struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// NewRouterAIClient creates a new client using OPENROUTER_AI_API_TOKEN env var
func NewRouterAIClient() LLMClient {
	apiKey := os.Getenv("OPENROUTER_AI_API_TOKEN")
	return &routerAIClient{
		baseURL:    "https://openrouter.ai/api/v1/chat/completions",
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}
}

// GenerateQlooSeed calls LLM to generate Qloo search seed
func (c *routerAIClient) LLMGenerateQlooSeed(inputs entities.JobInputs) (*entities.GeneratedSeed, error) {
	prompt := fmt.Sprintf(seedPromptTemplate,
		inputs.TargetPlatform,
		inputs.Product,
		inputs.Title,
	)
	body, err := c.callOpenRouter(prompt)
	if err != nil {
		return nil, err
	}
	raw, err := c.extractContent(body)
	if err != nil {
		return nil, err
	}
	clean := c.cleanMarkdown(raw)
	var seed entities.GeneratedSeed
	if err := json.Unmarshal([]byte(clean), &seed); err != nil {
		return nil, fmt.Errorf("unmarshal seed JSON: %w", err)
	}
	return &seed, nil
}

// GenerateAdsCampaign calls LLM to produce final campaign
func (c *routerAIClient) LLMGenerateAdsCampaign(job *entities.AdsInsightsJob) (*entities.AdsCampaign, error) {
	// format related entities list
	var relatedLines []string
	for _, e := range job.PopularityInsights {
		relatedLines = append(relatedLines, fmt.Sprintf("- %s (%s, popularity: %.2f)", e.Name, e.Subtype, e.Popularity))
	}
	relatedText := strings.Join(relatedLines, "\n")

	// format demographics
	demoBytes, _ := json.MarshalIndent(job.DemographicBuckets, "", "  ")
	// build prompt
	prompt := fmt.Sprintf(campaignPromptTemplate,
		job.JobInputs.Product,
		job.JobInputs.TargetPlatform,
		job.JobInputs.TargetPlatform,
		job.JobInputs.Product,
		job.JobInputs.Title,
		job.GeneratedSeed.Query,
		job.GeneratedSeed.Type,
		relatedText,
		string(demoBytes),
	)

	body, err := c.callOpenRouter(prompt)
	if err != nil {
		return nil, err
	}
	raw, err := c.extractContent(body)
	if err != nil {
		return nil, err
	}

	clean := c.cleanMarkdown(raw)
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(clean), &result); err != nil {
		return nil, fmt.Errorf("unmarshal campaign JSON: %w", err)
	}

	fmt.Printf("TODO: unmarshal to struct LLM campaign result: %s\n", clean)

	return nil, nil
}

// callOpenRouter sends prompt to OpenRouter and returns raw body
func (c *routerAIClient) callOpenRouter(prompt string) ([]byte, error) {
	reqBody := map[string]interface{}{
		"model": model,
		"messages": []map[string]string{{
			"role":    "user",
			"content": prompt,
		}},
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}
	req, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("LLM request failed: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	fmt.Println("LLM response status:", resp.StatusCode, ":", string(body))
	return body, err
}

// extractContent retrieves content field
func (c *routerAIClient) extractContent(body []byte) (string, error) {
	var wrap struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			}
		} `json:"choices"`
	}
	if err := json.Unmarshal(body, &wrap); err != nil {
		return "", fmt.Errorf("unmarshal LLM response: %w", err)
	}
	if len(wrap.Choices) == 0 {
		return "", fmt.Errorf("no content returned by the LLM")
	}
	return wrap.Choices[0].Message.Content, nil
}

// cleanMarkdown strips markdown wrappers and extracts the JSON block
func (c *routerAIClient) cleanMarkdown(raw string) string {
	// Remove everything before the first ```json
	if idx := strings.Index(raw, "```json"); idx != -1 {
		raw = raw[idx+len("```json"):]
	}
	// Remove everything after the last ```
	if idx := strings.LastIndex(raw, "```"); idx != -1 {
		raw = raw[:idx]
	}
	// Trim whitespace and return
	return strings.TrimSpace(raw)
}
