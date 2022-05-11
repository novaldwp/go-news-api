package tag

import (
	"math"

	"github.com/novaldwp/go-news-api/app/helper"
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

type TagRepositoryInterface interface {
	Pagination(pagination *helper.Pagination) *helper.PaginationRepository
	GetTags() ([]models.Tag, error)
	GetActiveTags() ([]models.Tag, error)
	GetNonActiveTags() ([]models.Tag, error)
	GetTagById(tagId int) (models.Tag, error)
	CreateTag(tag models.Tag) error
	UpdateTag(tag models.Tag) error
	DeleteTag(tag models.Tag) error
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *tagRepository {
	return &tagRepository{db}
}

func (r *tagRepository) GetTags() ([]models.Tag, error) {
	var tags []models.Tag

	err := r.db.Order("id desc").Find(&tags).Error

	return tags, err
}

func (r *tagRepository) GetActiveTags() ([]models.Tag, error) {
	var tags []models.Tag

	err := r.db.Order("id desc").Where("status = 1").Find(&tags).Error

	return tags, err
}

func (r *tagRepository) GetNonActiveTags() ([]models.Tag, error) {
	var tags []models.Tag

	err := r.db.Order("id desc").Where("status = 0").Find(&tags).Error

	return tags, err
}

func (r *tagRepository) GetTagById(tagId int) (models.Tag, error) {
	var tag models.Tag

	err := r.db.Where("id = ?", tagId).First(&tag).Error

	return tag, err
}

func (r *tagRepository) CreateTag(tag models.Tag) error {
	err := r.db.Create(&tag).Error

	return err
}

func (r *tagRepository) UpdateTag(tag models.Tag) error {
	err := r.db.Save(&tag).Error

	return err
}

func (r *tagRepository) DeleteTag(tag models.Tag) error {
	err := r.db.Delete(&tag).Error

	return err
}

func (r *tagRepository) Pagination(p *helper.Pagination) *helper.PaginationRepository {
	var tags []models.Category
	var totalRows int64 = 0

	// get data tags
	err := r.db.Order(p.GetSort()).Limit(p.GetLimit()).Offset(p.GetOffset()).Find(&tags).Error
	if err != nil {
		return &helper.PaginationRepository{Error: err}
	}

	// get total rows
	if err := r.db.Model(&tags).Count(&totalRows).Error; err != nil {
		return &helper.PaginationRepository{Error: err}
	}

	// get total page
	totalPages := int(math.Ceil(float64(totalRows) / float64(p.GetLimit())))

	// set pagination with value
	p.Data = tags
	p.TotalRows = int(totalRows)
	p.TotalPages = totalPages

	return &helper.PaginationRepository{Result: p}
}
