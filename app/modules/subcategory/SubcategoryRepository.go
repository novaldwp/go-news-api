package subcategory

import (
	"fmt"
	"math"

	"github.com/novaldwp/go-news-api/app/helper"
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

type SubcategoryRepositoryInterface interface {
	Pagination(pagination *helper.Pagination, query *helper.PaginationQuery) ([]models.Subcategory, error)
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

func (r *subcategoryRepository) Pagination(p *helper.Pagination, q *helper.PaginationQuery) ([]models.Subcategory, error) {
	var subcategory []models.Subcategory
	var totalRows int64 = 0

	// get data tags
	if err := r.db.Order(fmt.Sprintf("%s %s", q.Order, q.Sort)).Limit(q.GetLimit()).Offset(q.GetOffset()).Find(&subcategory).Error; err != nil {
		return nil, err
	}

	// get total rows
	if err := r.db.Model(&subcategory).Count(&totalRows).Error; err != nil {
		return nil, err
	}

	// get total page
	totalPages := int(math.Ceil(float64(totalRows) / float64(q.GetLimit())))

	// set pagination value
	p.TotalRows = int(totalRows)
	p.TotalPages = totalPages

	return subcategory, nil
}
