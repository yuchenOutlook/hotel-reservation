package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuchenOutlook/hotel-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User {
		FirstName: "James",
		LastName: "Big Bond",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("James Bond")
}