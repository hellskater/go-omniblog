package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/hellskater/omniblog/pkg/controllers"
	"github.com/jinzhu/gorm"
)

// AuthMiddleware is a middleware for authentication
func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authController := controllers.NewAuthController(db)
		authController.Authenticate(c)
	}
}


