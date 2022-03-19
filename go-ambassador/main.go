package main

import (
	"ambassador/src/database"
	"ambassador/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()
	database.AutoMigrate()
	// database.DB.Exec(`DELETE FROM users WHERE Email='test@test.com'`)
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		//allow frontend to get the cookies from the backend
		AllowCredentials: true,
	}))
	routes.Setup(app)
	app.Listen(":8000")

}
