package models

import (
	"time"
	"blog-cms/database"
	"gorm.io/gorm"

)

type Post struct {
	ID          uint  `gorm:"primarykey"`
	Tags        []Tag `gorm:"many2many:post_tag;"`
	Title       string
	Slug        string
	Content     string
	Status      string
	CategoryID  uint
	Category    Category
	Description string
	CreatedAt time.Time
}

func GetAllPosts(posts *[]Post) *gorm.DB {
	return database.DBGorm.Preload("Tags").Preload("Category").Find(&posts)
}

func (post *Post) GetPostById(id string) *gorm.DB{
	return database.DBGorm.Preload("Tags").Preload("Category").Find(&post, id )
}
func (post *Post) CreatePost(tags []Tag) (*gorm.DB, error){	
	return database.DBGorm.Create(post) , database.DBGorm.Model(post).Association("Tags").Replace(tags)
}

func (post *Post) DeletePostById(id string) (*gorm.DB, error){	
	err := database.DBGorm.Model(post).Association("Tags").Clear()
	tx := database.DBGorm.Where("id = ?", id).Delete(post)
	return tx, err
}

func (post *Post) UpdatePost(tags []Tag) (*gorm.DB, error){	
	return database.DBGorm.Save(post) , database.DBGorm.Model(post).Association("Tags").Replace(tags)
}




