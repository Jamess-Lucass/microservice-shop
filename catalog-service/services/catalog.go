package services

import (
	"github.com/Jamess-Lucass/microservice-shop/catalog-service/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CatalogService struct {
	db *gorm.DB
}

func NewCatalogService(db *gorm.DB) *CatalogService {
	return &CatalogService{
		db: db,
	}
}

func (s *CatalogService) List() *gorm.DB {
	query := s.db.Model(models.Catalog{}).Where("is_deleted <> ?", true)

	return query
}

func (s *CatalogService) Get(id uuid.UUID) (*models.Catalog, error) {
	var catalog models.Catalog
	if err := s.List().First(&catalog, id).Error; err != nil {
		return nil, err
	}

	return &catalog, nil
}
