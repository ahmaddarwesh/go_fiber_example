package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

var db = Database{}

func Posts(c *fiber.Ctx) error {

	jsonData, err := json.Marshal(db.getPostsList())
	if err != nil {
		fmt.Printf(err.Error())
	}
	return c.Send(jsonData)
}

func CreatePost(c *fiber.Ctx) error {

	return c.SendString("Hello, World!")
}
