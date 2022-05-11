package models

import "time"

type UserDetail struct {
	ID            uint        `json:"id" gorm:"primary_key"`
	Dob           time.Time   `json:"dob"`
	Address       string      `json:"address" gorm:"type:text"`
	UserID        uint        `json:"user_id"`
	ProvinceID    uint        `json:"province_id"`
	CityID        uint        `json:"city_id"`
	DistrictID    uint        `json:"district_id"`
	SubdistrictID uint        `json:"subdistrict_id"`
	Province      Province    `json:"provinces" gorm:"foreignKey:ProvinceID;references:id;OnUpdate:CASCADE,OnDelete:SET NULL"`
	City          City        `json:"cities" gorm:"foreignKey:CityID;references:id;OnUpdate:CASCADE,OnDelete:SET NULL"`
	District      District    `json:"districts" gorm:"foreignKey:DistrictID;references:id;OnUpdate:CASCADE,OnDelete:SET NULL"`
	Subdistrict   Subdistrict `json:"subdistricts" gorm:"foreignKey:SubdistrictID;references:id;OnUpdate:CASCADE,OnDelete:SET NULL"`
	User          User        `json:"users" gorm:"foreignKey:UserID;references:id;OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
