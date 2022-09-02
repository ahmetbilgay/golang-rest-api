package main

import (
	"github.com/gofiber/fiber/v2"
	"golang-rest-api/pkg/db"
	"golang-rest-api/pkg/routes"
)

func main() {
	// Init Fiber
	app := fiber.New()
	// Initialize Routes
	routes.Routes(app)
	// Connect MongoDB
	db.ConnectDB()
	// App listening on port 3000
	app.Listen(":3000")
}
