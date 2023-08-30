//go:build wireinject

package di

import (
	"X-TENTIONCREW/auth_svc/pkg/api"
	"X-TENTIONCREW/auth_svc/pkg/api/service"
	"X-TENTIONCREW/auth_svc/pkg/config"
	"X-TENTIONCREW/auth_svc/pkg/db"
	"X-TENTIONCREW/auth_svc/pkg/repository"

	"github.com/google/wire"
)

func InitializeServe(c *config.Config) (*api.Server, error) {
	wire.Build(db.Initdb,
		repository.NewauthRepo,
		service.NewAuthService,
		api.NewGrpcServe)
	return &api.Server{}, nil
}
