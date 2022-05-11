package subcategory

import (
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

type SubcategoryRepositoryInterface interface {
	GetSubcategories() ([]models.Subcategory, error)
	GetActiveSubcategories() ([]models.Subcategory, error)
	GetNonActiveSubcategories() ([]models.Subcategory, error)
	GetSubcategoryById(subcategoryId int) (models.Subcategory, error)
	CreateSubcategory(subcategory models.Subcategory) error
	UpdateSubcategory(subcategory models.Subcategory) error
	DeleteSubcategory(subcategory models.Subcategory) error
}

type subcategoryRepository struct {
	db *gorm.DB
}

func NewSubcategoryRepository(db *gorm.DB) *subcategoryRepository {
	return &subcategoryRepository{db}
}

func (r *subcategoryRepository) GetSubcategories() ([]models.Subcategory, error) {
	var subcategories []models.Subcategory

	err := r.db.Order("id desc").Find(&subcategories).Error

	return subcategories, err
}

func (r *subcategoryRepository) GetActiveSubcategories() ([]models.Subcategory, error) {
	var subcategories []models.Subcategory

	err := r.db.Order("id desc").Where("status = 1").Find(&subcategories).Error

	return subcategories, err
}

func (r *subcategoryRepository) GetNonActiveSubcategories() ([]models.Subcategory, error) {
	var subcategories []models.Subcategory

	err := r.db.Order("id desc").Where("status = 0").Find(&subcategories).Error

	return subcategories, err
}

func (r *subcategoryRepository) GetSubcategoryById(subcategoryId int) (models.Subcategory, error) {
	var subcategory models.Subcategory

	err := r.db.Where("id = ?", subcategoryId).First(&subcategory).Error

	return subcategory, err
}

func (r *subcategoryRepository) CreateSubcategory(subcategory models.Subcategory) error {
	err := r.db.Create(&subcategory).Error

	return err
}

func (r *subcategoryRepository) UpdateSubcategory(subcategory models.Subcategory) error {
	err := r.db.Save(&subcategory).Error

	return err
}

func (r *subcategoryRepository) DeleteSubcategory(subcategory models.Subcategory) error {
	err := r.db.Delete(&subcategory).Error

	return err
}
