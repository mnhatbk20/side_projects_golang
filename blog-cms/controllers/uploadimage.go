package controllers

import (
	"blog-cms/database"
	"blog-cms/models"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {

	SecretKey := os.Getenv("SecretKey")

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

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

	database.DBGorm.Where("id = ?", claims.Issuer).First(&user)
	user.Avatar = filename
	database.DBGorm.Save(&user)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
