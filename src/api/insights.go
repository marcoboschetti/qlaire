package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartAdsInsightJob(c *gin.Context) {
	input := struct {
		Query string `json:"query"`
	}{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"insights": nil})
}
