package district

import (
	"github.com/novaldwp/go-news-api/models"
)

type DistrictServiceInterface interface {
	GetDistricts(cityId int) ([]models.District, error)
	GetDistrictById(districtId int) (models.District, error)
}

type districtService struct {
	repository DistrictRepositoryInterface
}

func NewDistrictService(districtRepository DistrictRepositoryInterface) *districtService {
	return &districtService{districtRepository}
}

// fetch all data districtS
func (s *districtService) GetDistricts(cityId int) ([]models.District, error) {
	var districts []models.District
	var err error

	if cityId != 0 {
		districts, err = s.repository.GetDistrictByCityId(cityId)
	} else {
		districts, err = s.repository.GetDistricts()
	}

	return districts, err
}

// return data city by field "id"
func (s *districtService) GetDistrictById(districtId int) (models.District, error) {
	return s.repository.GetDistrictById(districtId)
}
