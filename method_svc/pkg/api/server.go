package api

import (
	"fmt"
	"method_svc/pkg/config"
	"method_svc/pkg/pb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	Gs   *grpc.Server
	Lis  net.Listener
	Port string
}

func NewGrpcServe(c *config.Config, service pb.MethodServiceServer) (*Server, error) {
	grpcServer := grpc.NewServer()
	pb.RegisterMethodServiceServer(grpcServer, service)
	reflection.Register(grpcServer)

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
	fmt.Println("Method service on:", s.Port)
	return s.Gs.Serve(s.Lis)
}
