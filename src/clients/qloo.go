package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/marcoboschetti/qlaire/src/entities"
)

type QlooClient interface {
	// Search for entities by query and type
	Search(query, entityType string) ([]entities.SearchResult, error)
	// GetInsights fetches related entities for a given entityID and type
	GetInsights(entityIDs []string, entityType string) ([]entities.InsightEntity, error)
	// GetDemographics fetches demographic insights for a given entityID
	GetDemographics(entityIDs []string) ([]entities.DemographicBucket, error)
}

// qlooClient handles communication with the Qloo Insights API
// Used for entity search, insights, and demographics
type qlooClient struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// NewQlooClient creates a new client using QLOO_API_KEY env var
func NewQlooClient() QlooClient {
	apiKey := os.Getenv("QLOO_HACKATHON_API_TOKEN")
	return &qlooClient{
		baseURL:    "https://hackathon.api.qloo.com",
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}
}

// Search searches for entities by query and type (urn:entity:...)
func (c *qlooClient) Search(query, entityType string) ([]entities.SearchResult, error) {
	u, _ := url.Parse(c.baseURL + "/search")
	q := u.Query()
	q.Set("query", query)
	q.Set("types", entityType)
	u.RawQuery = q.Encode()

	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("x-api-key", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("search request failed: %w", err)
	}
	defer resp.Body.Close()

	var wrapper struct {
		Results []struct {
			Name       string   `json:"name"`
			EntityID   string   `json:"entity_id"`
			Types      []string `json:"types"`
			Properties struct {
				ShortDescription string `json:"short_description"`
			} `json:"properties"`
		} `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, fmt.Errorf("failed to decode search response: %w", err)
	}

	var results []entities.SearchResult
	for _, r := range wrapper.Results {
		results = append(results, entities.SearchResult{
			Name:      r.Name,
			EntityID:  r.EntityID,
			Types:     r.Types,
			ShortDesc: r.Properties.ShortDescription,
		})
	}
	return results, nil
}

// GetInsights fetches related entities for a given entityID and type
func (c *qlooClient) GetInsights(entityIDs []string, entityType string) ([]entities.InsightEntity, error) {
	u, _ := url.Parse(c.baseURL + "/v2/insights")
	q := u.Query()
	q.Set("filter.type", entityType)
	q.Set("filter.entities", strings.Join(entityIDs, ","))
	u.RawQuery = q.Encode()

	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("x-api-key", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("insights request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("insights request failed with status %s", resp.Status)
	}

	var wrapper struct {
		Results struct {
			Entities []struct {
				Name       string  `json:"name"`
				EntityID   string  `json:"entity_id"`
				Subtype    string  `json:"subtype"`
				Popularity float64 `json:"popularity"`
			} `json:"entities"`
		} `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, fmt.Errorf("failed to decode insights response: %w", err)
	}

	var insights []entities.InsightEntity
	for _, e := range wrapper.Results.Entities {
		insights = append(insights, entities.InsightEntity{
			Name:       e.Name,
			EntityID:   e.EntityID,
			Subtype:    e.Subtype,
			Popularity: e.Popularity,
		})
	}
	return insights, nil
}

// GetDemographics fetches demographic insights for a given entityID
func (c *qlooClient) GetDemographics(entityIDs []string) ([]entities.DemographicBucket, error) {
	u, _ := url.Parse(c.baseURL + "/v2/insights")
	q := u.Query()
	q.Set("filter.type", "urn:demographics")
	q.Set("signal.interests.entities", strings.Join(entityIDs, ","))
	u.RawQuery = q.Encode()

	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("x-api-key", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("demographics request failed: %w", err)
	}
	defer resp.Body.Close()

	var wrapper struct {
		Results struct {
			Demographics []struct {
				EntityID string `json:"entity_id"`
				Query    struct {
					Age    map[string]float64 `json:"age"`
					Gender map[string]float64 `json:"gender"`
				} `json:"query"`
			} `json:"demographics"`
		} `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, fmt.Errorf("failed to decode demographics response: %w", err)
	}

	var buckets []entities.DemographicBucket
	for _, d := range wrapper.Results.Demographics {
		buckets = append(buckets, entities.DemographicBucket{
			EntityID: d.EntityID,
			Age:      d.Query.Age,
			Gender:   d.Query.Gender,
		})
	}
	return buckets, nil
}
