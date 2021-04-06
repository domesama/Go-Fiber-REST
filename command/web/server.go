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
var DB,err = config.NewMySQLSvc().InitDB()

func init()  {
	if err := cleanenv.ReadEnv(&appConf); err != nil {
		log.Fatal("Unable to initialize env")
	}
}

func Start(){
	if err != nil{
		log.Fatal(err)
	}
	//app.Use(middlewares.DatabaseMiddleware)
	app.Use(middlewares.WaifuMiddleware)
	routes.ExportWaifuRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
