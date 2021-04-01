package middlewares

import (
	"github.com/domesama/Go-Fiber-REST/package/structs"
	"github.com/gofiber/fiber/v2"
)

func WaifuMiddleware(c *fiber.Ctx)  error{
	 c.Locals( "waifu", structs.Waifu{"Pekora", "Usada"})
	 return c.Next()
}
