package models

import (
	"time"
)

type BlogInfo struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	BlogName   string
	AuthorName string
	Github     string
	Email      string
	Logo       string
	Favicon    string
}

type Post struct {
	ID        uint  `gorm:"primarykey"`
	Tags      []Tag `gorm:"many2many:post_tag;"`
	Title     string
	Slug      string
	Content   string
	Rating uint16
	CategoryID uint
	Category   Category
	CreatedAt time.Time
}

type Tag struct {
	ID   uint `gorm:"primarykey"`
	Slug string
	Name string
}

type Category struct {
	ID   uint `gorm:"primarykey"`
	Slug string
	Name string
}

type Comment struct {
	ID           uint `gorm:"primarykey"`
	UserID       uint
	PostID       uint
	PostParentID uint
	CreatedAt    time.Time
}

type User struct {
	ID       uint `gorm:"primarykey"`
	FirstName string
	LastName string
	UserName string
	Password string
	Email    string
	IsAdmin  bool `gorm:"type:bool;default:false"`
	Avatar string
	CreatedAt    time.Time
}
