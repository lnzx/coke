package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/lnzx/coke/api"
)

func MountUserRoutes(router fiber.Router) {
	router.Get("/users", getUsers)
	router.Post("/users", addUser)
	router.Put("/users", updateUser)
	router.Delete("/users", deleteUser)
}

func getUsers(c fiber.Ctx) error {
	fmt.Println("get users list")
	users := api.ListUsers()
	return c.JSON(fiber.Map{
		"users": users,
	})
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
