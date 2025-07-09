package entities

type (
	JobInputs struct {
		TargetPlatform string `json:"target_platform"`
		Product        string `json:"product"`
		Title          string `json:"title"`
	}

	GeneratedSeed struct {
		Query string `json:"query"`
		Type  string `json:"type"`
	}

	AdsInsightsJob struct {
		// ID is the unique identifier for the job
		ID string `json:"id"`

		Status     AdsInsightsJobStatus `json:"status"`
		FinalError string               `json:"final_error,omitempty"` // Optional error message if the job failed

		JobInputs     JobInputs     `json:"job_inputs"`
		GeneratedSeed GeneratedSeed `json:"generated_seed,omitempty"`
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
