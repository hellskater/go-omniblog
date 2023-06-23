package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name  string `gorm:"not null;unique"`
	Posts []Post `gorm:"many2many:post_categories"`
}
