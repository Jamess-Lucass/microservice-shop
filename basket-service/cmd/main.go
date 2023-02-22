package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/Jamess-Lucass/microservice-shop/basket-service/database"
	"github.com/Jamess-Lucass/microservice-shop/basket-service/handlers"
	"github.com/Jamess-Lucass/microservice-shop/basket-service/services"
	"github.com/rabbitmq/amqp091-go"
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

	basketService := services.NewBasketService(db)

	server := handlers.NewServer(logger, ch, basketService)

	if err := server.Start(); err != nil {
		logger.Sugar().Fatalf("error starting web server: %v", err)
	}
}
