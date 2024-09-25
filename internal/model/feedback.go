package model

import (
	"gorm.io/gorm"
	"time"
)

type Feedback struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Title      string    `json:"title"`
	Time       time.Time `json:"time"`
	Category   string    `json:"category"`
	IsUrgent   bool      `json:"is_urgent"`
	Name       string    `json:"name"`
	Content    string    `json:"content"`
	Images     string    `json:"images"`
	Reply      string    `json:"reply"`
	Evaluation string    `json:"evaluation"`
	DeletedAt  gorm.DeletedAt
}
