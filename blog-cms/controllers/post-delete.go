package controllers

import (
	"blog-cms/database"
	"blog-cms/models"

	"github.com/gofiber/fiber/v2"
)

func PostDeleteAPI(c *fiber.Ctx) error {

	idPost := c.Params("id")

	var post models.Post

	result := database.DBGorm.Where("id = ?", idPost).First(&post)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Item not exist",
		})
	}

	post.DeletePostById(idPost)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
