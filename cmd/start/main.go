package main

import (
	"os"

	"github.com/euklid-dev/deep_chained/docs"
	internal "github.com/euklid-dev/deep_chained/internal/api/alpha"
	"github.com/euklid-dev/deep_chained/internal/config"
	"github.com/euklid-dev/deep_chained/internal/langchain"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"

	_ "github.com/euklid-dev/deep_chained/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Deep Chained API
//	@version		alpha
//	@description	API for Deep Chained

// @contact.name	Deepchained
// @contact.url https;//euklid.dev
// @contact.email	tech@euklid.dev
func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize langchain
	langchain.Initialize()

	// db.ConnectToDatabase()

	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.BasePath = os.Getenv("SWAGGER_BASE_PATH")

	router := gin.Default()

	// Set up CORS configuration
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://localhost:3000", "https://localhost:3216", "https://automate-ai-c674b.web.app"},
		AllowMethods:     []string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	router.Use(cors.New(corsConfig))

	// Handle OPTIONS preflight requests
	router.OPTIONS("/*path", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, PUT, PATCH, POST, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Status(204)
	})

	// swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	alpha := router.Group("/deep-c")

	// Health Check
	alpha.GET("/health-check", internal.HealthCheck)

	router.Run(":" + config.GlobalAppConfig.APP_SERVICE_PORT)
}
