package responses

import "github.com/novaldwp/go-news-api/models"

type TagResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Status bool   `json:"status"`
}

func ToTagResponse(tag models.Tag) TagResponse {
	return TagResponse{
		ID:     int(tag.ID),
		Name:   tag.Name,
		Slug:   tag.Slug,
		Status: *tag.Status,
	}
}

func ToTagResponses(tags []models.Tag) []TagResponse {
	var tagResponses []TagResponse

	for _, tag := range tags {
		tagResponses = append(tagResponses, ToTagResponse(tag))
	}

	return tagResponses
}
