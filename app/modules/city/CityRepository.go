package city

import (
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

type CityRepositoryInterface interface {
	GetCities() ([]models.City, error)
	GetCityById(cityId int) (models.City, error)
	GetCityByProvinceId(provinceId int) ([]models.City, error)
}

type cityRepository struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) *cityRepository {
	return &cityRepository{db}
}

func (r *cityRepository) GetCities() ([]models.City, error) {
	var cities []models.City

	err := r.db.Order("name asc").Find(&cities).Error

	return cities, err
}

func (r *cityRepository) GetCityById(cityId int) (models.City, error) {
	var city models.City

	err := r.db.Where("id = ?", cityId).First(&city).Error

	return city, err
}

func (r *cityRepository) GetCityByProvinceId(provinceId int) ([]models.City, error) {
	var cities []models.City

	err := r.db.Order("name asc").Where("province_id = ?", provinceId).Find(&cities).Error

	return cities, err
}
