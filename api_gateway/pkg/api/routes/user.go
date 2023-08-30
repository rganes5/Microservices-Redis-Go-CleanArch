package routes

import (
	"X-TENTIONCREW/api_gateway/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(api *gin.RouterGroup, authHandler handlers.AuthHandler) {
	register := api.Group("/user")
	{
		register.POST("/register", authHandler.Register)
	}
}
