package models

import "gorm.io/gorm"

type ChatMessage struct {
	gorm.Model
	ChatSessionID uint   `gorm:"not null" json:"chat_session_id"`
	SenderID      uint   `gorm:"not null" json:"sender_id"`
	ReceiverID    uint   `gorm:"not null" json:"receiver_id"`
	Content       string `gorm:"not null"`
}
