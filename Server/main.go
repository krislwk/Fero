package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

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

	app.Patch("/api/changeinfo/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		newHW := Homework{}

		if err != nil {
			return c.Status(401).SendString("Invalid ID")
		}
		if err := c.BodyParser(&newHW); err != nil {
			return err
		}

		for i, hw := range hwList {
			if hw.ID == id {
				if hw.Title != newHW.Title {
					hwList[i].Title = newHW.Title
				}
				if hw.Body != newHW.Body {
					hwList[i].Body = newHW.Body
				}
				if hw.DueDate != newHW.DueDate {
					hwList[i].DueDate = newHW.DueDate
				}
				break
			}
		}

		return c.JSON(hwList)
	})

	log.Fatal(app.Listen(":4000"))
}
