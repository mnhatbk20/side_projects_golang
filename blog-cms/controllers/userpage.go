package controllers

import (
	"blog-cms/database"
	"blog-cms/models"
	"github.com/gofiber/fiber/v2"
	"github.com/dgrijalva/jwt-go"

)

func UserPage(c *fiber.Ctx) error {

	var token *jwt.Token
	token = c.Locals("token").(*jwt.Token)

	claims := token.Claims.(*jwt.StandardClaims)

	id := claims.Issuer
	
	var user []models.User
	database.DBGorm.Find(&user, id)

	return c.Render("user", user[0])

}
