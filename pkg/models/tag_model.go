package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name  string `gorm:"not null;unique"`
	Posts []Post `gorm:"many2many:post_tags"`
}


