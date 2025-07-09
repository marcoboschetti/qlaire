package entities

type (
	JobInputs struct {
		TargetPlatform string `json:"target_platform"`
		Product        string `json:"product"`
		Title          string `json:"title"`
	}

	AdsInsightsJob struct {
		// ID is the unique identifier for the job
		ID string `json:"id"`

		Status     AdsInsightsJobStatus `json:"status"`
		FinalError string               `json:"final_error,omitempty"` // Optional error message if the job failed

		JobInputs          JobInputs           `json:"job_inputs"`
		GeneratedSeed      GeneratedSeed       `json:"generated_seed,omitempty"`
		SearchResults      []SearchResult      `json:"search_results,omitempty"`
		PopularityInsights []InsightEntity     `json:"insights_response,omitempty"`
		DemographicBuckets []DemographicBucket `json:"demographics,omitempty"`
	}

	GeneratedSeed struct {
		Query string `json:"query"`
		Type  string `json:"type"`
	}

	// SearchResult represents a single entity returned by the search endpoint
	SearchResult struct {
		Name      string   `json:"name"`
		EntityID  string   `json:"entity_id"`
		Types     []string `json:"types"`
		ShortDesc string   `json:"short_description"`
	}

	// InsightEntity represents a single related entity from the insights endpoint
	InsightEntity struct {
		Name       string  `json:"name"`
		EntityID   string  `json:"entity_id"`
		Subtype    string  `json:"subtype"`
		Popularity float64 `json:"popularity"`
	}

	// DemographicBucket holds demographic metrics for one seed
	DemographicBucket struct {
		EntityID string             `json:"entity_id"`
		Age      map[string]float64 `json:"age"`
		Gender   map[string]float64 `json:"gender"`
	}

	AdsInsightsJobStatus string
)

const (
	AdsInsightsJobStatusPending AdsInsightsJobStatus = "pending"

	AdsInsightsJobStatusGeneratingSeed       AdsInsightsJobStatus = "generating_seed"       // Step 1
	AdsInsightsJobStatusSearchingEntity      AdsInsightsJobStatus = "searching_entity"      // Step 2
	AdsInsightsJobStatusFetchingInsights     AdsInsightsJobStatus = "fetching_insights"     // Step 3
	AdsInsightsJobStatusFetchingDemographics AdsInsightsJobStatus = "fetching_demographics" // Step 4
	AdsInsightsJobStatusGeneratingOutput     AdsInsightsJobStatus = "generating_output"     // Step 5

	AdsInsightsJobStatusCompleted AdsInsightsJobStatus = "completed"
	AdsInsightsJobStatusFailed    AdsInsightsJobStatus = "failed"
)
