package requests

// config request
type CreateUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	RoleId    uint   `json:"role_id"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" binding:"required"`
	RoleId    uint   `json:"role_id" binding:"required"`
	Status    *bool  `json:"status" gorm:"default:1"` // using type data pointer, so you can update value with "false / 0"
}

type UpdatePasswordUserRequest struct {
	CurrentPassword      string `json:"current_password" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
}

type UpdateUserDetailRequest struct {
	Dob           string `json:"dob" binding:"required"`
	Address       string `json:"address" binding:"required"`
	ProvinceID    uint   `json:"province_id" binding:"required"`
	CityID        uint   `json:"city_id" binding:"required"`
	DistrictID    uint   `json:"district_id" binding:"required"`
	SubdistrictID uint   `json:"subdistrict_id" binding:"required"`
}
