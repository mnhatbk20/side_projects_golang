package routes

import (
	"blog-cms/controllers"
	"blog-cms/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Static("/assets", "./public")

	app.Get("/", controllers.HomePage)
	app.Get("/login", controllers.LoginPage)
	app.Get("/register", controllers.RegisterPage)

	app.Get("/dashboard", middleware.Authencation, controllers.Dashboard)
	app.Get("/dashboard/posts", middleware.Authencation, controllers.AllPost)
	app.Get("/dashboard/posts/edit", middleware.Authencation, controllers.EditPost)
	app.Get("/dashboard/posts/add", middleware.Authencation, controllers.AddPost)

	app.Delete("/api/posts/delete/:id", middleware.AuthencationAPI, controllers.PostDeleteAPI)
	app.Post("/api/posts/update/:id", middleware.AuthencationAPI, controllers.PostUpdateAPI)
	app.Post("/api/posts/create", middleware.AuthencationAPI, controllers.PostCreateAPI)
	app.Post("/api/upload/image", middleware.AuthencationAPI, controllers.ImageUploadAPI)
	
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)

}
