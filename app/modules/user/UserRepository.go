package user

import (
	"fmt"
	"math"

	"github.com/novaldwp/go-news-api/app/helper"
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Pagination(pagination *helper.Pagination, query *helper.PaginationQuery) ([]models.User, error)
	GetUsers() ([]models.User, error)
	GetActiveUsers() ([]models.User, error)
	GetInactiveUsers() ([]models.User, error)
	GetUserById(userId int) (models.User, error)
	GetUserDetailById(userId int) (models.UserDetail, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByEmailAndPassword(email, password string) (models.User, error)
	CreateUser(User models.User) error
	CreateUserDetail(User models.UserDetail) error
	UpdateUser(User models.User) error
	UpdateUserDetail(User models.UserDetail) error
	DeleteUser(User models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUsers() ([]models.User, error) {
	var users []models.User

	err := r.db.Order("id desc").Find(&users).Error

	return users, err
}

func (r *userRepository) GetActiveUsers() ([]models.User, error) {
	var users []models.User

	err := r.db.Order("id desc").Where("status = 1").Find(&users).Error

	return users, err
}

func (r *userRepository) GetInactiveUsers() ([]models.User, error) {
	var users []models.User

	err := r.db.Order("id desc").Where("status = 0").Find(&users).Error

	return users, err
}

func (r *userRepository) GetUserById(userId int) (models.User, error) {
	var user models.User

	err := r.db.Where("id = ?", userId).First(&user).Error

	return user, err
}

func (r *userRepository) GetUserDetailById(userId int) (models.UserDetail, error) {
	var user models.UserDetail

	err := r.db.Where("user_id = ?", userId).First(&user).Error

	return user, err
}

func (r *userRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).Preload("Role").First(&user).Error

	return user, err
}

func (r *userRepository) GetUserByEmailAndPassword(email, password string) (models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).Where("password = ?", password).First(&user).Error

	return user, err
}

func (r *userRepository) CreateUser(user models.User) error {
	err := r.db.Create(&user).Error

	return err
}

func (r *userRepository) CreateUserDetail(user models.UserDetail) error {
	err := r.db.Create(&user).Error

	return err
}

func (r *userRepository) UpdateUser(user models.User) error {
	err := r.db.Save(&user).Error

	return err
}

func (r *userRepository) UpdateUserDetail(user models.UserDetail) error {
	err := r.db.Save(&user).Error

	return err
}

func (r *userRepository) DeleteUser(user models.User) error {
	err := r.db.Delete(&user).Error

	return err
}

func (r *userRepository) Pagination(p *helper.Pagination, q *helper.PaginationQuery) ([]models.User, error) {
	var users []models.User
	var totalRows int64 = 0

	// get data tags
	if err := r.db.Order(fmt.Sprintf("%s %s", q.Order, q.Sort)).Limit(q.GetLimit()).Offset(q.GetOffset()).Find(&users).Error; err != nil {
		return nil, err
	}

	// get total rows
	if err := r.db.Model(&users).Count(&totalRows).Error; err != nil {
		return nil, err
	}

	// get total page
	totalPages := int(math.Ceil(float64(totalRows) / float64(q.GetLimit())))

	// set pagination value
	p.TotalRows = int(totalRows)
	p.TotalPages = totalPages

	return users, nil
}
