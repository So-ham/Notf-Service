package Ngrpc

import (
	"context"

	"github.com/So-ham/notf-service/internal/entities"
	pb "github.com/So-ham/notf-service/internal/web/grpc"

	"google.golang.org/protobuf/types/known/emptypb"
)

// GetDepartmentCounts returns department count for that org
func (h *notfServiceServer) SendFlaggedNotification(ctx context.Context, req *pb.SendFlaggedNotificationReq) (*emptypb.Empty, error) {
	reqb := &entities.FlaggedNotfRequest{
		UserID:   uint(req.UserID),
		Severity: req.Severity,
		Content:  req.Content,
	}
	err := h.Service.AddFlaggedNotf(ctx, reqb)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
