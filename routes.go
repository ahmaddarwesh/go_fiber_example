package main

import (
	"github.com/gofiber/fiber/v2"
	"shortUrl/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Post("/createPost", handlers.CreatePost)
	app.Get("/posts", handlers.Posts)
}
