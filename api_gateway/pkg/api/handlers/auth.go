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

// USER SIGN-UP
//
//	@Summary		API FOR NEW USER SIGN UP
//	@ID				SIGNUP-USER
//	@Description	CREATE A NEW USER WITH REQUIRED DETAILS
//	@Tags			USER
//	@Accept			json
//	@Produce		json
//	@Param			user_details	body		utils.SignUpBody true "Enter the user details"
//	@Success		200				{object}	utils.Response
//	@Failure		400				{object}	utils.Response
//	@Failure		500				{object}	utils.Response
//	@Router			/user/register [post]
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

// LIST USER BASED ON CATEGORY
// @Summary API FOR LISTING USER BASED ON ID
// @ID LIST-USER
// @Description LISTING USER BASED ON ID
// @Tags USER
// @Accept json
// @Produce json
// @Param user_id path string true "Enter the user id"
// @Success 200 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /user/getuser/{user_id} [get]
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

// UPDATE USER
// @Summary API FOR UPDATING USER
// @ID UPDATE-USER
// @Description UPDATING USER DETAILS WITH ID
// @Tags USER
// @Accept json
// @Produce json
// @Param user_details body utils.UpdateBody true "Enter the user details"
// @Success 200 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /user/update [patch]
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

// DELETE USER
// @Summary API FOR DELETING A USER
// @ID DELETE-USER
// @Description DELETING A USER BASED ON ID
// @Tags USER
// @Accept json
// @Produce json
// @Param user_id path string true "Enter the user id that you would like to delete"
// @Success 200 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /user/delete/{user_id} [delete]
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
