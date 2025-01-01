package handler

import "github.com/gofiber/fiber/v3"

func Login(c fiber.Ctx) error {
	return c.SendString("login")
}

func Logout(c fiber.Ctx) error {
	return c.SendString("logout")
}
