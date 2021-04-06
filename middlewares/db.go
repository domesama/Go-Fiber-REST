package middlewares

import (
	"github.com/domesama/Go-Fiber-REST/config"
	"github.com/gofiber/fiber/v2"
	"log"
)


func DatabaseMiddleware(c *fiber.Ctx)  error{
	db, err:= config.NewMySQLSvc().InitDB();
	if err != nil{
		log.Fatal(err)
	}
	c.Locals("db", db)
	return c.Next()
}
