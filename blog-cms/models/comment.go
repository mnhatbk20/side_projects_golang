package models

import (
	"time"
	// "blog-cms/database"
)

type Comment struct {
	ID           uint `gorm:"primarykey"`
	UserID       uint
	PostID       uint
	PostParentID uint
	CreatedAt    time.Time
}

