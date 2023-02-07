package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Homework struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {
	fmt.Println("Hello World!")
	app := fiber.New()
	hwList := []Homework{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/homeworks", func(c *fiber.Ctx) error {
		hw := Homework{}

		if err := c.BodyParser(&hw); err != nil {
			return err
		}

		hw.ID = len(hwList)
		hw.Status = false
		hwList = append(hwList, hw)
		return c.JSON(hwList)
	})

	log.Fatal(app.Listen(":4000"))
}
