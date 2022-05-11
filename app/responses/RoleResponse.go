package responses

import "github.com/novaldwp/go-news-api/models"

type RoleResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

func ToRoleResponse(role models.Role) RoleResponse {
	return RoleResponse{
		ID:     int(role.ID),
		Name:   role.Name,
		Status: *role.Status,
	}
}

func ToRoleResponses(roles []models.Role) []RoleResponse {
	var roleResponses []RoleResponse

	for _, role := range roles {
		roleResponses = append(roleResponses, ToRoleResponse(role))
	}

	return roleResponses
}
