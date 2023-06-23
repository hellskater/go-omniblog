package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Title      string     `gorm:"not null"`
	Body       string     `gorm:"type:text"`
	Author     User       `gorm:"ForeignKey:AuthorID"`
	AuthorID   uint       `gorm:"not null"`
	Views      int        `gorm:"default:0"`
	Likes      int        `gorm:"default:0"`
	Comments   []Comment  `gorm:"ForeignKey:PostID"`
	Tags       []Tag      `gorm:"many2many:post_tags"`
	Categories []Category `gorm:"many2many:post_categories"`
}
