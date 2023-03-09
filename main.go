package main

import "github.com/gofiber/fiber/v2"
import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
  )

func main() {

	database.Connect()
	database.AutoMigrate()
	
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    app.Listen(":4000")
}