package handlers

import (
	client "X-TENTIONCREW/api_gateway/pkg/client/interfaces"
	"X-TENTIONCREW/api_gateway/pkg/utils"
	"context"
	"net/http"
	"strconv"

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
		"response": &res.Response,
	})
}

func (cr *AuthHandler) GetUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("user_id"), 10, 32)
	userId := uint32(id)
	res, err := cr.Client.GetUser(context.Background(), userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "communication error" + err.Error(),
		})
		return
	}

	c.JSON(int(res.Status), gin.H{
		"User": &res.User,
	})
}

func (cr *AuthHandler) UpdateUser(c *gin.Context) {
	var body utils.UpdateBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "failed to bind" + err.Error(),
		})
		return
	}

	res, err := cr.Client.UpdateUser(context.Background(), body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "communication error" + err.Error(),
		})
		return
	}

	c.JSON(int(res.Status), gin.H{
		"Response": &res.Response,
		"User":     &res.User,
	})

}

func (cr *AuthHandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("user_id"), 10, 32)
	userId := uint32(id)
	res, err := cr.Client.DeleteUser(context.Background(), userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "communication error" + err.Error(),
		})
		return
	}

	c.JSON(int(res.Status), gin.H{
		"Response": &res.Response,
	})
}
