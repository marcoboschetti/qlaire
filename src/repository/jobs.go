package repository

// A simple in-memory repository for running and finished jobs.
import (
	"sync"

	"github.com/marcoboschetti/qlaire/src/entities"
)

var (
	jobsRepository = map[string]*entities.AdsInsightsJob{}
	jobsMutex      = &sync.RWMutex{}
)

// UpsertJob adds a new job to the repository
func UpsertJob(job *entities.AdsInsightsJob) {
	jobsMutex.Lock()
	defer jobsMutex.Unlock()

	jobsRepository[job.ID] = job
	// TODO: If len(jobsRepository) > 50, prune oldest or define eviction policy
}

// GetJob retrieves a job by its ID from the repository
func GetJob(jobId string) (*entities.AdsInsightsJob, bool) {
	jobsMutex.RLock()
	defer jobsMutex.RUnlock()

	job, exists := jobsRepository[jobId]
	return job, exists
}
