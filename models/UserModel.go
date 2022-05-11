package models

import "time"

type User struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	FirstName    string `json:"first_name" gorm:"varchar(125);not null"`
	LastName     string `json:"last_name" gorm:"varchar(125)"`
	Email        string `json:"email" gorm:"varchar(125);unique;not null"`
	Password     string `json:"password" gorm:"not null"`
	RefreshToken string `json:"refresh_token"`
	IsVerif      *bool  `json:"is_verif" gorm:"default:0"`
	RoleID       uint   `json:"role_id"`
	Status       *bool  `json:"status" gorm:"default:1"` // using type data pointer, so you can update value with "false / 0"
	Role         Role   `json:"roles"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
