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
	app.Get("/mycv", middleware.Authencation, controllers.MyCv)


	app.Post("/api/upload/image", middleware.Authencation, controllers.UploadImage)
	
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

}
