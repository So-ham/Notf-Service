package handlers

import (
	v1 "github.com/So-ham/notf-service/internal/handlers/v1"
	"github.com/So-ham/notf-service/internal/services"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	V1 v1.HandlerV1
}

// New creates a new instance of handlers
func New(service services.Service, validate *validator.Validate) *Handler {
	return &Handler{
		V1: v1.New(service, validate),
	}

}
