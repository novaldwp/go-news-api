package models

import (
	"time"
)

type Role struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"varchar(125);not null"`
	Status    *bool  `json:"status" gorm:"default:1"` // using type data pointer, so you can update value with "false / 0"
	Users     []User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
