package responses

import "github.com/novaldwp/go-news-api/models"

type DistrictResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ToDistrictResponse(district models.District) DistrictResponse {
	return DistrictResponse{
		ID:   int(district.ID),
		Name: district.Name,
	}
}

func ToDistrictResponses(districts []models.District) []DistrictResponse {
	var districtResponses []DistrictResponse

	for _, district := range districts {
		districtResponses = append(districtResponses, ToDistrictResponse(district))
	}

	return districtResponses
}
