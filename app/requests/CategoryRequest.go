package requests

// config request
type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategoryRequest struct {
	Name   string `json:"name" binding:"required"`
	Status *bool  `json:"status" binding:"required"`
}
