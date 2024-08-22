package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"gophKeeper/internal/config"
	"gophKeeper/internal/server/grpc/interceptors"
	"net"
)

type GRPCServer struct {
	Srv      *grpc.Server
	cfg      config.ConfigServer
	services []func(*grpc.Server)
}

func NewGRPCServer(cfg config.ConfigServer, srvOptions ...grpc.ServerOption) *GRPCServer {
	var srv *grpc.Server
	if len(srvOptions) > 0 {
		srv = grpc.NewServer(srvOptions...)
	} else {
		srv = grpc.NewServer(grpc.UnaryInterceptor(interceptors.AuthInterceptor))
	}

	return &GRPCServer{
		Srv: srv,
		cfg: cfg,
	}
}

func (s *GRPCServer) RegisterService(registerFunc func(*grpc.Server)) {
	s.services = append(s.services, registerFunc)
}

func (s *GRPCServer) Start() error {
	for _, registerFunc := range s.services {
		registerFunc(s.Srv)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s", s.cfg.GetGrpcAddress()))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	return s.Srv.Serve(lis)
}

func (s *GRPCServer) Stop() error {
	s.Srv.GracefulStop()
	return nil
}
