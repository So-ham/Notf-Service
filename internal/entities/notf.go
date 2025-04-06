package entities

import (
	"time"
)

type FlaggedNotf struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	UserID     uint      `json:"user_id" gorm:"index"`
	Severity   string    `json:"severity"`
	Content    string    `json:"content"`
	NotifiedAt time.Time `json:"notified_at,omitempty"`
	Notified   bool      `json:"notified" gorm:"default:false"`
}

type FlaggedNotfRequest struct {
	UserID   uint   `json:"user_id"`
	Severity string `json:"severity"`
	Content  string `json:"content"`
}
