package models

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID          primitive.ObjectID `bson:"_id"`
	Address     string             `bson:"address"`
	Email       string             `bson:"email"`
	Name        string             `bson:"name"`
	PhoneNumber string             `bson:"phoneNumber"`
	Items       []OrderItem        `bson:"items"`
}

type OrderItem struct {
	ID        primitive.ObjectID `bson:"_id"`
	CatalogId uuid.UUID          `bson:"catalogId"`
	Price     string             `bson:"price"`
	Quantity  uint               `bson:"quantity"`
}
