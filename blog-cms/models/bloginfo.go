package models

import (
	"time"
	"blog-cms/database"
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


func (bloginfo *BlogInfo) Migrate(){
	database.DBGorm.AutoMigrate(bloginfo)
}

