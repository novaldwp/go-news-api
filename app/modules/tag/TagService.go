package tag

import (
	"github.com/gosimple/slug"
	"github.com/novaldwp/go-news-api/app/helper"
	"github.com/novaldwp/go-news-api/app/requests"
	"github.com/novaldwp/go-news-api/models"
)

type TagServiceInterface interface {
	Pagination(pagination *helper.Pagination, query *helper.PaginationQuery, link *helper.PaginationLink, urlPath string) ([]models.Tag, error)
	GetTags(status string) ([]models.Tag, error)
	GetTagById(tagId int) (models.Tag, error)
	CreateTag(tagRequest requests.CreateTagRequest) error
	UpdateTag(tagId int, tagRequest requests.UpdateTagRequest) error
	UpdateTagStatus(tagId int) error
	DeleteTag(tagId int) error
}

type tagService struct {
	repository TagRepositoryInterface
}

func NewTagService(tagRepository TagRepositoryInterface) *tagService {
	return &tagService{tagRepository}
}

// fetch all tags with parameter status. If status true, then will fetch active tag
// which have value 1 on field "status", if status false, then will fetch non active tag
// which have value 0 on field "status", if not both status true/false then will fetch all tag
func (s *tagService) GetTags(status string) ([]models.Tag, error) {
	var tags []models.Tag
	var errors error

	switch status {
	case "true":
		tags, errors = s.repository.GetActiveTags()
	case "false":
		tags, errors = s.repository.GetNonActiveTags()
	default:
		tags, errors = s.repository.GetTags()
	}

	return tags, errors
}

// return data tag by field "id"
func (s *tagService) GetTagById(tagId int) (models.Tag, error) {
	return s.repository.GetTagById(tagId)
}

// create new tag
func (s *tagService) CreateTag(tagRequest requests.CreateTagRequest) error {
	tag := models.Tag{
		Name: tagRequest.Name,
		Slug: slug.Make(tagRequest.Name),
	}

	err := s.repository.CreateTag(tag)

	return err
}

// update selected tag
func (s *tagService) UpdateTag(tagId int, tagRequest requests.UpdateTagRequest) error {
	tag, _ := s.repository.GetTagById(tagId)

	tag.Name = tagRequest.Name
	tag.Slug = slug.Make(tagRequest.Name)
	tag.Status = tagRequest.Status

	err := s.repository.UpdateTag(tag)

	return err
}

func (s *tagService) UpdateTagStatus(tagId int) error {
	tag, _ := s.repository.GetTagById(tagId)

	// if tag status == false
	if !*tag.Status {
		// set into true
		*tag.Status = true
	} else {
		// set into false
		*tag.Status = false
	}

	err := s.repository.UpdateTag(tag)

	return err
}

// delete selected tag
func (s *tagService) DeleteTag(tagId int) error {
	tag, _ := s.repository.GetTagById(tagId)

	err := s.repository.DeleteTag(tag)

	return err
}

func (s *tagService) Pagination(pagination *helper.Pagination, query *helper.PaginationQuery, link *helper.PaginationLink, appPATH string) ([]models.Tag, error) {
	result, err := s.repository.Pagination(pagination, query)

	if err != nil {
		return nil, err
	}

	// set pagination link
	helper.GeneratePaginationLink(appPATH, pagination, query, link)

	return result, nil
}
