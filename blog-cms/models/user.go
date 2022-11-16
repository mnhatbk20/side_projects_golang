package models

import (
	"time"
	"blog-cms/database"
	"gorm.io/gorm"
)


type User struct {
	ID        uint `gorm:"primarykey"`
	FirstName string
	LastName  string
	UserName  string
	Password  string
	Email     string
	IsAdmin   bool `gorm:"type:bool;default:false"`
	Avatar    string
	CreatedAt time.Time
}

func (user *User) GetUserById(id string) *gorm.DB{
	return database.DBGorm.First(&user, "id = ?", id)	
}

func (user *User) CreateUser() *gorm.DB{	
	return database.DBGorm.Create(user)
}

func (user *User) UpdateUser() *gorm.DB{
	return database.DBGorm.Save(&user)
}

func UserNameExist(username string) bool{
	var user User
	var count int64
	database.DBGorm.Model(&user).Where("user_name = ?", username).Count(&count)
	return (count >0)
}


