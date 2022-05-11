package responses

import "github.com/novaldwp/go-news-api/models"

type CityResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ToCityResponse(city models.City) CityResponse {
	return CityResponse{
		ID:   int(city.ID),
		Name: city.Name,
	}
}

func ToCityResponses(cities []models.City) []CityResponse {
	var cityResponses []CityResponse

	for _, city := range cities {
		cityResponses = append(cityResponses, ToCityResponse(city))
	}

	return cityResponses
}
