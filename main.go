package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shinochiha/fiber_go/controllers/bookcontroller"
	"github.com/shinochiha/fiber_go/models"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	api := app.Group("/api")

	api.Get("books", bookcontroller.GetList)
	api.Get("books/:id", bookcontroller.GetById)
	api.Post("books", bookcontroller.Create)
	api.Put("books/:id", bookcontroller.Update)
	api.Delete("books/:id", bookcontroller.Delete)

	app.Listen(":8000")
}
