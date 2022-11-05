package controllers

import (
	"blog-cms/database"
	"blog-cms/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type viewPostCreate struct {
	Title      string
	Breadcrumb string
	User       models.User
	Category   []models.Category
	Tag        []models.Tag
}


func PostAddPage(c *fiber.Ctx) error {

	var token *jwt.Token
	token = c.Locals("token").(*jwt.Token)

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DBGorm.First(&user, "id = ?", claims.Issuer)

	var categories []models.Category
	database.DBGorm.Find(&categories)

	var tags []models.Tag
	database.DBGorm.Find(&tags)

	view := viewPostCreate{"Add post", "Posts", user, categories, tags}

	return c.Render("postadd", view)
}
