package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Route(a *fiber.App) {
	router := a.Group("/api/v1")
	router.Get("/example", example)
}
