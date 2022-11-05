package controllers

import (
	"blog-cms/database"
	"blog-cms/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type viewPostEdit struct {
	Title      string
	Breadcrumb string
	User       models.User
	Post       models.Post
	Category   []models.Category
	Tag        []models.Tag
}

func (v viewPostEdit) CheckStatus(option string) bool {
	return v.Post.Status == option
}

func (v viewPostEdit) CheckCategory(option string) bool {
	return v.Post.Category.Name == option
}
func (v viewPostEdit) CheckSelectedTag(option string) bool {
	for _, s := range v.Post.Tags {
		if (option == s.Name) {
			return true
		}
	}
	return false
}


func PostEditPage(c *fiber.Ctx) error {

	var token *jwt.Token
	token = c.Locals("token").(*jwt.Token)

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DBGorm.First(&user, "id = ?", claims.Issuer)

	idPost := c.Query("id")
	if idPost == "" {
		return c.Redirect("/dashboard/posts")
	}

	var post models.Post
	result := database.DBGorm.Preload("Tags").Preload("Category").First(&post, idPost)
	if result.RowsAffected == 0 {
		return c.Redirect("/dashboard/posts")
	}

	var categories []models.Category
	database.DBGorm.Find(&categories)

	var tags []models.Tag
	database.DBGorm.Find(&tags)

	view := viewPostEdit{"Edit post", "Posts", user, post, categories, tags}

	return c.Render("postedit", view)
}
