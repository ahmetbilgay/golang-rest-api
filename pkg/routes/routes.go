package routes

import (
	"github.com/gofiber/fiber/v2"
	"golang-rest-api/pkg/controllers"
)

func Routes(app *fiber.App) {
	app.Post("/register", controllers.RegisterUser)
	app.Post("/login", controllers.LoginUser)
}
