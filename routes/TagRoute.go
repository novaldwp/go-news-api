package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/modules/tag"
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

func TagRoutes(db *gorm.DB, v1 *gin.RouterGroup) {
	db.AutoMigrate(&models.Tag{})

	tagRepository := tag.NewTagRepository(db)
	tagService := tag.NewTagService(tagRepository)
	tagHandler := tag.NewTagHandler(tagService)

	route := v1.Group("/tags")
	{
		route.GET("/", tagHandler.GetTags)
		route.GET("/paginate", tagHandler.Paginate)
		route.GET("/:id", tagHandler.GetTag)
		route.POST("/", tagHandler.CreateTag)
		route.PATCH("/:id", tagHandler.UpdateTag)
		route.PATCH("/:id/status", tagHandler.UpdateTagStatus)
		route.DELETE("/:id", tagHandler.DeleteTag)
	}
}
