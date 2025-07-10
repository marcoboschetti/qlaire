package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marcoboschetti/qlaire/src/api"
)

func main() {

	gin.SetMode(gin.DebugMode)
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("$PORT must be set. Defaulted to 8080")
		port = "8080"
	}

	// Start server
	r := gin.Default()
	adsInsightsController := api.NewAdsInsightsController()

	// // *************** API **************
	public := r.Group("/v1/api/ads")
	public.POST("/insights", adsInsightsController.StartAdsInsightJob)
	public.GET("/insights/:jobId", adsInsightsController.GetAdsInsightJob)

	// *************** SITE **************
	// Public Static Resources
	r.GET("/", func(c *gin.Context) { http.ServeFile(c.Writer, c.Request, "./site/index.html") })
	r.GET("/favicon.ico", func(c *gin.Context) { http.ServeFile(c.Writer, c.Request, "./site/favicon.ico") })

	// Serve React app static files
	r.Static("/static", "./site/static")
	r.StaticFile("/manifest.json", "./site/manifest.json")

	// Catch-all route for React Router
	r.NoRoute(func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "./site/index.html")
	})

	r.Run(":" + port)
}
