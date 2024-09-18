package controllers

import (
	"chat-service/models"
	"chat-service/services"

	"github.com/gin-gonic/gin"
)

func CreateChatSession(c *gin.Context) {
	var chatSession models.ChatSession
	if err := c.ShouldBindJSON(&chatSession); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateChatSession(&chatSession); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": chatSession})
}

func GetChatSessions(c *gin.Context) {
	uid := c.Param("id")
	chatSessions, err := services.GetChatSessions(uid)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": chatSessions})
}

func UpdateLastMessage(c *gin.Context) {
	chatSessionID := c.Param("id")
	messageID := c.Param("message_id")
	if err := services.UpdateLastMessage(chatSessionID, messageID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Last message updated successfully"})
}
