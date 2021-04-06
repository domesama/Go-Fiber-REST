package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)


func HelloFiber(c *fiber.Ctx) error{
	return c.JSON(fiber.Map{
		"Hello": "Fiber",
	})
}

func GetWaifu(c *fiber.Ctx) error{
	json := c.Locals("waifu")
	fmt.Println(json)
	return c.JSON(json)
}

func GetWaifuFromDB(c *fiber.Ctx)  error{
}
