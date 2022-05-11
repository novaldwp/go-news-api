package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/helper"
	"github.com/novaldwp/go-news-api/app/requests"
	"github.com/novaldwp/go-news-api/app/responses"
)

type authHandler struct {
	aService AuthServiceInterface
}

func NewAuthHandler(aService AuthServiceInterface) *authHandler {
	return &authHandler{aService}
}

// get all users
func (h *authHandler) Login(c *gin.Context) {
	var input requests.AuthLoginRequest

	// validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helper.ErrorRequestResponse(err)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	// checking credentials from database
	user, err := h.aService.CheckCredentials(input)

	// if had error
	if err != nil {
		// set error response
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	// generate token after success check credentials
	token, err := helper.GenerateJWT(user.Email, user.Role.Name)

	// validate
	if err != nil {
		response := helper.ErrorResponse(err)
		c.JSON(http.StatusInternalServerError, response)

		return
	}

	// set success response
	convertResponse := responses.ToAuthResponse(user, token)
	response := helper.SuccessResponseWithData("users", "get", convertResponse)

	// set return
	c.JSON(http.StatusOK, response)
}
