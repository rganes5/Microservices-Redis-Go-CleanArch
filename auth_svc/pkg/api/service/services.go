package service

import (
	"X-TENTIONCREW/auth_svc/pkg/pb"
	"X-TENTIONCREW/auth_svc/pkg/repository/interfaces"
)

type authService struct {
	Repo interfaces.AuthRepo
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(repo interfaces.AuthRepo) pb.AuthServiceServer {
	return &authService{
		Repo: repo,
	}
}
