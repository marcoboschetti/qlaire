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
Product description: "%s"  
Product title: "%s"

Requirements:
1. Choose **one** seed entity (movie, videogame, book, artist, brand, etc.) **most closely aligned** with the product’s genre, tone, and target audience. It must be a well‑known example that directly reflects the user’s description.
2. Return only the exact name of that entity.
3. Also return its Qloo entity type from the following list:

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
	campaignPromptTemplate = `You are a senior digital marketing strategist.
Client: generate a Meta Ads campaign for "%s" (%s).

User Inputs:
- Platform: %s
- Product: %s
- Title: %s

Seed Entity:
- Query: %s
- Type: %s

Top Related Entities (from Qloo Insights):
%s

Demographic Insights:
%s

Tasks:
1. Summarize the audience persona based on demographics.
2. Recommend audience segmentation rules for the platform.
3. Provide 2 ad headlines and descriptions (tone: relevant to product).
4. Suggest creative concepts (image/video ideas).
5. Outline campaign settings: objective, placements, budget, A/B testing.

**Output** in JSON with keys: persona_summary, segmentation, ad_copy, creative_concepts, campaign_config.
`
)

type RouterAIClient interface {
	LLMGenerateQlooSeed(inputs entities.JobInputs) (*entities.GeneratedSeed, error)
	LLMGenerateAdsCampaign(job *entities.AdsInsightsJob) (map[string]interface{}, error)
}

type routerAIClient struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// NewRouterAIClient creates a new client using OPENROUTER_AI_API_TOKEN env var
func NewRouterAIClient() RouterAIClient {
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
	raw, err := extractContent(body)
	if err != nil {
		return nil, err
	}
	clean := cleanMarkdown(raw)
	var seed entities.GeneratedSeed
	if err := json.Unmarshal([]byte(clean), &seed); err != nil {
		return nil, fmt.Errorf("unmarshal seed JSON: %w", err)
	}
	return &seed, nil
}

// GenerateAdsCampaign calls LLM to produce final campaign
func (c *routerAIClient) LLMGenerateAdsCampaign(job *entities.AdsInsightsJob) (map[string]interface{}, error) {
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
	raw, err := extractContent(body)
	if err != nil {
		return nil, err
	}
	clean := cleanMarkdown(raw)
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(clean), &result); err != nil {
		return nil, fmt.Errorf("unmarshal campaign JSON: %w", err)
	}

	fmt.Printf("LLM campaign result: %s\n", clean)
	return result, nil
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
	return io.ReadAll(resp.Body)
}

// extractContent retrieves content field
func extractContent(body []byte) (string, error) {
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
		return "", fmt.Errorf("no choices returned by the LLM")
	}
	return wrap.Choices[0].Message.Content, nil
}

// cleanMarkdown strips markdown wrappers
func cleanMarkdown(raw string) string {
	clean := strings.TrimSpace(raw)
	clean = strings.TrimPrefix(clean, "```json")
	clean = strings.TrimPrefix(clean, "```")
	clean = strings.TrimSuffix(clean, "```")
	return strings.TrimSpace(clean)
}
