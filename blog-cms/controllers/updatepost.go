package controllers

import (
	"blog-cms/database"
	"blog-cms/models"
	"strings"
	"strconv"


	"github.com/gofiber/fiber/v2"
	// "github.com/go-playground/validator/v10"
)

func UpdatePost(c *fiber.Ctx) error {

	idPost := c.Params("id")

	var post models.Post

	result := database.DBGorm.Where("id = ?", idPost).First(&post)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Item not exist",
		})
	}

	data := new(struct {
		Title       string `form:"title" validate:"required"`
		Slug        string `form:"slug" validate:"required"`
		Content     string `form:"content"`
		Status      string `form:"status"`
		CategoryID  uint   `form:"category-id"`
		Description string `form:"description"`
		Tags        string	`form:"tags"`
	})

	if err := c.BodyParser(data); err != nil {
		return err
	}

	post.Title = data.Title
	post.Slug = data.Slug
	post.Content = data.Content
	post.Status = data.Status
	post.CategoryID = data.CategoryID
	post.Description = data.Description
	 
	IDTags := strings.Split(data.Tags, ",")

	var tags []models.Tag

	for _, id := range IDTags {
		idnum ,err := strconv.Atoi(id)
		if err !=nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Tag List is invalid",
			})
		}
		tags = append(tags, models.Tag{ID: uint(idnum)})
	}

	database.DBGorm.Save(&post)
	database.DBGorm.Model(&post).Association("Tags").Replace(tags)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
