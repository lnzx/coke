package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/lnzx/coke/api"
	"log"
	"strconv"
)

func MountNodeRoutes(router fiber.Router) {
	router.Get("/nodes", GetNodes)
	router.Put("/nodes/:id", RenameNode)
	router.Delete("/nodes/:id", DeleteNode)
}

func GetNodes(c fiber.Ctx) error {
	user := c.Get("user")
	log.Println(c.Path(), "user:", user)
	return c.JSON(api.GetNodes(user))
}

func RenameNode(c fiber.Ctx) error {
	// 获取路径参数 ID
	id := c.Params("id")
	node := new(api.Node)
	if err := c.Bind().JSON(node); err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	log.Println(c.Path(), "id:", id, "new name:", node.Name)
	// 验证 ID 是否为数字
	n, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	r := api.RenameNode(node.Name, n)
	return c.JSON(r)
}

func DeleteNode(c fiber.Ctx) error {
	// 获取路径参数 ID
	id := c.Params("id")
	log.Println(c.Path(), "id:", id)
	// 验证 ID 是否为数字
	n, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	r := api.DeleteNode(n)
	return c.JSON(r)
}
