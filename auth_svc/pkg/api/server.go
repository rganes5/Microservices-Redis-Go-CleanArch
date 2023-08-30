package api

import (
	"X-TENTIONCREW/auth_svc/pkg/config"
	"X-TENTIONCREW/auth_svc/pkg/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	Gs   *grpc.Server
	Lis  net.Listener
	Port string
}

func NewGrpcServe(c *config.Config, service pb.AuthServiceServer) (*Server, error) {
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, service)
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		return nil, err
	}
	return &Server{
		Gs:   grpcServer,
		Lis:  lis,
		Port: c.Port,
	}, nil
}

func (s *Server) Start() error {
	fmt.Println("Authentication service on:", s.Port)
	return s.Gs.Serve(s.Lis)
}
