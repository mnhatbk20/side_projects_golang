package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/dgrijalva/jwt-go"
	"blog-cms/database"
	"blog-cms/models"

)

type dataViewAllPostPage struct {
	User models.User
	Post []models.Post
}
func AllPostPage(c *fiber.Ctx) error {

	var token *jwt.Token
	token = c.Locals("token").(*jwt.Token)

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DBGorm.First(&user, "id = ?", claims.Issuer)
	var posts []models.Post
	database.DBGorm.Preload("Tags").Find(&posts)
	dataView :=dataViewAllPostPage{user, posts}

	return c.Render("allpost", dataView)
}
