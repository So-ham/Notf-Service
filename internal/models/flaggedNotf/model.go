package flaggedNotf

import (
	"context"

	"github.com/So-ham/notf-service/internal/entities"
	"gorm.io/gorm"
)

type flaggedNotf struct {
	DB *gorm.DB
}

type FlaggedNotf interface {
	AddFlaggedNotf(ctx context.Context, req *entities.FlaggedNotfRequest) (err error)
	GetNotifications(ctx context.Context, userID uint) (notifications []entities.FlaggedNotf, err error)
}

func New(db *gorm.DB) FlaggedNotf {
	return &flaggedNotf{DB: db}
}
