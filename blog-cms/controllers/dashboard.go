package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/dgrijalva/jwt-go"
	"blog-cms/database"
	"blog-cms/models"

)
type viewDashboard struct {
	Breadcrumb string
	User models.User
}
func DashboardPage(c *fiber.Ctx) error {

	var token *jwt.Token
	token = c.Locals("token").(*jwt.Token)

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DBGorm.First(&user, "id = ?", claims.Issuer)	
	view := viewDashboard{"Dashboard",user}


	return c.Render("dashboard", view)
}
