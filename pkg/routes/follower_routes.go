package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hellskater/omniblog/pkg/controllers"
	"github.com/hellskater/omniblog/pkg/middlewares"
	"github.com/jinzhu/gorm"
)

func FollowerRoutes(r *gin.Engine, db *gorm.DB) {
	followerController := controllers.NewFollowerController(db)
	r.Use(
		middlewares.AuthMiddleware(db),
	)
	r.POST("/follow", followerController.Follow)
	r.DELETE("/unfollow/:follower_id/:following_id", followerController.Unfollow)
	r.GET("/followers/:follower_id", followerController.GetFollowers)
	r.GET("/following/:following_id", followerController.GetFollowing)
}
