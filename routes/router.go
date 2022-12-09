package routes

import (
	"ddrag23/gae-soal/features/auth"
	"ddrag23/gae-soal/features/user"
	"ddrag23/gae-soal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitRouter(app *fiber.App) {
	// app.Get("/test",func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello Word")
	// })
	api := app.Group("/api", logger.New())
	api.Post("/login", auth.Login)
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to gae soal")
	})
	api.Post("/register", user.Store)
	api.Use(middleware.Protected())
	api.Use(middleware.RoleAccess())
	userRoute := api.Group("/user")
	userRoute.Get("/", user.Index)
	userRoute.Post("/", user.Store)
	userRoute.Put("/:id", user.Updated)
	userRoute.Put("/:id/change-password", user.ChangePassword)
	userRoute.Delete("/:id", user.Destroy)
}
