//go:build wireinject

package di

import (
	"X-TENTIONCREW/method_svc/pkg/api"
	"X-TENTIONCREW/method_svc/pkg/api/service"
	"X-TENTIONCREW/method_svc/pkg/client"
	"X-TENTIONCREW/method_svc/pkg/config"

	"github.com/google/wire"
)

func InitializeServe(c *config.Config) (*api.Server, error) {
	wire.Build(service.NewMethodService,
		client.NewAuthClient,
		api.NewGrpcServe)
	return &api.Server{}, nil
}
