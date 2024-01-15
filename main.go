package main

import (
	"github.com/deliskyxd/myfilemanager/database"
	"github.com/gofiber/fiber/v2"
)

func main(){
    database.Connect()
    PORT := "8500"
    app := fiber.New()
    app.Listen(":" + PORT)
}
