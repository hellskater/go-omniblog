package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName    string    `gorm:"not null" json:"firstName" validate:"required,min=2,max=20"`
	LastName     string    `gorm:"not null" json:"lastName" validate:"required,min=2,max=20"`
	Username     string    `gorm:"not null;unique" json:"username" validate:"required,min=3,max=20"`
	Email        string    `gorm:"not null;unique" json:"email" validate:"required,email"`
	UserType     string    `gorm:"not null" json:"userType" validate:"required,oneof=user admin"`
	Password     string    `gorm:"not null" json:"-" validate:"required,min=8,max=30"`
	Bio          string    `gorm:"type:text" json:"bio" validate:"max=300"`
	AvatarURL    string    `gorm:"type:text" json:"avatarUrl" validate:"url"`
	AccessToken  string    `gorm:"type:text"`
	RefreshToken string    `gorm:"type:text"`
	Posts        []Post    `gorm:"ForeignKey:AuthorID" json:"posts"`
	Followers    []*User   `gorm:"many2many:followers" json:"followers"`
	Likes        []*Post   `gorm:"many2many:likes;association_jointable_foreignkey:liked_post_id" json:"likes"`
	Comments     []Comment `gorm:"ForeignKey:AuthorID" json:"comments"`
}
