package handlers

import (
	"github.com/Jamess-Lucass/microservice-shop/email-service/services"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Server struct {
	validator    *validator.Validate
	logger       *zap.Logger
	emailService *services.EmailService
}

func NewServer(logger *zap.Logger, emailService *services.EmailService) *Server {
	return &Server{
		validator:    validator.New(),
		logger:       logger,
		emailService: emailService,
	}
}
