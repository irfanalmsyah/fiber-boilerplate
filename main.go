package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/irfanalmsyah/fiber-boilerplate/config"
	"github.com/irfanalmsyah/fiber-boilerplate/database"
	"github.com/irfanalmsyah/fiber-boilerplate/router"
)

func init() {
    config.LoadEnv()
    database.ConnectDB()
}

func main() {
    app := fiber.New()

    database.ConnectDB()

    router.SetupRoutes(app)

    app.Listen(":3000")
}