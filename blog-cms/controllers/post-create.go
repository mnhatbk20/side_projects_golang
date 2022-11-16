package controllers

import (
	"blog-cms/models"
	"strconv"
	"strings"
	"time"
	"github.com/gofiber/fiber/v2"
	// "github.com/go-playground/validator/v10"
)

func PostCreateAPI(c *fiber.Ctx) error {

	var post models.Post

	data := new(struct {
		Title       string `form:"title" validate:"required"`
		Slug        string `form:"slug" validate:"required"`
		Content     string `form:"content"`
		Status      string `form:"status"`
		CategoryID  uint   `form:"category-id"`
		Description string `form:"description"`
		Tags        string `form:"tags"`
		CreatedAt   time.Time
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
	post.CreatedAt = time.Now()

	IDTags := strings.Split(data.Tags, ",")

	var tags []models.Tag

	for _, id := range IDTags {
		idnum, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Tag List is invalid",
			})
		}
		tags = append(tags, models.Tag{ID: uint(idnum)})
	}

	post.CreatePost(tags) 

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
