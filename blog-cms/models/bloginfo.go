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
