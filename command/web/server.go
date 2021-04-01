package web

import (
	"github.com/domesama/Go-Fiber-REST/config"
	"github.com/domesama/Go-Fiber-REST/middlewares"
	"github.com/domesama/Go-Fiber-REST/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

var app = fiber.New()
var appConf = config.AppConfig{}

func init()  {
	if err := cleanenv.ReadEnv(&appConf); err != nil {
		log.Fatal("Unable to initialize env")
	}
}

func Start(){
	app.Use(middlewares.WaifuMiddleware)
	routes.ExportWaifuRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
