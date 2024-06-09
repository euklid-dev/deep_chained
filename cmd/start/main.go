package main

import (
	"github.com/euklid-dev/deep_chained/db"
	"github.com/euklid-dev/deep_chained/internal/config"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// loading the config
	config.LoadConfig()

	db.ConnectToDatabase()

	defer db.SQLx.Close()

	router := gin.Default()

	router.Run(":8080")
}
