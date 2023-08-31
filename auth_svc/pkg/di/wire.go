//go:build wireinject

package di

import (
	"auth_svc/pkg/api"
	"auth_svc/pkg/api/service"
	"auth_svc/pkg/config"
	"auth_svc/pkg/db"
	"auth_svc/pkg/repository"

	"github.com/google/wire"
)

func InitializeServe(c *config.Config) (*api.Server, error) {
	wire.Build(db.Initdb,
		db.InitRedis,
		repository.NewauthRepo,
		service.NewAuthService,
		api.NewGrpcServe)
	return &api.Server{}, nil
}
