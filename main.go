package main

import (
	"fmt"
	//"github.com/deliskyxd/myfilemanager/models"
	"github.com/deliskyxd/myfilemanager/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
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

    //Routers(app)
    app.Get("/hello", hello)
	//Starting the web server
	log.Fatal(app.Listen(":" + port))
}

//func Routers(app *fiber.App) {
//    app.Get("/users", models.GetUsers)
//    app.Get("/users/:id", models.GetUser)
//    app.Post("/users", models.CreateUser)
//    app.Delete("/users/:id", models.DeleteUser)
//    app.Put("/users/:id", models.UpdateUser)
//}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
