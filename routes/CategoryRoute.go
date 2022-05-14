package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/modules/category"
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

func CategoryRoutes(db *gorm.DB, v1 *gin.RouterGroup) {
	db.AutoMigrate(&models.Category{})

	categoryRepository := category.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(categoryRepository)
	categoryHandler := category.NewCategoryHandler(categoryService)

	route := v1.Group("/categories")
	{
		route.GET("/", categoryHandler.GetCategories)
		route.GET("/paginate", categoryHandler.Paginate)
		route.GET("/:id", categoryHandler.GetCategory)
		route.POST("/", categoryHandler.CreateCategory)
		route.PATCH("/:id", categoryHandler.UpdateCategory)
		route.PATCH("/:id/status", categoryHandler.UpdateCategoryStatus)
		route.DELETE("/:id", categoryHandler.DeleteCategory)
	}
}
