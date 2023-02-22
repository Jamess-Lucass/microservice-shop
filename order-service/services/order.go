package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/Jamess-Lucass/microservice-shop/basket-service/utils"
	"github.com/Jamess-Lucass/microservice-shop/order-service/emails"
	"github.com/Jamess-Lucass/microservice-shop/order-service/models"
	"github.com/Jamess-Lucass/microservice-shop/order-service/requests"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderService struct {
	db *mongo.Database
}

func NewOrderService(db *mongo.Database) *OrderService {
	return &OrderService{
		db: db,
	}
}

func (s *OrderService) List(ctx *fasthttp.RequestCtx) ([]*models.Order, error) {
	var orders []*models.Order
	cur, err := s.db.Collection("orders").Find(ctx, nil)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var order models.Order
		err := cur.Decode(&order)
		if err != nil {
			return orders, err
		}

		orders = append(orders, &order)
	}

	if err := cur.Err(); err != nil {
		return orders, err
	}

	cur.Close(ctx)

	if len(orders) == 0 {
		return orders, mongo.ErrNoDocuments
	}

	return orders, nil
}

func (s *OrderService) Create(ctx context.Context, order *models.Order) error {
	_, err := s.db.Collection("orders").InsertOne(ctx, order)

	return err
}

type CatalogResponse struct {
	Name string `json:"name"`
}

func (s *OrderService) SendPurchasedEmail(order *models.Order) error {
	template, err := template.ParseFiles("./emails/purchase_email.html")
	if err != nil {
		return err
	}

	purchaseEmail := emails.PurchaseEmail{Name: order.Name, ID: order.ID.Hex(), Address: order.Address}
	for _, item := range order.Items {
		uri := fmt.Sprintf("%s/api/v1/catalog/%s", os.Getenv("CATALOG_SERVICE_BASE_URL"), item.CatalogId)
		body, err := utils.HttpGet(uri)
		if err != nil {
			return err
		}

		var catalog CatalogResponse
		if err := json.NewDecoder(body).Decode(&catalog); err != nil {
			return err
		}

		purchaseEmailItem := emails.PurchaseEmailItem{
			Name:     catalog.Name,
			Price:    item.Price,
			Quantity: item.Quantity,
		}
		purchaseEmail.Items = append(purchaseEmail.Items, purchaseEmailItem)
	}

	var html bytes.Buffer
	template.Execute(&html, purchaseEmail)

	email := requests.Email{To: []string{order.Email}, From: "no-reply@microservice-shop.com", Subject: "Your order is being processed", Body: html.String()}

	uri := fmt.Sprintf("%s/api/v1/emails", os.Getenv("EMAIL_SERVICE_BASE_URL"))

	body, err := json.Marshal(email)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", uri, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return err
	}

	return err
}
