package responses

import "github.com/novaldwp/go-news-api/models"

type ProvinceResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ToProvinceResponse(province models.Province) ProvinceResponse {
	return ProvinceResponse{
		ID:   int(province.ID),
		Name: province.Name,
	}
}

func ToProvinceResponses(provinces []models.Province) []ProvinceResponse {
	var provinceResponses []ProvinceResponse

	for _, province := range provinces {
		provinceResponses = append(provinceResponses, ToProvinceResponse(province))
	}

	return provinceResponses
}
