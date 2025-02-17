package main

import (
	"log"

	"github.com/ChristianSilvaDev/GoMail/src/internal/container"
	"github.com/gin-gonic/gin"
)

func setupRoutes() *gin.Engine {
	router := gin.Default()
	container.ProvideRoutesV1().Setup(router)
	return router
}

func main() {
	router := setupRoutes()

	log.Println("Server runnin on port 8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}