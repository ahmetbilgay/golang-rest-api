package routes

import (
	"github.com/gofiber/fiber/v2"
	"golang-rest-api/pkg/controllers"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.HelloWorld)
}
