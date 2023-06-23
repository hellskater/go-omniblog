package models

import "github.com/jinzhu/gorm"

type Like struct {
	gorm.Model
	LikerID     uint `gorm:"not null"`
	LikedPostID uint `gorm:"not null"`
}
