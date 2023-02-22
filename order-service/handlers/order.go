package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// GET: /api/v1/orders
func (s *Server) GetOrders(c *fiber.Ctx) error {
	orders, err := s.orderService.List(c.Context())
	if err != nil {
		s.logger.Sugar().Debugf("error getting orders: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{Code: fiber.StatusNotFound, Message: "Could not load orders."})
	}

	return c.Status(fiber.StatusOK).JSON(orders)
}
