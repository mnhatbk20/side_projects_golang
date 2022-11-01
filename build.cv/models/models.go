package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName        string
	Password        string
	FirstName       string
	LastName        string
	Position        string
	DateOfBirth     time.Time
	Address         string
	Website         string
	Phone           string
	Profile         string
	Image           string
	ContactSocials  []ContactSocial
	Educations      []Education
	WorkExperiences []WorkExperience
	Referencess     []References
	Hobbies         []Hobby
	Skills          []Skill
	Email           string
}

type ContactSocial struct {
	ID     uint `gorm:"primarykey"`
	Link   string
	Icon   string
	UserID uint
}

type Education struct {
	ID         uint `gorm:"primarykey"`
	University string
	StartTime  int
	EndTime    int
	Degree     string
	UserID     uint
}

type WorkExperience struct {
	ID          uint `gorm:"primarykey"`
	Company     string
	Address     string
	Position    string
	StartTime   int
	EndTime     int
	Description string
	UserID      uint
}

type Hobby struct {
	ID     uint `gorm:"primarykey"`
	Name   string
	Icon   string
	UserID uint
}

type Skill struct {
	ID     uint `gorm:"primarykey"`
	Name   string
	Level  int
	UserID uint
}

type References struct {
	ID       uint `gorm:"primarykey"`
	Name     string
	Phone    string
	Position string
	UserID   uint
}
