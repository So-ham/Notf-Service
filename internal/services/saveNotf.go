package services

import (
	"context"

	"github.com/So-ham/notf-service/internal/entities"
)

func (s *service) AddFlaggedNotf(ctx context.Context, req *entities.FlaggedNotfRequest) (err error) {

	err = s.model.FlaggedNotf.AddFlaggedNotf(ctx, req)
	if err != nil {
		return err
	}
	return nil

}
