//go:build wireinject
// +build wireinject

package di

import (
	"api_gateway/pkg/api"
	"api_gateway/pkg/api/handlers"
	"api_gateway/pkg/client"
	"api_gateway/pkg/config"
	"api_gateway/pkg/service"

	"github.com/google/wire"
)

func InitializeAPI(c *config.Config) (*api.Server, error) {
	wire.Build(service.InitClient,
		client.NewauthClient, client.NewMethodClient,
		handlers.NewUserHandler, handlers.NewMethodHandler,
		api.NewServerHTTP)
	return &api.Server{}, nil
}
