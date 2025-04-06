package grpc

import (
	"github.com/So-ham/notf-service/internal/services"
	pb "github.com/So-ham/notf-service/internal/web/grpc"
)

type notfServiceServer struct {
	pb.UnimplementedNotfServiceServer
	Service services.Service
}

type PoutServiceServer interface {
	pb.UnsafeNotfServiceServer
}

func New(s services.Service) pb.NotfServiceServer {
	return &notfServiceServer{
		Service: s,
	}
}
