package services

import (
	"chat-service/configs"
	"chat-service/models"
)

func CreateChatMessage(message *models.ChatMessage) error {
	result := configs.DB.Create(&message)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetChatMessages(chatSessionID string) ([]models.ChatMessage, error) {
	var messages []models.ChatMessage
	result := configs.DB.Where("chat_session_id = ?", chatSessionID).Find(&messages)

	if result.Error != nil {
		return nil, result.Error
	}
	return messages, nil
}
