package controllers

import (
	"blog-cms/models"
	"blog-cms/views"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AddPost(c *fiber.Ctx) error {

	var token *jwt.Token
	token = c.Locals("token").(*jwt.Token)

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	user.GetUserById(claims.Issuer)


	var categories []models.Category
	models.GetAllCategories(&categories)

	var tags []models.Tag
	models.GetAllTags(&tags)

	view := views.PostCreate{"Add post", "Posts", user, categories, tags}

	return c.Render("add-post", view)
}

func AllPost(c *fiber.Ctx) error {

	var token *jwt.Token
	token = c.Locals("token").(*jwt.Token)

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	user.GetUserById(claims.Issuer)

	var posts []models.Post
	models.GetAllPosts(&posts)

	view := views.AllPost{"All Post", "Posts", user, posts}

	return c.Render("all-post", view)
}


func EditPost(c *fiber.Ctx) error {

	var token *jwt.Token
	token = c.Locals("token").(*jwt.Token)

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	user.GetUserById(claims.Issuer)

	idPost := c.Query("id")
	if idPost == "" {
		return c.Redirect("/dashboard/posts")
	}

	var post models.Post
	result := post.GetPostById(idPost)
	if result.RowsAffected == 0 {
		return c.Redirect("/dashboard/posts")
	}

	var categories []models.Category
	models.GetAllCategories(&categories)

	var tags []models.Tag
	models.GetAllTags(&tags)

	view := views.PostEdit{"Edit post", "Posts", user, post, categories, tags}

	return c.Render("edit-post", view)
}

