package role

import (
	"github.com/novaldwp/go-news-api/app/requests"
	"github.com/novaldwp/go-news-api/models"
)

type RoleServiceInterface interface {
	GetRoles(status string) ([]models.Role, error)
	GetRoleById(roleId int) (models.Role, error)
	CreateRole(request requests.CreateRoleRequest) error
	UpdateRole(roleId int, request requests.UpdateRoleRequest) error
	UpdateRoleStatus(roleId int) error
	DeleteRole(roleId int) error
}

type roleService struct {
	repository RoleRepositoryInterface
}

func NewRoleService(roleRepository RoleRepositoryInterface) *roleService {
	return &roleService{roleRepository}
}

// fetch all categories with parameter status. If status true, then will fetch active role
// which have value 1 on field "status", if status false, then will fetch non active role
// which have value 0 on field "status", if not both status true/false then will fetch all role
func (s *roleService) GetRoles(status string) ([]models.Role, error) {
	var roles []models.Role
	var errors error

	switch status {
	case "true":
		roles, errors = s.repository.GetActiveRoles()
	case "false":
		roles, errors = s.repository.GetNonActiveRoles()
	default:
		roles, errors = s.repository.GetRoles()
	}

	return roles, errors
}

// return data role by field "id"
func (s *roleService) GetRoleById(roleId int) (models.Role, error) {
	return s.repository.GetRoleById(roleId)
}

// create new role
func (s *roleService) CreateRole(request requests.CreateRoleRequest) error {
	role := models.Role{
		Name: request.Name,
	}

	err := s.repository.CreateRole(role)

	return err
}

// update selected role
func (s *roleService) UpdateRole(roleId int, request requests.UpdateRoleRequest) error {
	role, _ := s.repository.GetRoleById(roleId)

	role.Name = request.Name
	role.Status = request.Status

	err := s.repository.UpdateRole(role)

	return err
}

func (s *roleService) UpdateRoleStatus(roleId int) error {
	role, _ := s.repository.GetRoleById(roleId)

	// if role status == false
	if !*role.Status {
		// set into true
		*role.Status = true
	} else {
		// set into false
		*role.Status = false
	}

	err := s.repository.UpdateRole(role)

	return err
}

// delete selected role
func (s *roleService) DeleteRole(roleId int) error {
	role, _ := s.repository.GetRoleById(roleId)

	err := s.repository.DeleteRole(role)

	return err
}
