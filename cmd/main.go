package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/unq-arq2-ecommerce-team/payments-service/internal/application"
	"github.com/unq-arq2-ecommerce-team/payments-service/internal/infrastructure"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoUri := os.Getenv("MONGO_URI")
	mongoDbName := os.Getenv("MONGO_DATABASE")
	if mongoUri == "" || mongoDbName == "" {
		panic("env vars MONGO_URI or MONGO_DATABASE not found")
	}
	mongoConn, err := createConnection(mongoUri)
	defer func() {
		if err = mongoConn.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		panic(err)
	}

	paymentsRepository := infrastructure.NewMongoPaymentRepository(mongoConn, mongoDbName, 10*time.Second)

	createPaymentsUseCase := application.NewCreatePaymentUseCase(paymentsRepository)
	confirmPaymentsUseCase := application.NewConfirmPaymentUseCase(paymentsRepository)
	updatePaymentMethodUseCase := application.NewUpdatePaymentMethodUsecaseUseCase(paymentsRepository)
	rejectPaymentUseCase := application.NewRejectPaymentUseCase(paymentsRepository)

	app := infrastructure.NewGinApplication(createPaymentsUseCase, updatePaymentMethodUseCase, confirmPaymentsUseCase, rejectPaymentUseCase)
	log.Fatal(app.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func createConnection(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return client, nil
}
