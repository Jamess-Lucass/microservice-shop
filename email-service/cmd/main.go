package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/smtp"
	"net/url"
	"os"

	"github.com/Jamess-Lucass/microservice-shop/email-service/handlers"
	"github.com/Jamess-Lucass/microservice-shop/email-service/models"
	"github.com/Jamess-Lucass/microservice-shop/email-service/services"
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

	emailService := services.NewEmailService(ch)

	server := handlers.NewServer(logger, emailService)

	q, err := ch.QueueDeclare(
		"emails", // name
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
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		logger.Sugar().Fatalf("error occured consuming rabbitMQ queue: %v", err)
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpAddress := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	tmpl := template.Must(template.ParseFiles("./templates/email-1.html"))

	var forever chan struct{}

	go func() {
		for d := range msgs {
			logger.Sugar().Infof("message received from queue: %v", string(d.Body))

			var email models.Email
			if err := json.Unmarshal(d.Body, &email); err != nil {
				logger.Sugar().Errorf("error occured unmarshalling message from queue: %v", err)
				continue
			}

			var body bytes.Buffer
			mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
			body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", email.Subject, mimeHeaders)))

			tmpl.Execute(&body, email)

			err := smtp.SendMail(smtpAddress, nil, email.From, email.To, body.Bytes())
			if err != nil {
				logger.Sugar().Errorf("error occured while sending email: %v", err)
				d.Nack(false, true)
				continue
			}

			d.Ack(false)
		}
	}()

	if err := server.Start(); err != nil {
		logger.Sugar().Fatalf("error starting web server: %v", err)
	}

	<-forever
}
