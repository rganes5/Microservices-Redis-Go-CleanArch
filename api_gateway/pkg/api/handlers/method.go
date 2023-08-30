package handlers

import (
	client "X-TENTIONCREW/api_gateway/pkg/client/interfaces"
	"X-TENTIONCREW/api_gateway/pkg/utils"

	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MethodHandler struct {
	Client client.MethodClient
}

func NewMethodHandler(client client.MethodClient) MethodHandler {
	return MethodHandler{
		Client: client,
	}
}

// METHODS
// @Summary API FOR METHOD DEMONSTRATION USING CONCURRENCY AND PARALLELISM
// @ID METHODS
// @Description API FOR METHOD DEMONSTRATION USING CONCURRENCY AND PARALLELISM
// @Tags METHOD
// @Accept json
// @Produce json
// @Param MethodsRequest body utils.MethodsRequest true "Enter the method number and waittime"
// @Success 200 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /user/method [post]
func (cr *MethodHandler) MethodsHandler(c *gin.Context) {
	var request utils.MethodsRequest
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid request format" + err.Error(),
		})
		return
	}

	res, err := cr.Client.MethodService(context.Background(), request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "communication error" + err.Error(),
		})
		return
	}

	c.JSON(int(res.Status), gin.H{
		"Response": &res.Response,
		"Count":    &res.Count,
		"User":     &res.Users,
	})

}
