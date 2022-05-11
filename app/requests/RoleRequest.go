package requests

// config request
type CreateRoleRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateRoleRequest struct {
	Name   string `json:"name" binding:"required"`
	Status *bool  `json:"status" binding:"required"`
}
