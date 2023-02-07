package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type Homework struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Body    string `json: "body"`
	DueDate string `json: "duedate"` //Format: DD/MM/YYYY
	Status  bool   `json:"status"`
}

func main() {
	app := fiber.New()
	hwList := []Homework{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/add", func(c *fiber.Ctx) error {
		hw := Homework{}

		if err := c.BodyParser(&hw); err != nil {
			return err
		}

		hw.ID = len(hwList)
		hw.Status = false
		hwList = append(hwList, hw)
		return c.JSON(hwList)
	})

	app.Patch("/api/changestatus/:id/:done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		var done string = c.Params("done")

		if err != nil {
			return c.Status(401).SendString("Invalid ID")
		}

		for i, hw := range hwList {
			if hw.ID == id {
				hwList[i].Status = (done == "done")
				break
			}
		}

		return c.JSON(hwList)
	})

	log.Fatal(app.Listen(":4000"))
}
