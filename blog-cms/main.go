package main

import (
	"blog-cms/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"

	"blog-cms/database"
	"blog-cms/models"
	"gorm.io/gorm"
)

func main() {

	migrate(database.Connect())

	engine := html.New("./views/templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)
	app.Listen(":8000")
}

func migrate(*gorm.DB) {
	(&models.User{}).Migrate()
	(&models.Post{}).Migrate()
	(&models.Tag{}).Migrate()
	(&models.Comment{}).Migrate()
	(&models.Category{}).Migrate()
	(&models.BlogInfo{}).Migrate()
}
