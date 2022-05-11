package role

import (
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

type RoleRepositoryInterface interface {
	GetRoles() ([]models.Role, error)
	GetActiveRoles() ([]models.Role, error)
	GetNonActiveRoles() ([]models.Role, error)
	GetRoleById(roleId int) (models.Role, error)
	CreateRole(role models.Role) error
	UpdateRole(role models.Role) error
	DeleteRole(role models.Role) error
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) GetRoles() ([]models.Role, error) {
	var roles []models.Role

	err := r.db.Order("id desc").Find(&roles).Error

	return roles, err
}

func (r *roleRepository) GetActiveRoles() ([]models.Role, error) {
	var roles []models.Role

	err := r.db.Order("id desc").Where("status = 1").Find(&roles).Error

	return roles, err
}

func (r *roleRepository) GetNonActiveRoles() ([]models.Role, error) {
	var roles []models.Role

	err := r.db.Order("id desc").Where("status = 0").Find(&roles).Error

	return roles, err
}

func (r *roleRepository) GetRoleById(roleId int) (models.Role, error) {
	var role models.Role

	err := r.db.Where("id = ?", roleId).First(&role).Error

	return role, err
}

func (r *roleRepository) CreateRole(role models.Role) error {
	err := r.db.Create(&role).Error

	return err
}

func (r *roleRepository) UpdateRole(role models.Role) error {
	err := r.db.Save(&role).Error

	return err
}

func (r *roleRepository) DeleteRole(role models.Role) error {
	err := r.db.Delete(&role).Error

	return err
}
