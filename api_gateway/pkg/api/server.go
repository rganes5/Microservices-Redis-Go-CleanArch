package api

import (
	_ "X-TENTIONCREW/api_gateway/docs"
	"X-TENTIONCREW/api_gateway/pkg/api/handlers"
	"X-TENTIONCREW/api_gateway/pkg/api/routes"

	"X-TENTIONCREW/api_gateway/pkg/config"
	"net/http"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
	Port   string
}

func NewServerHTTP(c *config.Config, authHandler handlers.AuthHandler, methodHandler handlers.MethodHandler) (*Server, error) {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.RegisterUserRoutes(engine.Group("/"), authHandler, methodHandler)
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
