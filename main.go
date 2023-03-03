package main

import "github.com/gofiber/fiber/v2"
import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
  )

func main() {

	_, err := gorm.Open(mysql.Open("root:root@tcp(db:3306)/ambassador"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}
	
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    app.Listen(":4000")
}