package main

import (
	"log"

	"github.com/ChristianSilvaDev/GoMail/src/internal"
	"github.com/ChristianSilvaDev/GoMail/src/internal/container"
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	container.ProvideRoutesV1().Setup(router)
}

func main() {
	router := gin.Default()
	
	setupRoutes(router)
	internal.Setup()

	log.Println("Server running on port 8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}