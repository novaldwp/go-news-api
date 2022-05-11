package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/modules/subdistrict"
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

func SubdistrictRoutes(db *gorm.DB, v1 *gin.RouterGroup) {
	db.AutoMigrate(&models.Subdistrict{})

	subdistrictRepository := subdistrict.NewSubdistrictRepository(db)
	subdistrictService := subdistrict.NewSubdistrictService(subdistrictRepository)
	subdistrictHandler := subdistrict.NewSubdistrictHandler(subdistrictService)

	route := v1.Group("/subdistricts")
	{
		route.GET("/", subdistrictHandler.GetSubdistricts)
		route.GET("/:id", subdistrictHandler.GetSubdistrict)
	}
}
