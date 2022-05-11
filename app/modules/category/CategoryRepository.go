package category

import (
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

type CategoryRepositoryInterface interface {
	GetCategories() ([]models.Category, error)
	GetActiveCategories() ([]models.Category, error)
	GetNonActiveCategories() ([]models.Category, error)
	GetCategoryById(categoryId int) (models.Category, error)
	CreateCategory(category models.Category) error
	UpdateCategory(category models.Category) error
	DeleteCategory(category models.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategories() ([]models.Category, error) {
	var categories []models.Category

	err := r.db.Order("id desc").Find(&categories).Error

	return categories, err
}

func (r *categoryRepository) GetActiveCategories() ([]models.Category, error) {
	var categories []models.Category

	err := r.db.Order("id desc").Where("status = 1").Find(&categories).Error

	return categories, err
}

func (r *categoryRepository) GetNonActiveCategories() ([]models.Category, error) {
	var categories []models.Category

	err := r.db.Order("id desc").Where("status = 0").Find(&categories).Error

	return categories, err
}

func (r *categoryRepository) GetCategoryById(categoryId int) (models.Category, error) {
	var category models.Category

	err := r.db.Where("id = ?", categoryId).First(&category).Error

	return category, err
}

func (r *categoryRepository) CreateCategory(category models.Category) error {
	err := r.db.Create(&category).Error

	return err
}

func (r *categoryRepository) UpdateCategory(category models.Category) error {
	err := r.db.Save(&category).Error

	return err
}

func (r *categoryRepository) DeleteCategory(category models.Category) error {
	err := r.db.Delete(&category).Error

	return err
}
