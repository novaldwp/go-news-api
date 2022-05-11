package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/modules/auth"
	"github.com/novaldwp/go-news-api/app/modules/user"
	"gorm.io/gorm"
)

func AuthRoutes(db *gorm.DB, v1 *gin.RouterGroup) {

	userRepository := user.NewUserRepository(db)
	authService := auth.NewAuthService(userRepository)
	authHandler := auth.NewAuthHandler(authService)

	v1.POST("/login", authHandler.Login)

}
