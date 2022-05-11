package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/novaldwp/go-news-api/app/helper"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		currPath := c.FullPath()
		loginPath := "/api/v1/login"

		if currPath != loginPath {
			tokenString := c.GetHeader("Authorization")

			if tokenString == "" {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Invalid token provided",
				})
				c.Abort()
				return
			}

			err := helper.ValidateToken(tokenString)

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
