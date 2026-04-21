package main

import (
	"os"

	logger "github.com/aleluchesi/crud_mvc/src/configuration"
	"github.com/aleluchesi/crud_mvc/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err, zap.String("journey", "main"))
	}

	router := gin.Default()

	routes.InitiRoutes(&router.RouterGroup)

	if err := router.Run(":" + os.Getenv("PORT_ROUTE")); err != nil {
		logger.Error("Error starting server", err, zap.String("journey", "main"))
	}

}
