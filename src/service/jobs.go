package service

import (
	"github.com/marcoboschetti/qlaire/src/clients"
	"github.com/marcoboschetti/qlaire/src/entities"
	"github.com/marcoboschetti/qlaire/src/repository"
)

func runJob(job *entities.AdsInsightsJob) {
	// This function will be responsible for running the job in the background.

	// 1. Reques the LLM to generate Qloo's query (seed) and insight type
	job.Status = entities.AdsInsightsJobStatusGeneratingSeed
	repository.UpsertJob(job)
	seed, err := clients.LLMGenerateQlooSeed(job.JobInputs)
	if err != nil {
		job.Status = entities.AdsInsightsJobStatusFailed
		job.FinalError = err.Error()
		repository.UpsertJob(job)
		return
	}
	job.GeneratedSeed = *seed

	// 2. Qloo's entity /search with seed query to retrieve results with names, entity_ids, and short_description
	// 3. Attempt Qloo's insights with retrieved entity_ids to find popularity
	// 4. Attempt Qloo's demographic insights with retrieved entity_ids to find demographic information

	// 5. From 3 and 4 information, request the LLM to generate job results with enriched insights

}
