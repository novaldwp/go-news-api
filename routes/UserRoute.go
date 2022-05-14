package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/modules/city"
	"github.com/novaldwp/go-news-api/app/modules/district"
	"github.com/novaldwp/go-news-api/app/modules/province"
	"github.com/novaldwp/go-news-api/app/modules/role"
	"github.com/novaldwp/go-news-api/app/modules/subdistrict"
	"github.com/novaldwp/go-news-api/app/modules/user"
	"github.com/novaldwp/go-news-api/models"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB, v1 *gin.RouterGroup) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.UserDetail{})

	roleRepository := role.NewRoleRepository(db)
	roleService := role.NewRoleService(roleRepository)
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	provinceRepository := province.NewProvinceRepository(db)
	provinceService := province.NewProvinceService(provinceRepository)
	cityRepository := city.NewCityRepository(db)
	cityService := city.NewCityService(cityRepository)
	districtRepository := district.NewDistrictRepository(db)
	districtService := district.NewDistrictService(districtRepository)
	subdistrictRepository := subdistrict.NewSubdistrictRepository(db)
	subdistrictService := subdistrict.NewSubdistrictService(subdistrictRepository)
	userHandler := user.NewUserHandler(userService, roleService, provinceService, cityService, districtService, subdistrictService)

	route := v1.Group("/users")
	{
		route.GET("/", userHandler.GetUsers)
		route.GET("/paginate", userHandler.Paginate)
		route.GET("/:id", userHandler.GetUser)
		route.POST("/", userHandler.CreateUser)
		route.PATCH("/:id", userHandler.UpdateUser)
		route.PATCH("/:id/status", userHandler.UpdateUserStatus)
		route.PATCH("/:id/password", userHandler.UpdateUserPassword)
		route.PATCH("/:id/detail", userHandler.UpdateUserDetail)
		route.DELETE("/:id", userHandler.DeleteUser)
	}

}
