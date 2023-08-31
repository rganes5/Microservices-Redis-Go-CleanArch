//go:build wireinject

package di

import (
	"method_svc/pkg/api"
	"method_svc/pkg/api/service"
	"method_svc/pkg/client"
	"method_svc/pkg/config"

	"github.com/google/wire"
)

func InitializeServe(c *config.Config) (*api.Server, error) {
	wire.Build(
		client.InitAuthClient,
		client.NewAuthClient,
		service.NewMethodService,
		api.NewGrpcServe)
	return &api.Server{}, nil
}
