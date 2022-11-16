package views

import (
	"blog-cms/models"
)


type Dashboard struct {
	Title string
	Breadcrumb string
	User models.User
}