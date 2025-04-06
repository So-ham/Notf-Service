package models

import (
	"github.com/So-ham/notf-service/internal/models/flaggedNotf"
	"gorm.io/gorm"
)

type Model struct {
	FlaggedNotf flaggedNotf.FlaggedNotf
}

// New creates a new instance of Model
func New(gdb *gorm.DB) *Model {
	return &Model{
		FlaggedNotf: flaggedNotf.New(gdb),
	}
}
