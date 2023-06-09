package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	setupRoutes(app)

	var err = app.Listen(":3000")
	if err != nil {
		fmt.Printf("Error: ", err)
	}

}
