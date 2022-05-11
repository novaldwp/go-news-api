package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/modules/role"
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

func RoleRoutes(db *gorm.DB, v1 *gin.RouterGroup) {
	db.AutoMigrate(&models.Role{})

	roleRepository := role.NewRoleRepository(db)
	roleService := role.NewRoleService(roleRepository)
	roleHandler := role.NewRoleHandler(roleService)

	route := v1.Group("/roles")
	{
		route.GET("/", roleHandler.GetRoles)
		route.GET("/:id", roleHandler.GetRole)
		route.POST("/", roleHandler.CreateRole)
		route.PATCH("/:id", roleHandler.UpdateRole)
		route.PATCH("/:id/status", roleHandler.UpdateRoleStatus)
		route.DELETE("/:id", roleHandler.DeleteRole)
	}

}
