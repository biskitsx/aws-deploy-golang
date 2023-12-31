package main

import (
	"github.com/biskitsx/aws-deploy/database"
	"github.com/biskitsx/aws-deploy/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	database.ConnectDb()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg": "hello world",
		})
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg": "pong",
		})
	})

	app.Get("/attractions", func(c *fiber.Ctx) error {
		attractions := &[]model.Attraction{}
		database.Db.Find(attractions)
		return c.JSON(attractions)
	})

	app.Get("/attractions/:id", func(c *fiber.Ctx) error {
		attraction := &model.Attraction{}
		database.Db.Where("id = ?", c.Params("id")).First(attraction)
		if attraction.Name == "" {
			return c.JSON(fiber.Map{"msg": "no attraction with this id"})
		}
		return c.JSON(attraction)
	})

	app.Post("/attractions/", func(c *fiber.Ctx) error {
		attraction := &model.Attraction{}
		if err := c.BodyParser(attraction); err != nil {
			return c.JSON(fiber.Map{"msg": "error kub"})
		}
		database.Db.Create(attraction)
		return c.JSON(attraction)
	})

	app.Listen(":8000")
}
