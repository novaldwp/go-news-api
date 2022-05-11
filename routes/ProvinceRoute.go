package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/modules/province"
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

func ProvinceRoutes(db *gorm.DB, v1 *gin.RouterGroup) {
	db.AutoMigrate(&models.Province{})

	provinceRepository := province.NewProvinceRepository(db)
	provinceService := province.NewProvinceService(provinceRepository)
	provinceHandler := province.NewProvinceHandler(provinceService)

	route := v1.Group("/provinces")
	{
		route.GET("/", provinceHandler.GetProvinces)
		route.GET("/:id", provinceHandler.GetProvince)
	}
}
