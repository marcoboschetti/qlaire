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

		// Status is the current status of the job (e.g., "pending", "running", "completed", "failed")
		Status AdsInsightsJobStatus `json:"status"`

		JobInputs JobInputs `json:"job_inputs"`
		// TODO: Target audience, similar concepts, target platform, etc
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
