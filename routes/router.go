package routes

import (
	"ddrag23/gae-soal/features/auth"
	"ddrag23/gae-soal/features/user"
	"ddrag23/gae-soal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitRouter(app *fiber.App)  {
	// app.Get("/test",func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello Word")
	// })
	api := app.Group("/api",logger.New())
	api.Post("/login",auth.Login)
	api.Get("/", user.Index)


	userRoute := api.Group("/user",middleware.Protected())
	userRoute.Get("/",user.Index)
	userRoute.Post("/",user.Store)
}