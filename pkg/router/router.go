package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hellskater/omniblog/pkg/routes"
	"github.com/hellskater/omniblog/pkg/utils"
	"github.com/jinzhu/gorm"
)

func Initialize(router *gin.Engine, db *gorm.DB) {

	// Health check GET
	router.GET("/health", func(c *gin.Context) {
		payload := map[string]string{
			"status": "OK",
		}

		utils.SendSuccessMessage(
			c,
			"OK",
			payload,
			200,
		)
	})

	// Health check POST
	router.POST("/health", func(c *gin.Context) {
		var response interface{}
		c.BindJSON(&response)

		utils.SendSuccessMessage(
			c,
			"OK",
			response,
			200,
		)
	})

	// List of routes
	routes.FollowerRoutes(router, db)

	// No route found
	router.NoRoute(func(c *gin.Context) {
		utils.SendErrorMessage(
			c,
			"Route not found",
			404,
		)
	},
	)
}
