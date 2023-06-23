package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	Body     string `gorm:"type:text"`
	Author   User   `gorm:"ForeignKey:AuthorID"`
	Post     Post   `gorm:"ForeignKey:PostID"`
	AuthorID uint   `gorm:"not null"`
	PostID   uint   `gorm:"not null"`
}
