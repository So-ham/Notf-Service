package services

import "github.com/So-ham/notf-service/internal/models"

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
}
