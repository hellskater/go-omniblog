package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hellskater/omniblog/pkg/models"
	"github.com/hellskater/omniblog/pkg/repositories"
	"github.com/hellskater/omniblog/pkg/utils"
	"github.com/hellskater/omniblog/pkg/validators"
	"github.com/jinzhu/gorm"
)

type FollowerController struct {
	followerRepo      repositories.FollowerRepositoryInterface
	followerValidator validators.FollowerValidatorInterface
}

func NewFollowerController(db *gorm.DB) *FollowerController {
	return &FollowerController{followerRepo: repositories.NewFollowerRepository(db), followerValidator: validators.NewFollowerValidator()}
}

func (fc *FollowerController) Follow(c *gin.Context) {
	var follower models.Follower
	if err := c.ShouldBindJSON(&follower); err != nil {
		utils.SendErrorMessage(
			c,
			"Invalid JSON",
			400,
		)
		return
	}

	// Validate

	err := fc.followerValidator.ValidateCreate(&follower)
	if err != nil {
		utils.SendErrorMessage(
			c,
			err.Error(),
			400,
		)
		return
	}

	if err := fc.followerRepo.Follow(&follower); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Successfully followed"})
}

func (fc *FollowerController) Unfollow(c *gin.Context) {
	followerID, _ := strconv.ParseUint(c.Param("follower_id"), 10, 32)
	followingID, _ := strconv.ParseUint(c.Param("following_id"), 10, 32)

	if err := fc.followerRepo.Unfollow(uint(followerID), uint(followingID)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Successfully unfollowed"})
}

func (fc *FollowerController) GetFollowers(c *gin.Context) {
	followerID, _ := strconv.ParseUint(c.Param("follower_id"), 10, 32)
	followers, err := fc.followerRepo.FindByFollowerID(uint(followerID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"followers": followers})
}

func (fc *FollowerController) GetFollowing(c *gin.Context) {
	followingID, _ := strconv.ParseUint(c.Param("following_id"), 10, 32)
	followers, err := fc.followerRepo.FindByFollowingID(uint(followingID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"followers": followers})
}
