package api

import (
	"X-TENTIONCREW/api_gateway/pkg/api/handlers"
	"X-TENTIONCREW/api_gateway/pkg/api/routes"
	"X-TENTIONCREW/api_gateway/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
	Port   string
}

func NewServerHTTP(c *config.Config, authHandler handlers.AuthHandler) (*Server, error) {
	engine := gin.New()
	engine.Use(gin.Logger())
	routes.RegisterUserRoutes(engine.Group("/"), authHandler)
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"statuscode": 404,
			"message":    "invalid url",
		})
	})
	return &Server{
		Engine: engine,
		Port:   c.Port,
	}, nil
}

func (c *Server) Start() {
	c.Engine.Run(c.Port)
}
