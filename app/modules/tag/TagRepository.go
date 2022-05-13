package tag

import (
	"fmt"
	"math"

	"github.com/novaldwp/go-news-api/app/helper"
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

type TagRepositoryInterface interface {
	Pagination(pagination *helper.Pagination, query *helper.PaginationQuery) ([]models.Tag, error)
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

func (r *tagRepository) Pagination(p *helper.Pagination, q *helper.PaginationQuery) ([]models.Tag, error) {
	var tags []models.Tag
	var totalRows int64 = 0

	// get data tags
	if err := r.db.Order(fmt.Sprintf("%s %s", q.Order, q.Sort)).Limit(q.GetLimit()).Offset(q.GetOffset()).Find(&tags).Error; err != nil {
		return nil, err
	}

	// get total rows
	if err := r.db.Model(&tags).Count(&totalRows).Error; err != nil {
		return nil, err
	}

	// get total page
	totalPages := int(math.Ceil(float64(totalRows) / float64(q.GetLimit())))

	// set pagination value
	p.TotalRows = int(totalRows)
	p.TotalPages = totalPages

	return tags, nil
}
