package main

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	. "github.com/lnzx/coke/handler"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	// /api/login endpoint does not require authentication
	app.Get("/api/login", Login)

	// /api/* endpoint requires authentication
	api := app.Group("/api", auth)

	api.Post("/logout", Logout)

	// users
	MountUserRoutes(api)

	// preauthkeys
	api.Get("/preauthkeys", func(c fiber.Ctx) error {

		return c.SendString("login")
	})

	// nodes
	api.Get("/nodes", func(c fiber.Ctx) error {

		return c.SendString("login")
	})

	// SPA (Single Page Application)
	app.Get("/*", func(c fiber.Ctx) error {
		return c.SendFile("web/index.html")
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}

func auth(c fiber.Ctx) error {
	log.Println("auth...", c.Path())

	return c.SendString("invalid token")
}
