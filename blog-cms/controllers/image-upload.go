package controllers

import (
	"blog-cms/models"
	"fmt"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func ImageUploadAPI(c *fiber.Ctx) error {

	var token *jwt.Token
	token = c.Locals("token").(*jwt.Token)

	claims := token.Claims.(*jwt.StandardClaims)

	file, err := c.FormFile("photo")
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "upload failed",
		})

	}

	filename := "image_" + claims.Issuer + "." + strings.Split(file.Filename, ".")[1]

	c.SaveFile(file, fmt.Sprintf("./public/images/gallery/%s", filename))

	var user models.User

	user.GetUserById(claims.Issuer)	
	user.Avatar = filename
	user.UpdateUser()	

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
