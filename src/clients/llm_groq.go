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
	groqURL   = "https://api.groq.com/openai/v1/chat/completions"
	groqModel = "meta-llama/llama-4-scout-17b-16e-instruct"
)

type groqClient struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// NewGroqClient creates a new client using GROQ_API_TOKEN env var
func NewGroqClient() LLMClient {
	apiKey := os.Getenv("GROQ_API_TOKEN")
	return &groqClient{
		baseURL:    "https://api.groq.com/openai/v1/chat/completions",
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}
}

// GenerateQlooSeed calls LLM to generate Qloo search seed
func (c *groqClient) LLMGenerateQlooSeed(inputs entities.JobInputs) (*entities.GeneratedSeed, error) {
	prompt := fmt.Sprintf(seedPromptTemplate,
		inputs.TargetPlatform,
		inputs.Product,
		inputs.Title,
	)
	body, err := c.callGroq(prompt)
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

func (c *groqClient) LLMGenerateAdsCampaign(job *entities.AdsInsightsJob) (*entities.AdsCampaign, error) {
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

	// call LLM
	body, err := c.callGroq(prompt)
	if err != nil {
		return nil, err
	}

	raw, err := c.extractContent(body)
	if err != nil {
		return nil, err
	}

	clean := c.cleanMarkdown(raw)

	adsCampaignWrapper := struct {
		AdsCampaignResult entities.AdsCampaign `json:"ads_campaign_result"`
	}{}

	if err := json.Unmarshal([]byte(clean), &adsCampaignWrapper); err != nil {
		return nil, fmt.Errorf("unmarshal AdsCampaign JSON: %w", err)
	}

	return &adsCampaignWrapper.AdsCampaignResult, nil
}

// callGroq sends the given prompt to the Groq chat completions endpoint and returns the raw response body.
func (c *groqClient) callGroq(prompt string) ([]byte, error) {
	// build the request payload
	reqBody := map[string]interface{}{
		"model": groqModel,
		"messages": []map[string]string{{
			"role":    "user",
			"content": prompt,
		}},
	}

	// marshal to JSON
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	// create HTTP request against Groq endpoint
	req, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	// set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	// execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("LLM request failed: %w", err)
	}
	defer resp.Body.Close()

	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	// optional: log status and body for debugging
	fmt.Println("Groq response status:", resp.StatusCode)
	fmt.Println("Groq response body:", string(body))

	// check for non-2xx
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("groq API error %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// extractContent retrieves content field
func (c *groqClient) extractContent(body []byte) (string, error) {
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
func (c *groqClient) cleanMarkdown(raw string) string {
	// Remove everything before the first ```json
	if idx := strings.Index(raw, "```"); idx != -1 {
		raw = raw[idx+len("```"):]
	}
	// Remove everything after the last ```
	if idx := strings.LastIndex(raw, "```"); idx != -1 {
		raw = raw[:idx]
	}
	// Trim whitespace and return
	return strings.TrimSpace(raw)
}
