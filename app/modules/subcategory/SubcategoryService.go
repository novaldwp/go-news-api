package subcategory

import (
	"github.com/gosimple/slug"
	"github.com/novaldwp/go-news-api/app/requests"
	"github.com/novaldwp/go-news-api/models"
)

type SubcategoryServiceInterface interface {
	GetSubcategories(status string) ([]models.Subcategory, error)
	GetSubcategoryById(subcategoryId int) (models.Subcategory, error)
	CreateSubcategory(subcategoryRequest requests.CreateSubcategoryRequest) error
	UpdateSubcategory(subsubcategoryId int, categoryRequest requests.UpdateSubcategoryRequest) error
	UpdateSubcategoryStatus(subsubcategoryId int) error
	DeleteSubcategory(subsubcategoryId int) error
}

type subcategoryService struct {
	repository SubcategoryRepositoryInterface
}

func NewSubcategoryService(subcategoryRepository SubcategoryRepositoryInterface) *subcategoryService {
	return &subcategoryService{subcategoryRepository}
}

// fetch all subcategorieswith parameter status. If status true, then will fetch active category
// which have value 1 on field "status", if status false, then will fetch non active category
// which have value 0 on field "status", if not both status true/false then will fetch all category
func (s *subcategoryService) GetSubcategories(status string) ([]models.Subcategory, error) {
	var subcategories []models.Subcategory
	var errors error

	switch status {
	case "true":
		subcategories, errors = s.repository.GetActiveSubcategories()
	case "false":
		subcategories, errors = s.repository.GetNonActiveSubcategories()
	default:
		subcategories, errors = s.repository.GetSubcategories()
	}

	return subcategories, errors
}

// return data category by field "id"
func (s *subcategoryService) GetSubcategoryById(subcategoryId int) (models.Subcategory, error) {
	return s.repository.GetSubcategoryById(subcategoryId)
}

// create new category
func (s *subcategoryService) CreateSubcategory(subcategoryRequest requests.CreateSubcategoryRequest) error {
	category := models.Subcategory{
		Name:       subcategoryRequest.Name,
		Slug:       slug.Make(subcategoryRequest.Name),
		CategoryID: subcategoryRequest.CategoryId,
	}

	err := s.repository.CreateSubcategory(category)

	return err
}

// update selected category
func (s *subcategoryService) UpdateSubcategory(subcategoryId int, subcategoryRequest requests.UpdateSubcategoryRequest) error {
	subcategory, _ := s.repository.GetSubcategoryById(subcategoryId)

	subcategory.Name = subcategoryRequest.Name
	subcategory.Slug = slug.Make(subcategoryRequest.Name)
	subcategory.Status = subcategoryRequest.Status

	err := s.repository.UpdateSubcategory(subcategory)

	return err
}

func (s *subcategoryService) UpdateSubcategoryStatus(subcategoryId int) error {
	subcategory, _ := s.repository.GetSubcategoryById(subcategoryId)

	// if Category status == false
	if !*subcategory.Status {
		// set into true
		*subcategory.Status = true
	} else {
		// set into false
		*subcategory.Status = false
	}

	err := s.repository.UpdateSubcategory(subcategory)

	return err
}

// delete selected category
func (s *subcategoryService) DeleteSubcategory(subcategoryId int) error {
	subcategory, _ := s.repository.GetSubcategoryById(subcategoryId)

	err := s.repository.DeleteSubcategory(subcategory)

	return err
}
