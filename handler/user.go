package handler

import (
	"github.com/gofiber/fiber/v3"
)

func MountUserRoutes(router fiber.Router) {
	router.Get("/users", getUsers)
	router.Post("/users", addUser)
	router.Put("/users", updateUser)
	router.Delete("/users", deleteUser)
}

func getUsers(c fiber.Ctx) error {
	return c.SendString("user list")
}

func addUser(c fiber.Ctx) error {

	return c.SendString("user list")
}

func updateUser(c fiber.Ctx) error {

	return c.SendString("user list")
}

func deleteUser(c fiber.Ctx) error {

	return c.SendString("user list")
}
