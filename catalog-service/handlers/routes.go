package handlers

import (
	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (s *Server) Start() error {
	f := fiber.New()
	f.Use(cors.New(cors.Config{AllowOrigins: "*", AllowCredentials: true, MaxAge: 60}))

	f.Use(fiberzap.New(fiberzap.Config{
		Logger: s.logger,
	}))

	f.Get("/api/v1/catalog", s.GetAllCatalogItems)
	f.Get("/api/v1/catalog/:id", s.GetCatalogItem)

	f.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{"code": fiber.StatusNotFound, "message": "No resource found"})
	})

	return f.Listen(":8080")
}
