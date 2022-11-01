package controllers

import (
	// "fmt"

	"blog-cms/database"
	"blog-cms/models"
	"github.com/gofiber/fiber/v2"
	"github.com/dgrijalva/jwt-go"

)

func MyCv(c *fiber.Ctx) error {

	var token *jwt.Token
	token = c.Locals("token").(*jwt.Token)

	claims := token.Claims.(*jwt.StandardClaims)

	id := claims.Issuer
	
	var user []models.User
	database.DBGorm.Preload("Skills").Preload("Educations").Preload("WorkExperiences").Preload("Referencess").Preload("Hobbies").Preload("ContactSocials").Find(&user, id)

	return c.Render("mycv", user[0])

}
