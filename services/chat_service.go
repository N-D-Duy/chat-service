package services

import (
	"chat-service/configs"
	"chat-service/models"
	"log"
)

/*
GetChatSessions retrieves all chat sessions for a given user ID.
It joins the chat_sessions and chat_messages tables to find sessions
where the user is either the sender or receiver of messages.
The results are grouped by chat session ID to avoid duplicates.

Parameters:
- uid: The user ID for which to retrieve chat sessions.

Returns:
- A slice of ChatSession models representing the chat sessions.
- An error if there is an issue with the database query.
*/
func GetChatSessions(uid string) ([]models.ChatSession, error) {
	var chatSessions []models.ChatSession
	err := configs.DB.
		Joins("JOIN chat_messages ON chat_messages.chat_session_id = chat_sessions.id").
		Where("chat_messages.sender_id = ? OR chat_messages.receiver_id = ?", uid, uid).
		Group("chat_sessions.id").
		Find(&chatSessions).Error

	if err != nil {
		return nil, err
	}

	return chatSessions, nil
}

func CreateChatSession(chatSession *models.ChatSession) error {
	result := configs.DB.Create(
		&chatSession,
	)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("Created chat session: %v", chatSession)
	return nil
}

func UpdateLastMessage(chatSessionID string, messageID string) error {
	result := configs.DB.Exec(
		"UPDATE chat_sessions SET last_message = ? WHERE id = ?",
		messageID, chatSessionID,
	)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
