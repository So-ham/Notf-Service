package services

import (
	"context"

	"github.com/So-ham/notf-service/internal/entities"
	"github.com/So-ham/notf-service/pkg/middlewares"
)

func (s *service) GetNotifications(ctx context.Context) ([]entities.FlaggedNotf, error) {
	curUser := middlewares.GetUserContext(ctx)
	notfs, err := s.model.FlaggedNotf.GetNotifications(ctx, curUser.UserID)
	if err != nil {
		return nil, err
	}
	//

	return notfs, nil
}
