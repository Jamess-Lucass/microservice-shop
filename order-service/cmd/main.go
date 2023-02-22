package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/Jamess-Lucass/microservice-shop/order-service/database"
	"github.com/Jamess-Lucass/microservice-shop/order-service/handlers"
	"github.com/Jamess-Lucass/microservice-shop/order-service/models"
	"github.com/Jamess-Lucass/microservice-shop/order-service/requests"
	"github.com/Jamess-Lucass/microservice-shop/order-service/services"
	"github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	atom := zap.NewAtomicLevel()
	atom.SetLevel(zapcore.DebugLevel)

	// logger, err := zap.NewProduction()
	// if err != nil {
	// 	panic(err)
	// }

	encoderCfg := zap.NewProductionEncoderConfig()
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))

	defer logger.Sync()

	db := database.Connect(logger)

	orderService := services.NewOrderService(db)

	server := handlers.NewServer(logger, orderService)

	// Rabbit MQ
	user := os.Getenv("RABBITMQ_USERNAME")
	pass := os.Getenv("RABBITMQ_PASSWORD")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")

	u := &url.URL{
		Scheme: "amqp",
		User:   url.UserPassword(user, pass),
		Host:   fmt.Sprintf("%s:%s", host, port),
	}

	rabbitMQClient, err := amqp091.Dial(u.String())
	if err != nil {
		logger.Sugar().Fatalf("error occured connecting to rabbit MQ: %v", err)
	}
	defer rabbitMQClient.Close()

	ch, err := rabbitMQClient.Channel()
	if err != nil {
		logger.Sugar().Fatalf("error occured opening rabbitMQ channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"orders", // name
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		logger.Sugar().Fatalf("error occured delcaring rabbitMQ queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		logger.Sugar().Fatalf("error occured consuming rabbitMQ queue: %v", err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			logger.Sugar().Infof("processing message %s", string(d.Body))

			var request requests.CreateOrderRequest
			if err := json.Unmarshal(d.Body, &request); err != nil {
				logger.Sugar().Errorf("error occured unmarshalling queue message: %v", err)
				continue
			}

			order := &models.Order{
				ID:          primitive.NewObjectID(),
				Address:     request.Address,
				Email:       request.Email,
				Name:        request.Name,
				PhoneNumber: request.PhoneNumber,
			}

			for _, item := range request.Basket.Items {
				orderItem := models.OrderItem{
					ID:        primitive.NewObjectID(),
					CatalogId: item.CatalogId,
					Price:     fmt.Sprintf("%f", item.Price),
					Quantity:  item.Quantity,
				}
				order.Items = append(order.Items, orderItem)
			}

			if err := orderService.Create(context.TODO(), order); err != nil {
				logger.Sugar().Errorf("error creating order record: %v", err)
				continue
			}

			if err := orderService.SendPurchasedEmail(order); err != nil {
				logger.Sugar().Errorf("error sending order purchased email: %v", err)
			}

			logger.Sugar().Infof("order :%v", order)
		}
	}()

	logger.Info("waiting for notifications from message broker")

	if err := server.Start(); err != nil {
		logger.Sugar().Fatalf("error starting web server: %v", err)
	}

	<-forever
}
