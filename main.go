package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/teamanict/marks/controllers"
)

func main() {
	// Create a new engine
	engine := html.New("./views", ".html")
	engine.Reload(true)
	engine.Layout("embed")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"Title": "Upload sheet",
		})
	})

	app.Post("/", controllers.File)
	log.Fatal(app.Listen(":3000"))
}
