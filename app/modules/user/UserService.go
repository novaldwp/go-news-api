package user

import (
	"errors"
	"time"

	"github.com/novaldwp/go-news-api/app/helper"
	"github.com/novaldwp/go-news-api/app/requests"
	"github.com/novaldwp/go-news-api/models"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	Pagination(pagination *helper.Pagination, query *helper.PaginationQuery, link *helper.PaginationLink, urlPath string) ([]models.User, error)
	GetUsers(status string) ([]models.User, error)
	GetUserById(userId int) (models.User, error)
	CreateUser(userRequest requests.CreateUserRequest) error
	UpdateUser(userId int, userRequest requests.UpdateUserRequest) error
	UpdateUserStatus(models.User) error
	UpdateUserDetail(userId int, userRequest requests.UpdateUserDetailRequest) error
	UpdateUserPassword(user models.User, userRequest requests.UpdatePasswordUserRequest) error
	DeleteUser(userId int) error
}

type userService struct {
	repository UserRepositoryInterface
}

func NewUserService(userRepository UserRepositoryInterface) *userService {
	return &userService{userRepository}
}

// fetch all users with parameter status. If status true, then will fetch active user
// which have value 1 on field "status", if status false, then will fetch non active user
// which have value 0 on field "status", if not both status true/false then will fetch all user
func (s *userService) GetUsers(status string) ([]models.User, error) {
	var users []models.User
	var errors error

	switch status {
	case "true":
		users, errors = s.repository.GetActiveUsers()
	case "false":
		users, errors = s.repository.GetInactiveUsers()
	default:
		users, errors = s.repository.GetUsers()
	}

	return users, errors
}

// return data user by field "id"
func (s *userService) GetUserById(userId int) (models.User, error) {
	return s.repository.GetUserById(userId)
}

// create new user
func (s *userService) CreateUser(userRequest requests.CreateUserRequest) error {
	user := models.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Email:     userRequest.Email,
		Password:  userRequest.Password,
		RoleID:    userRequest.RoleId,
	}

	err := s.repository.CreateUser(user)

	return err
}

// update selected user
func (s *userService) UpdateUser(userId int, userRequest requests.UpdateUserRequest) error {
	user, _ := s.repository.GetUserById(userId)

	user.FirstName = userRequest.FirstName
	user.LastName = userRequest.LastName
	user.Email = userRequest.Email
	user.Status = userRequest.Status
	user.RoleID = userRequest.RoleId

	err := s.repository.UpdateUser(user)

	return err
}

func (s *userService) UpdateUserDetail(userId int, userRequest requests.UpdateUserDetailRequest) error {
	user, _ := s.repository.GetUserDetailById(userId)

	// date converter from string to time
	dateFormat := "2006-01-02"
	dateBirth, err := time.Parse(dateFormat, userRequest.Dob)

	if err != nil {
		return errors.New("error parsing date, dob format must be yyyy-mm-dd")
	}

	// if user not found
	if user.ID == 0 {
		// create new user detail
		NewUserDetail := models.UserDetail{
			Dob:           dateBirth,
			Address:       userRequest.Address,
			UserID:        uint(userId),
			ProvinceID:    userRequest.ProvinceID,
			CityID:        userRequest.CityID,
			DistrictID:    userRequest.DistrictID,
			SubdistrictID: userRequest.SubdistrictID,
		}

		err := s.repository.CreateUserDetail(NewUserDetail)

		return err
	} else {
		// update data user detail
		user.Dob = dateBirth
		user.Address = userRequest.Address
		user.UserID = uint(userId)
		user.ProvinceID = userRequest.ProvinceID
		user.CityID = userRequest.CityID
		user.DistrictID = userRequest.DistrictID
		user.SubdistrictID = userRequest.SubdistrictID

		err := s.repository.UpdateUserDetail(user)

		return err
	}
}

func (s *userService) UpdateUserStatus(user models.User) error {

	// if user status == false
	if !*user.Status {
		// set into true
		*user.Status = true
	} else {
		// set into false
		*user.Status = false
	}

	err := s.repository.UpdateUser(user)

	return err
}

func (s *userService) UpdateUserPassword(user models.User, input requests.UpdatePasswordUserRequest) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14) // hash password
	user.Password = string(hash)
	err := s.repository.UpdateUser(user)

	return err
}

// delete selected user
func (s *userService) DeleteUser(userId int) error {
	user, _ := s.repository.GetUserById(userId)

	err := s.repository.DeleteUser(user)

	return err
}

func (s *userService) Pagination(pagination *helper.Pagination, query *helper.PaginationQuery, link *helper.PaginationLink, appPATH string) ([]models.User, error) {
	result, err := s.repository.Pagination(pagination, query)

	if err != nil {
		return nil, err
	}

	// set pagination link
	helper.GeneratePaginationLink(appPATH, pagination, query, link)

	return result, nil
}
