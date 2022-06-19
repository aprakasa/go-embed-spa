package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aprakasa/go-embed-spa/frontend"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func main() {
	app := fiber.New()
	app.Get("/hello.json", handleHello)
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         frontend.BuildHTTPFS(),
		NotFoundFile: "index.html",
	}))
	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}

func handleHello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "hello from the fiber server"})
}
