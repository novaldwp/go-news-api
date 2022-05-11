package models

import "time"

type City struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	Name       string   `json:"name" gorm:"varchar(255)"`
	ProvinceID uint     `json:"province_id"`
	Province   Province `json:"provinces" gorm:"foreignKey:ProvinceID;references:id;onUpdate:CASCADE;onDelete:SET NULL"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
