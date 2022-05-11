package subdistrict

import (
	"github.com/novaldwp/go-news-api/models"
)

type SubdistrictServiceInterface interface {
	GetSubdistricts(districtId int) ([]models.Subdistrict, error)
	GetSubdistrictById(districtId int) (models.Subdistrict, error)
}

type subdistrictService struct {
	repository SubdistrictRepositoryInterface
}

func NewSubdistrictService(subdistrictRepository SubdistrictRepositoryInterface) *subdistrictService {
	return &subdistrictService{subdistrictRepository}
}

// fetch all data subdistrictS
func (s *subdistrictService) GetSubdistricts(districtId int) ([]models.Subdistrict, error) {
	var subdistricts []models.Subdistrict
	var err error

	if districtId != 0 {
		subdistricts, err = s.repository.GetSubdistrictByDistrictId(districtId)
	} else {
		subdistricts, err = s.repository.GetSubdistricts()
	}

	return subdistricts, err
}

// return data city by field "id"
func (s *subdistrictService) GetSubdistrictById(districtId int) (models.Subdistrict, error) {
	return s.repository.GetSubdistrictById(districtId)
}
