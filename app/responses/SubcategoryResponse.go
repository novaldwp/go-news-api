package responses

import "github.com/novaldwp/go-news-api/models"

type SubcategoryResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Status bool   `json:"status"`
}

func ToSubcategoryResponse(subcategory models.Subcategory) SubcategoryResponse {
	return SubcategoryResponse{
		ID:     int(subcategory.ID),
		Name:   subcategory.Name,
		Slug:   subcategory.Slug,
		Status: *subcategory.Status,
	}
}

func ToSubcategoryResponses(subcategories []models.Subcategory) []SubcategoryResponse {
	var subcategoryResponses []SubcategoryResponse

	for _, subcategory := range subcategories {
		subcategoryResponses = append(subcategoryResponses, ToSubcategoryResponse(subcategory))
	}

	return subcategoryResponses
}
