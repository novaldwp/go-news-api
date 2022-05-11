package requests

// config request
type CreateSubcategoryRequest struct {
	Name       string `json:"name" binding:"required"`
	CategoryId uint   `json:"category_id" binding:"required"`
}

type UpdateSubcategoryRequest struct {
	Name       string `json:"name" binding:"required"`
	CategoryId uint   `json:"category_id" binding:"required"`
	Status     *bool  `json:"status" binding:"required"`
}
