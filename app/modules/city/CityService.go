package city

import (
	"github.com/novaldwp/go-news-api/models"
)

type CityServiceInterface interface {
	GetCities(provinceId int) ([]models.City, error)
	GetCityById(cityId int) (models.City, error)
}

type cityService struct {
	repository CityRepositoryInterface
}

func NewCityService(cityRepository CityRepositoryInterface) *cityService {
	return &cityService{cityRepository}
}

// fetch all data citys
func (s *cityService) GetCities(provinceId int) ([]models.City, error) {
	var cities []models.City
	var err error

	if provinceId != 0 {
		cities, err = s.repository.GetCityByProvinceId(provinceId)
	} else {
		cities, err = s.repository.GetCities()
	}

	return cities, err
}

// return data city by field "id"
func (s *cityService) GetCityById(cityId int) (models.City, error) {
	return s.repository.GetCityById(cityId)
}
