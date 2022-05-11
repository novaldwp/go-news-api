package auth

import (
	"fmt"

	"github.com/novaldwp/go-news-api/app/modules/user"
	"github.com/novaldwp/go-news-api/app/requests"
	"github.com/novaldwp/go-news-api/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	CheckCredentials(loginRequest requests.AuthLoginRequest) (models.User, error)
	// Register(userRequest requests.CreateUserRequest) (interface{}, error)
}

type authService struct {
	repository user.UserRepositoryInterface
}

func NewAuthService(userRepository user.UserRepositoryInterface) *authService {
	return &authService{userRepository}
}

func (s *authService) CheckCredentials(loginRequest requests.AuthLoginRequest) (models.User, error) {
	var user models.User
	var err error

	// validate email
	user, err = s.repository.GetUserByEmail(loginRequest.Email)

	if err != nil {
		err = fmt.Errorf("email not registered, please try again")
	}

	// validate password
	checkHash := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))

	if checkHash != nil {
		err = fmt.Errorf("invalid login credentials, please try again")
	}

	return user, err
}
