package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/marcoboschetti/qlaire/src/entities"
	"github.com/marcoboschetti/qlaire/src/repository"
)

type AdsInsightsService interface {
	// StartAdsInsightJob starts a new ads insight job
	StartAdsInsightJob(entities.JobInputs) (*entities.AdsInsightsJob, error)

	// GetAdsInsightJob gets the status of an ads insight job
	GetAdsInsightJob(jobId string) (*entities.AdsInsightsJob, error)
}

type adsInsightsService struct {
}

// NewAdsInsightsService creates a new AdsInsightsService
func NewAdsInsightsService() AdsInsightsService {
	return &adsInsightsService{}
}

func (a *adsInsightsService) StartAdsInsightJob(input entities.JobInputs) (*entities.AdsInsightsJob, error) {
	newJob := &entities.AdsInsightsJob{
		ID:        uuid.New().String(),
		Status:    entities.AdsInsightsJobStatusPending,
		JobInputs: input,
	}

	repository.AddJob(newJob)

	// Trigger run job async
	go runJob(newJob)

	return newJob, nil
}

func (a *adsInsightsService) GetAdsInsightJob(jobId string) (*entities.AdsInsightsJob, error) {
	job, exists := repository.GetJob(jobId)
	if exists {
		return job, nil
	}

	return nil, fmt.Errorf("job with ID %s not found", jobId)
}
