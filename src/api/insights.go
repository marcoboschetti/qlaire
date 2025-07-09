package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcoboschetti/qlaire/src/entities"
	"github.com/marcoboschetti/qlaire/src/service"
)

type AdsInsightsController interface {
	// StartAdsInsightJob starts a new ads insight job
	StartAdsInsightJob(c *gin.Context)

	// GetAdsInsightJob gets the status of an ads insight job
	GetAdsInsightJob(c *gin.Context)
}

type adsInsightsController struct {
	service service.AdsInsightsService
}

// NewAdsInsightsController creates a new AdsInsightsController
func NewAdsInsightsController() AdsInsightsController {
	return &adsInsightsController{
		service: service.NewAdsInsightsService(),
	}
}

// GetAdsInsightJob gets the status of an ads insight job
func (a *adsInsightsController) GetAdsInsightJob(c *gin.Context) {
	jobId := c.Param("jobId")

	if jobId == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Job ID is required"})
		return
	}

	job, err := a.service.GetAdsInsightJob(jobId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if job == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Job not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"job": job})
}

func (a *adsInsightsController) StartAdsInsightJob(c *gin.Context) {
	input := entities.JobInputs{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	newJob, err := a.service.StartAdsInsightJob(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"insights": newJob})
}
