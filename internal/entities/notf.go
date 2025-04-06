package entities

import "gorm.io/gorm"

type FlaggedNotf struct {
	gorm.Model
	UserID   uint
	Severity string `json:"severity;omitempty"`
	Content  string
}
