package main

import (
	"fmt"
	//"github.com/deliskyxd/myfilemanager/models"
	"log"
	"os"

	"github.com/deliskyxd/myfilemanager/database"
	"github.com/deliskyxd/myfilemanager/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// connecting to the database / getting env variables
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	fmt.Println("Server is running...")

	// Basic site routing
	app.Static("/", "./src") // get files from src folder = base URL
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index.html", fiber.Map{})
	})

	createRoutes(app)
	//Starting the web server
	log.Fatal(app.Listen(":" + port))
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello from API 👋!")
}

func createRoutes(app *fiber.App) {
	app.Get("/api", hello)
	//app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	// jwt auth
	app.Post("/signup", routes.Signup)
	app.Post("/login", routes.Login)
	app.Get("/private", routes.Private)
	app.Get("/public", routes.Public)
}
