package main

import (
	"log"
	"os"

	"github.com/aleluchesi/crud_mvc/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitiRoutes(&router.RouterGroup)

	if err := router.Run(":" + os.Getenv("PORT_ROUTE")); err != nil {
		log.Fatal(err)
	}

}
