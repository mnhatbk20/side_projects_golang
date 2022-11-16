package views

import (
	"blog-cms/models"
	"github.com/itchyny/timefmt-go"
)

// View Post Create
type PostCreate struct {
	Title      string
	Breadcrumb string
	User       models.User
	Category   []models.Category
	Tag        []models.Tag
}

// View All Post
type AllPost struct {
	Title      string
	Breadcrumb string
	User       models.User
	Post       []models.Post
}

func (v AllPost) GetClassStatus(status string) string {
	if status == "public" {
		return "bg-green"
	} else {
		return "bg-red"
	}
}

func (v AllPost) GetCreateAtTimeFormat(index int) string {
	time := v.Post[index].CreatedAt
	return timefmt.Format(time, "%Y/%m/%d %H:%M:%S")

}

// View Post Edit
type PostEdit struct {
	Title      string
	Breadcrumb string
	User       models.User
	Post       models.Post
	Category   []models.Category
	Tag        []models.Tag
}

func (v PostEdit) CheckStatus(option string) bool {
	return v.Post.Status == option
}

func (v PostEdit) CheckCategory(option string) bool {
	return v.Post.Category.Name == option
}
func (v PostEdit) CheckSelectedTag(option string) bool {
	for _, s := range v.Post.Tags {
		if (option == s.Name) {
			return true
		}
	}
	return false
}
