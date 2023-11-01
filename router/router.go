package router

import (
	"github.com/irfanalmsyah/fiber-boilerplate/controller"
	"github.com/irfanalmsyah/fiber-boilerplate/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))

	app.Route("/auth", func(router fiber.Router) {
		router.Post("/register", controller.SignUpUser)
		router.Post("/login", controller.LogInUser)
		router.Get("/logout", middleware.LoginRequired, controller.LogoutUser)
	})

	app.Get("/users/me", middleware.LoginRequired, controller.GetMe)

}
