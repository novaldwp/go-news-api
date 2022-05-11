package models

import (
	"time"
)

type Subcategory struct {
	ID         uint     `json:"id" gorm:"primary_key"`
	Name       string   `json:"name" gorm:"varchar(125);not null"`
	Slug       string   `json:"slug" gorm:"unique;not null"`
	CategoryID uint     `json:"category_id"`
	Status     *bool    `json:"status" gorm:"default:1"` // using type data pointer, so you can update value with "false / 0"
	Category   Category `json:"categories"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
