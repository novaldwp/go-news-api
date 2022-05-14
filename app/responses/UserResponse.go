package responses

import "github.com/novaldwp/go-news-api/models"

type UserResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	IsVerif   *bool  `json:"is_verif"` // using type data pointer, so you can update value with "false / 0"
	Status    *bool  `json:"status"`   // using type data pointer, so you can update value with "false / 0"
}

func ToUserResponse(user models.User) UserResponse {
	return UserResponse{
		ID:        int(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		IsVerif:   user.IsVerif,
		Status:    user.Status,
	}
}

func ToUserResponses(users []models.User) []UserResponse {
	var userResponses []UserResponse

	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}
