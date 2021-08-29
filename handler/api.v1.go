package handler

import (
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	. "kovercheng/middleware"
	"kovercheng/service/database"
)

var call = resty.New().R()

func example(c *fiber.Ctx) error {
	Logger.Debug("API been called.")
	// HTTP call via go-resty
	if _, err := call.Get("https://www.microsoft.com/"); err != nil {
		return err
	}

	if err := database.TestPostgres(); err != nil {
		return err
	}

	if err := database.TestMongo(); err != nil {
		return err
	}

	if err := database.TestRedis(); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"msg": "example",
	})
}
