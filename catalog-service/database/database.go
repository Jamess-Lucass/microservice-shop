package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Jamess-Lucass/microservice-shop/catalog-service/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(log *zap.Logger) *gorm.DB {
	server := os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic("Could not parse PORT to an integar")
	}
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%d sslmode=disable", server, user, pass, name, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Sugar().Fatalf("failed to connect to the database. %v", err)
	}

	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Catalog{})
}
