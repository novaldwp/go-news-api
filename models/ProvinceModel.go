package models

import "time"

type Province struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
