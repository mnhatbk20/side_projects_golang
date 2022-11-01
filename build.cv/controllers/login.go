package controllers

import (
	// "fmt"
	"github.com/gofiber/fiber/v2"
)

func LoginPage(c *fiber.Ctx) error {

	return c.Render("login", nil)

}
