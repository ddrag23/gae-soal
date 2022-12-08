package main

import (
	db "ddrag23/gae-soal/database"
	"ddrag23/gae-soal/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main()  {
	app := fiber.New()
	app.Use(cors.New())
	db.ConnectDB()
	routes.InitRouter(app)
	app.Listen(":3030")
	fmt.Println("Hello word")
}