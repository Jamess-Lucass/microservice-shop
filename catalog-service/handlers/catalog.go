package handlers

import (
	"github.com/Jamess-Lucass/microservice-shop/catalog-service/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GET: /api/v1/catalog
func (s *Server) GetAllCatalogItems(c *fiber.Ctx) error {
	var items []models.Catalog
	if err := s.catalogService.List().Find(&items).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": err.Error()})
	}

	return c.Status(200).JSON(items)
}

// GET: /api/v1/catalog/1
func (s *Server) GetCatalogItem(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	item, err := s.catalogService.Get(id)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(item)
}
