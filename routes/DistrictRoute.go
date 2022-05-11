package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/modules/district"
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

func DistrictRoutes(db *gorm.DB, v1 *gin.RouterGroup) {
	db.AutoMigrate(&models.District{})

	districtRepository := district.NewDistrictRepository(db)
	districtService := district.NewDistrictService(districtRepository)
	districtHandler := district.NewDistrictHandler(districtService)

	route := v1.Group("/districts")
	{
		route.GET("/", districtHandler.GetDistricts)
		route.GET("/:id", districtHandler.GetDistrict)
	}
}
