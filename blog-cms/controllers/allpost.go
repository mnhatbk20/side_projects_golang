package controllers

import (
	"blog-cms/database"
	"blog-cms/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/itchyny/timefmt-go"

)

type viewAllPost struct {
	Title      string
	Breadcrumb string
	User       models.User
	Post       []models.Post
}

func (v viewAllPost) GetClassStatus(status string) string {
	if status == "public" {
		return "bg-green"
	} else {
		return "bg-red"
	}
}

func (v viewAllPost) GetCreateAtTimeFormat(index int) string {
	time := v.Post[index].CreatedAt
	return timefmt.Format(time, "%Y/%m/%d %H:%M:%S")

}

func AllPostPage(c *fiber.Ctx) error {

	var token *jwt.Token
	token = c.Locals("token").(*jwt.Token)

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DBGorm.First(&user, "id = ?", claims.Issuer)
	var posts []models.Post
	database.DBGorm.Preload("Tags").Preload("Category").Find(&posts)

	view := viewAllPost{"All Post", "Posts", user, posts}

	return c.Render("allpost", view)
}
