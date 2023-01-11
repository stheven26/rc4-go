package main

import (
	"fmt"
	"hashing-file/db"
	"hashing-file/routes"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var (
	port = "8080"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	db.SetupDB()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Routes(app)
	app.Listen(fmt.Sprintf(":%s", port))
}
