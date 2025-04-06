package Ngrpc

import (
	"context"

	"github.com/So-ham/notf-service/internal/services"
	pb "github.com/So-ham/notf-service/internal/web/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type notfServiceServer struct {
	pb.UnimplementedNotfServiceServer
	Service services.Service
}

type NotfServiceServer interface {
	SendFlaggedNotification(ctx context.Context, req *pb.SendFlaggedNotificationReq) (*emptypb.Empty, error)
	pb.UnsafeNotfServiceServer
}

func New(s services.Service) pb.NotfServiceServer {
	return &notfServiceServer{
		Service: s,
	}
}
