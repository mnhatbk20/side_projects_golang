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

	app.Get("/dashboard", middleware.Authencation, controllers.DashboardPage)
	app.Get("/dashboard/posts", middleware.Authencation, controllers.AllPostPage)
	app.Get("/dashboard/posts/edit", middleware.Authencation, controllers.PostEditPage)
	app.Get("/dashboard/posts/add", middleware.Authencation, controllers.PostAddPage)


	app.Post("/api/posts/update/:id", middleware.AuthencationAPI, controllers.UpdatePost)
	app.Post("/api/posts/create", middleware.AuthencationAPI, controllers.CreatePost)
	app.Post("/api/upload/image", middleware.AuthencationAPI, controllers.UploadImage)
	
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)

}
