package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marcoboschetti/qlaire/src/api"
)

func main() {

	gin.SetMode(gin.DebugMode)
	godotenv.Load(".env")
	rand.Seed(time.Now().Unix())

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("$PORT must be set. Defaulted to 8080")
		port = "8080"
	}

	// Start server
	r := gin.Default()
	// // *************** API **************
	public := r.Group("/api")
	public.POST("/insights", api.StartAdsInsightJob)

	// *************** SITE **************
	// Public Static Resources
	r.GET("/", func(c *gin.Context) { http.ServeFile(c.Writer, c.Request, "./site/index.html") })
	r.GET("/favicon.ico", func(c *gin.Context) { http.ServeFile(c.Writer, c.Request, "./site/logo.ico") })
	publicSite := r.Group("/site")
	publicSite.Static("/", "./site/")

	r.Run(":" + port)
}
