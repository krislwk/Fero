package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type Homework struct {
	ID       int    `json:"id"`
	taskName string `json:"name"`
	body     string `json:"body"`
	dueDate  string `json:"duedate"` //"DD/MM/YYYY"
	status   bool   `json:"status"`
}

func main() {
	app := fiber.New()
	hwList := []Homework{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/homeworks", func(c *fiber.Ctx) error {
		hw := &Homework{}
		err := c.BodyParser(&hw)

		if err != nil {
			return err
		}

		hw.ID = len(hwList)
		hw.status = false
		hwList = append(hwList, *hw)
		return c.JSON(hwList)
	})

	app.Get("/getHomeworks", func(c *fiber.Ctx) error {
		return c.JSON(hwList)
	})

	app.Patch("/api/homeworks/:id/:done", func(c *fiber.Ctx) error {
		id, idErr := c.ParamsInt("id")
		done := c.Params("done")

		if idErr != nil {
			return c.Status(401).SendString("Invalid ID")
		}

		for i, t := range hwList {
			if t.ID == id {
				hwList[i].status = (done == "done")
				break
			}
		}

		return c.JSON(hwList)
	})

	log.Fatal(app.Listen(":4000"))
}
