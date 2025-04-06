package services

import (
	"context"

	"github.com/So-ham/notf-service/internal/entities"
	"github.com/So-ham/notf-service/internal/models"
)

// Service represents the service layer having
// all the services from all service packages
type service struct {
	model models.Model
}

// New creates a new instance of Service
func New(model *models.Model) Service {
	m := &service{model: *model}
	return m
}

type Service interface {
	AddFlaggedNotf(ctx context.Context, req *entities.FlaggedNotfRequest) (err error)
	GetNotifications(ctx context.Context) ([]entities.FlaggedNotf, error)
}
