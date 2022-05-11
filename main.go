package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/novaldwp/go-news-api/app/middleware"
	"github.com/novaldwp/go-news-api/config"
	"github.com/novaldwp/go-news-api/routes"
)

func main() {
	if env := godotenv.Load(); env != nil {
		log.Fatal("Failed to load .env file")
	}

	db := config.InitDb()              // init db
	r := gin.Default()                 // init default gin
	r.Use(middleware.Authentication()) // set middleware auth
	r.Use(func(c *gin.Context) {

	})
	v1 := r.Group("api/v1") // create route group called "api/v1"

	// define all routes
	routes.AuthRoutes(db, v1)
	routes.CategoryRoutes(db, v1)
	routes.CityRoutes(db, v1)
	routes.DistrictRoutes(db, v1)
	routes.ProvinceRoutes(db, v1)
	routes.RoleRoutes(db, v1)
	routes.SubcategoryRoutes(db, v1)
	routes.SubdistrictRoutes(db, v1)
	routes.TagRoutes(db, v1)
	routes.UserRoutes(db, v1)

	r.Run()
}
