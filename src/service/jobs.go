package service

import (
	"fmt"

	"github.com/marcoboschetti/qlaire/src/clients"
	"github.com/marcoboschetti/qlaire/src/entities"
	"github.com/marcoboschetti/qlaire/src/repository"
)

type AdsInsightsJobsService interface {
	TriggerJobProcessing(job *entities.AdsInsightsJob)
}

type adsInsightsJobsService struct {
	qlooClient clients.QlooClient
	llmClient  clients.LLMClient
}

// NewAdsInsightsJobsService creates a new AdsInsightsJobsService
func NewAdsInsightsJobsService() AdsInsightsJobsService {
	return &adsInsightsJobsService{
		qlooClient: clients.NewQlooClient(),
		llmClient:  clients.NewGroqClient(),
		// openRouterURL: clients.NewRouterAIClient(),
	}
}

// TriggerJobProcessing runs all the steps of the AdsInsights job
func (a *adsInsightsJobsService) TriggerJobProcessing(job *entities.AdsInsightsJob) {
	// 1. Generate Qloo seed via LLM
	job.Status = entities.AdsInsightsJobStatusGeneratingSeed
	repository.UpsertJob(job)
	seed, err := a.llmClient.LLMGenerateQlooSeed(job.JobInputs)
	if err != nil {
		job.Status = entities.AdsInsightsJobStatusFailed
		job.FinalError = err.Error()
		repository.UpsertJob(job)
		return
	}
	job.GeneratedSeed = *seed

	// 2. Search entities via Qloo
	job.Status = entities.AdsInsightsJobStatusSearchingEntity
	repository.UpsertJob(job)
	searchResults, err := a.qlooClient.Search(seed.Query, seed.Type)
	if err != nil {
		job.Status = entities.AdsInsightsJobStatusFailed
		job.FinalError = fmt.Sprintf("search error: %v", err)
		repository.UpsertJob(job)
		return
	}
	job.SearchResults = searchResults

	if len(searchResults) > 0 {
		entityIds := a.getEntityIDs(searchResults)

		// 3. Fetch popularity insights for each entity
		job.Status = entities.AdsInsightsJobStatusFetchingInsights
		repository.UpsertJob(job)

		allInsights, err := a.qlooClient.GetInsights(entityIds, seed.Type)
		if err != nil {
			job.Status = entities.AdsInsightsJobStatusFailed
			job.FinalError = fmt.Sprintf("insights error: %v", err)
			repository.UpsertJob(job)
			return
		}
		job.PopularityInsights = allInsights

		// 4. Fetch demographics for top entity
		job.Status = entities.AdsInsightsJobStatusFetchingDemographics
		repository.UpsertJob(job)
		demoResp, err := a.qlooClient.GetDemographics(entityIds)
		if err != nil {
			job.Status = entities.AdsInsightsJobStatusFailed
			job.FinalError = fmt.Sprintf("demographics error: %v", err)
			repository.UpsertJob(job)
			return
		}
		job.DemographicBuckets = demoResp
	}

	// 5. Generate final enriched insights via LLM
	job.Status = entities.AdsInsightsJobStatusGeneratingOutput
	repository.UpsertJob(job)

	output, err := a.llmClient.LLMGenerateAdsCampaign(job)
	if err != nil {
		job.Status = entities.AdsInsightsJobStatusFailed
		job.FinalError = fmt.Sprintf("LLM ads campaign generation error: %v", err)
		repository.UpsertJob(job)
		return
	}
	job.Status = entities.AdsInsightsJobStatusCompleted
	job.AdsCampaignResult = output
}

// getEntityIDs extracts entity IDs from search results
func (a *adsInsightsJobsService) getEntityIDs(results []entities.SearchResult) []string {
	ids := make([]string, len(results))
	for i, r := range results {
		ids[i] = r.EntityID
	}
	return ids
}
