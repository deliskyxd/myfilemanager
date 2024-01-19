package main

import (
	"fmt"
	"log"
	"os"

	"github.com/deliskyxd/myfilemanager/database"
	"github.com/deliskyxd/myfilemanager/routes"
	jwtware "github.com/gofiber/contrib/jwt"
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
	app.Static("/style", "./src/output.css")
	app.Static("/img", "./src/img")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("./src/templates/login.html", fiber.Map{})
	})

	guardedRoute := app.Group("/files")
	guardedRoute.Use(jwtware.New(jwtware.Config{
		Claims:       &routes.Claims{},
		SigningKey:   jwtware.SigningKey{Key: []byte(routes.JwtSecretKey)},
		TokenLookup:  "cookie:access-token",
		ErrorHandler: routes.JWTErrorChecker,
	}))
	guardedRoute.Get("", func(c *fiber.Ctx) error {
		return c.Render("./src/index.html", fiber.Map{})
	})

	// API routing
	createRoutes(app)
	//Starting the web server
	log.Fatal(app.Listen(":" + port))
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello from API 👋!")
}

func createRoutes(app *fiber.App) {
	app.Get("/api", hello)
	app.Get("/api/users", routes.GetUsers)
	app.Delete("/api/users/:id", routes.DeleteUser)
	// jwt auth
	app.Post("/signup", routes.Signup)
	app.Post("/login", routes.Login)
	app.Get("/private", routes.Private)
	app.Get("/public", routes.Public)
}
