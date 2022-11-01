package middleware
import (

	"os"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)


func Authencation(c *fiber.Ctx) (error) {
	
	SecretKey := os.Getenv("SecretKey")

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Redirect("/login")
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		}) 
	}

	
	c.Locals("token", token)
	return c.Next()

  }