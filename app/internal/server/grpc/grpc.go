package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"gophKeeper/internal/config"
)

type GRPCServer struct {
	srv *grpc.Server
	cfg config.ConfigServer
}

func NewGRPCServer(cfg config.ConfigServer) *GRPCServer {
	return &GRPCServer{
		srv: grpc.NewServer(),
		cfg: cfg,
	}
}

func (s *GRPCServer) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s", s.cfg.GetGrpcAddress()))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// зарегистрировать gRPC сервисы

	return s.srv.Serve(lis)
}

func (s *GRPCServer) Stop() error {
	s.srv.GracefulStop()
	return nil
}
