package models

import "gorm.io/gorm"

type ChatSession struct {
	gorm.Model
	LastMessage string
}
