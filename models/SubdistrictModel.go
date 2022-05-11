package models

import "time"

type Subdistrict struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	Name       string   `json:"name" gorm:"varchar(255)"`
	DistrictID uint     `json:"district_id"`
	District   District `json:"district" gorm:"foreignKey:DistrictID;references:id;onUpdate:CASCADE;onDelete:SET NULL"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
