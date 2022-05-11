package province

import (
	"github.com/novaldwp/go-news-api/models"
)

type ProvinceServiceInterface interface {
	GetProvinces() ([]models.Province, error)
	GetProvinceById(provinceId int) (models.Province, error)
}

type provinceService struct {
	repository ProvinceRepositoryInterface
}

func NewProvinceService(provinceRepository ProvinceRepositoryInterface) *provinceService {
	return &provinceService{provinceRepository}
}

// fetch all data provinces
func (s *provinceService) GetProvinces() ([]models.Province, error) {
	provinces, errors := s.repository.GetProvinces()

	return provinces, errors
}

// return data province by field "id"
func (s *provinceService) GetProvinceById(provinceId int) (models.Province, error) {
	return s.repository.GetProvinceById(provinceId)
}
