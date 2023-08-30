package handlers

import (
	client "X-TENTIONCREW/api_gateway/pkg/client/interfaces"
	"X-TENTIONCREW/api_gateway/pkg/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Client client.AuthClient
}

func NewUserHandler(client client.AuthClient) AuthHandler {
	return AuthHandler{
		Client: client,
	}
}

func (cr *AuthHandler) Register(c *gin.Context) {
	var body utils.SignUpBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind" + err.Error(),
		})
		return
	}
	res, err := cr.Client.Register(context.Background(), body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "commnuication error" + err.Error(),
		})
		return
	}
	c.JSON(int(res.Status), gin.H{
		"responseid": &res.Response,
	})
}
