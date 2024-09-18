package main

import (
	"chat-service/configs"
	"chat-service/models"
	"chat-service/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDatabase()

	err := configs.DB.AutoMigrate(&models.ChatSession{}, &models.ChatMessage{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	r := gin.Default()
	routes.SetupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
