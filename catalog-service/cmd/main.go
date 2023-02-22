package main

import (
	"github.com/Jamess-Lucass/microservice-shop/catalog-service/database"
	"github.com/Jamess-Lucass/microservice-shop/catalog-service/handlers"
	"github.com/Jamess-Lucass/microservice-shop/catalog-service/services"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	db := database.Connect(logger)

	catalogService := services.NewCatalogService(db)

	server := handlers.NewServer(logger, catalogService)

	if err := server.Start(); err != nil {
		logger.Sugar().Fatalf("error starting web server: %v", err)
	}
}
