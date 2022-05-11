package district

import (
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

type DistrictRepositoryInterface interface {
	GetDistricts() ([]models.District, error)
	GetDistrictById(districtId int) (models.District, error)
	GetDistrictByCityId(cityId int) ([]models.District, error)
}

type districtRepository struct {
	db *gorm.DB
}

func NewDistrictRepository(db *gorm.DB) *districtRepository {
	return &districtRepository{db}
}

func (r *districtRepository) GetDistricts() ([]models.District, error) {
	var districts []models.District

	err := r.db.Order("name asc").Find(&districts).Error

	return districts, err
}

func (r *districtRepository) GetDistrictById(districtId int) (models.District, error) {
	var district models.District

	err := r.db.Where("id = ?", districtId).First(&district).Error

	return district, err
}

func (r *districtRepository) GetDistrictByCityId(cityId int) ([]models.District, error) {
	var districts []models.District

	err := r.db.Order("name asc").Where("city_id = ?", cityId).Find(&districts).Error

	return districts, err
}
