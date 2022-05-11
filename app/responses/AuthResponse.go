package responses

import (
	"fmt"
	"strings"

	"github.com/novaldwp/go-news-api/models"
)

type LoginResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Token     string `json:"token"`
}

func ToAuthResponse(user models.User, token string) LoginResponse {
	return LoginResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		FullName:  fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		Email:     user.Email,
		Role:      strings.ToLower(user.Role.Name),
		Token:     token,
	}
}
