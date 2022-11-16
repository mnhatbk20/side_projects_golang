package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/dgrijalva/jwt-go"
	"blog-cms/models"
	"blog-cms/views"
)


func Dashboard(c *fiber.Ctx) error {

	var token *jwt.Token
	token = c.Locals("token").(*jwt.Token)

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	user.GetUserById(claims.Issuer)

	view := views.Dashboard{"Dashboard", "Dashboard", user}

	return c.Render("dashboard-home", view)
}
