package models

import "time"

type District struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"varchar(255);not null"`
	CityID    uint   `json:"city_id"`
	City      City   `json:"cities" gorm:"foreignKey:CityID;references:id;onUpdate:CASCADE;onDelete:SET NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
