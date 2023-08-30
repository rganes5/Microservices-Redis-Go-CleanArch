//go:build wireinject
// +build wireinject

package di

import (
	"X-TENTIONCREW/api_gateway/pkg/api"
	"X-TENTIONCREW/api_gateway/pkg/api/handlers"
	"X-TENTIONCREW/api_gateway/pkg/client"
	"X-TENTIONCREW/api_gateway/pkg/config"
	"X-TENTIONCREW/api_gateway/pkg/service"

	"github.com/google/wire"
)

func InitializeAPI(c *config.Config) (*api.Server, error) {
	wire.Build(service.InitClient,
		client.NewauthClient, client.NewMethodClient,
		handlers.NewUserHandler, handlers.NewMethodHandler,
		api.NewServerHTTP)
	return &api.Server{}, nil
}
