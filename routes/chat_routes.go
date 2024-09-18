package routes

import (
	"chat-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/api/v1/chat", controllers.CreateChatSession)
	r.POST("/api/v1/chat/:id/last-message/:message_id", controllers.UpdateLastMessage)
	r.GET("/api/v1/chat/:id", controllers.GetChatSessions)
	r.POST("/api/v1/chat/:id/message", controllers.CreateChatMessage)
	r.GET("/api/v1/chat/:id/messages", controllers.GetChatMessages)
}
