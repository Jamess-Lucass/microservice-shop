package handlers

import (
	"github.com/Jamess-Lucass/microservice-shop/catalog-service/services"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Server struct {
	validator      *validator.Validate
	logger         *zap.Logger
	catalogService *services.CatalogService
}

func NewServer(logger *zap.Logger, catalogService *services.CatalogService) *Server {
	return &Server{
		validator:      validator.New(),
		logger:         logger,
		catalogService: catalogService,
	}
}
