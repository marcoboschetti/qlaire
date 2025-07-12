package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
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

	// CORS configuration
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "http://localhost:3000,http://localhost:8080"
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{allowedOrigins}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true

	r.Use(cors.New(config))

	adsInsightsController := api.NewAdsInsightsController()

	// // *************** API **************
	public := r.Group("/v1/api/ads")
	public.POST("/insights", adsInsightsController.StartAdsInsightJob)
	public.GET("/insights/:jobId", adsInsightsController.GetAdsInsightJob)

	// Health check endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Qlaire API is running",
		})
	})

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
