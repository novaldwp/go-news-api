package responses

import "github.com/novaldwp/go-news-api/models"

type CategoryResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Status bool   `json:"status"`
}

func ToCategoryResponse(category models.Category) CategoryResponse {
	return CategoryResponse{
		ID:     int(category.ID),
		Name:   category.Name,
		Slug:   category.Slug,
		Status: *category.Status,
	}
}

func ToCategoryResponses(categories []models.Category) []CategoryResponse {
	var categoryResponses []CategoryResponse

	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}
