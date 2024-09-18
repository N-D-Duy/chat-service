package controllers

import (
	"chat-service/models"
	"chat-service/services"
	"chat-service/utils"

	"github.com/gin-gonic/gin"
)

func CreateChatMessage(c *gin.Context) {
	var chatMessage models.ChatMessage
	if err := c.ShouldBindJSON(&chatMessage); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	chatSessionId := c.Param("id")
	chatSessionIdUint, err := utils.StringToUint(chatSessionId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	chatMessage.ChatSessionID = uint(chatSessionIdUint)

	if err := services.CreateChatMessage(&chatMessage); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"successful": chatMessage})

}

func GetChatMessages(c *gin.Context) {
	chatSessionID := c.Param("id")
	chatMessages, err := services.GetChatMessages(chatSessionID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": chatMessages})

}
