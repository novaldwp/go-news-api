package category

import (
	"github.com/gosimple/slug"
	"github.com/novaldwp/go-news-api/app/helper"
	"github.com/novaldwp/go-news-api/app/requests"
	"github.com/novaldwp/go-news-api/models"
)

type CategoryServiceInterface interface {
	Pagination(pagination *helper.Pagination, query *helper.PaginationQuery, link *helper.PaginationLink, urlPath string) ([]models.Category, error)
	GetCategories(status string) ([]models.Category, error)
	GetCategoryById(categoryId int) (models.Category, error)
	CreateCategory(categoryRequest requests.CreateCategoryRequest) error
	UpdateCategory(categoryId int, categoryRequest requests.UpdateCategoryRequest) error
	UpdateCategoryStatus(categoryId int) error
	DeleteCategory(categoryId int) error
}

type categoryService struct {
	repository CategoryRepositoryInterface
}

func NewCategoryService(categoryRepository CategoryRepositoryInterface) *categoryService {
	return &categoryService{categoryRepository}
}

// fetch all categories with parameter status. If status true, then will fetch active category
// which have value 1 on field "status", if status false, then will fetch non active category
// which have value 0 on field "status", if not both status true/false then will fetch all category
func (s *categoryService) GetCategories(status string) ([]models.Category, error) {
	var categories []models.Category
	var errors error

	switch status {
	case "true":
		categories, errors = s.repository.GetActiveCategories()
	case "false":
		categories, errors = s.repository.GetNonActiveCategories()
	default:
		categories, errors = s.repository.GetCategories()
	}

	return categories, errors
}

// return data category by field "id"
func (s *categoryService) GetCategoryById(categoryId int) (models.Category, error) {
	return s.repository.GetCategoryById(categoryId)
}

// create new category
func (s *categoryService) CreateCategory(categoryRequest requests.CreateCategoryRequest) error {
	category := models.Category{
		Name: categoryRequest.Name,
		Slug: slug.Make(categoryRequest.Name),
	}

	err := s.repository.CreateCategory(category)

	return err
}

// update selected category
func (s *categoryService) UpdateCategory(categoryId int, categoryRequest requests.UpdateCategoryRequest) error {
	category, _ := s.repository.GetCategoryById(categoryId)

	category.Name = categoryRequest.Name
	category.Slug = slug.Make(categoryRequest.Name)
	category.Status = categoryRequest.Status

	err := s.repository.UpdateCategory(category)

	return err
}

func (s *categoryService) UpdateCategoryStatus(categoryId int) error {
	category, _ := s.repository.GetCategoryById(categoryId)

	// if Category status == false
	if !*category.Status {
		// set into true
		*category.Status = true
	} else {
		// set into false
		*category.Status = false
	}

	err := s.repository.UpdateCategory(category)

	return err
}

// delete selected category
func (s *categoryService) DeleteCategory(categoryId int) error {
	category, _ := s.repository.GetCategoryById(categoryId)

	err := s.repository.DeleteCategory(category)

	return err
}

func (s *categoryService) Pagination(pagination *helper.Pagination, query *helper.PaginationQuery, link *helper.PaginationLink, appPATH string) ([]models.Category, error) {
	result, err := s.repository.Pagination(pagination, query)

	if err != nil {
		return nil, err
	}

	// set pagination link
	helper.GeneratePaginationLink(appPATH, pagination, query, link)

	return result, nil
}
