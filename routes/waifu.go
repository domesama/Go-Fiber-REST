package routes

import (
	"github.com/domesama/Go-Fiber-REST/handlers"
	"github.com/gofiber/fiber/v2"
)

func ExportWaifuRoutes(app *fiber.App)  {
	app.Get("/", handlers.HelloFiber)
	app.Get("/waifus", handlers.GetWaifu)
}
