package routes

import (
	"api_gateway/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(api *gin.RouterGroup, authHandler handlers.AuthHandler, methodHandler handlers.MethodHandler) {
	register := api.Group("/user")
	{
		register.POST("/register", authHandler.Register)

	}
	operations := api.Group("/user")
	{
		operations.GET("/getuser/:user_id", authHandler.GetUser)
		operations.PATCH("/update", authHandler.UpdateUser)
		operations.DELETE("/delete/:user_id", authHandler.DeleteUser)
	}
	methods := api.Group("/user")
	{
		methods.POST("/method", methodHandler.MethodsHandler)
	}
}
