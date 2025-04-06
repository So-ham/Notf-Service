package flaggedNotf

import (
	"gorm.io/gorm"
)

type flaggedNotf struct {
	DB *gorm.DB
}

type FlaggedNotf interface {
	AddFlaggedNotf(userID uint, severity, content string) (err error)
	UpdateFlagStatus(id uint, isFlagged bool) error
}

func New(db *gorm.DB) FlaggedNotf {
	return &flaggedNotf{DB: db}
}
