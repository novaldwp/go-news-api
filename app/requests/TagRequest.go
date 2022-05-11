package requests

// config request
type CreateTagRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTagRequest struct {
	Name   string `json:"name" binding:"required"`
	Status *bool  `json:"status" binding:"required"`
}
