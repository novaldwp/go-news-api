package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/modules/city"
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

func CityRoutes(db *gorm.DB, v1 *gin.RouterGroup) {
	db.AutoMigrate(&models.City{})

	cityRepository := city.NewCityRepository(db)
	cityService := city.NewCityService(cityRepository)
	cityHandler := city.NewCityHandler(cityService)

	route := v1.Group("/cities")
	{
		route.GET("/", cityHandler.GetCities)
		route.GET("/:id", cityHandler.GetCity)
	}
}
