package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/modules/category"
	"github.com/novaldwp/go-news-api/app/modules/subcategory"
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

func SubcategoryRoutes(db *gorm.DB, v1 *gin.RouterGroup) {
	db.AutoMigrate(&models.Subcategory{})

	categoryRepository := category.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(categoryRepository)
	subcategoryRepository := subcategory.NewSubcategoryRepository(db)
	subcategoryService := subcategory.NewSubcategoryService(subcategoryRepository)
	subcategoryHandler := subcategory.NewSubcategoryHandler(subcategoryService, categoryService)

	route := v1.Group("/subcategories")
	{
		route.GET("/", subcategoryHandler.GetSubcategories)
		route.GET("/:id", subcategoryHandler.GetSubcategory)
		route.POST("/", subcategoryHandler.CreateSubcategory)
		route.PATCH("/:id", subcategoryHandler.UpdateSubcategory)
		route.PATCH("/:id/status", subcategoryHandler.UpdateSubcategoryStatus)
		route.DELETE("/:id", subcategoryHandler.DeleteSubcategory)
	}
}
