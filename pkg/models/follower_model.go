package models

import (
	"github.com/jinzhu/gorm"
)

type Follower struct {
	gorm.Model
	FollowerID uint `gorm:"not null" json:"followerId" validate:"required,number"`
	FollowedID uint `gorm:"not null" json:"followedId" validate:"required,number"`
}
