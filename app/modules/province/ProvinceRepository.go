package province

import (
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

type ProvinceRepositoryInterface interface {
	GetProvinces() ([]models.Province, error)
	GetProvinceById(provinceId int) (models.Province, error)
}

type provinceRepository struct {
	db *gorm.DB
}

func NewProvinceRepository(db *gorm.DB) *provinceRepository {
	return &provinceRepository{db}
}

func (r *provinceRepository) GetProvinces() ([]models.Province, error) {
	var provinces []models.Province

	err := r.db.Order("name asc").Find(&provinces).Error

	return provinces, err
}

func (r *provinceRepository) GetProvinceById(provinceId int) (models.Province, error) {
	var province models.Province

	err := r.db.Where("id = ?", provinceId).First(&province).Error

	return province, err
}
