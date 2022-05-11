package subdistrict

import (
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

type SubdistrictRepositoryInterface interface {
	GetSubdistricts() ([]models.Subdistrict, error)
	GetSubdistrictById(subdistrictId int) (models.Subdistrict, error)
	GetSubdistrictByDistrictId(districtId int) ([]models.Subdistrict, error)
}

type subdistrictRepository struct {
	db *gorm.DB
}

func NewSubdistrictRepository(db *gorm.DB) *subdistrictRepository {
	return &subdistrictRepository{db}
}

func (r *subdistrictRepository) GetSubdistricts() ([]models.Subdistrict, error) {
	var subdistricts []models.Subdistrict

	err := r.db.Order("name asc").Find(&subdistricts).Error

	return subdistricts, err
}

func (r *subdistrictRepository) GetSubdistrictById(subdistrictId int) (models.Subdistrict, error) {
	var subdistrict models.Subdistrict

	err := r.db.Where("id = ?", subdistrictId).First(&subdistrict).Error

	return subdistrict, err
}

func (r *subdistrictRepository) GetSubdistrictByDistrictId(districtId int) ([]models.Subdistrict, error) {
	var subdistricts []models.Subdistrict

	err := r.db.Order("name asc").Where("district_id = ?", districtId).Find(&subdistricts).Error

	return subdistricts, err
}
