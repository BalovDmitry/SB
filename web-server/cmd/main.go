package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"web-server/internal"
)

func main() {
	handler := internal.LinkHandler{
		Storage: &internal.LinkStorageMap{},
	}
	webApp := fiber.New(fiber.Config{
		ReadBufferSize: 16 * 1024})
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
	webApp.Get("/links/:external", func(c *fiber.Ctx) error {
		return handler.GetLink(c.Params("external", ""), c)
	})
	webApp.Post("/links", handler.AddLink)

	logrus.Fatal(webApp.Listen(":8080"))
}
