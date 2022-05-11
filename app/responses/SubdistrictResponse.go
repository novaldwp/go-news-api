package responses

import "github.com/novaldwp/go-news-api/models"

type SubdistrictResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ToSubdistrictResponse(subdistrict models.Subdistrict) SubdistrictResponse {
	return SubdistrictResponse{
		ID:   int(subdistrict.ID),
		Name: subdistrict.Name,
	}
}

func ToSubdistrictResponses(subdistricts []models.Subdistrict) []SubdistrictResponse {
	var subdistrictResponses []SubdistrictResponse

	for _, subdistrict := range subdistricts {
		subdistrictResponses = append(subdistrictResponses, ToSubdistrictResponse(subdistrict))
	}

	return subdistrictResponses
}
