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
	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASS")
	mongoConn, err := createConnection(user, password)
	defer func() {
		if err = mongoConn.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		panic(err)
	}

	paymentsRepository := infrastructure.NewMongoPaymentRepository(mongoConn, "payments_service", 10*time.Second)

	createPaymentsUseCase := application.NewCreatePaymentUseCase(paymentsRepository)
	confirmPaymentsUseCase := application.NewConfirmPaymentUseCase(paymentsRepository)
	updatePaymentMethodUseCase := application.NewUpdatePaymentMethodUsecaseUseCase(paymentsRepository)
	rejectPaymentUseCase := application.NewRejectPaymentUseCase(paymentsRepository)

	app := infrastructure.NewGinApplication(createPaymentsUseCase, updatePaymentMethodUseCase, confirmPaymentsUseCase, rejectPaymentUseCase)
	log.Fatal(app.Run(fmt.Sprintf(":%d", os.Getenv("PORT"))))
}

func createConnection(user, password string) (*mongo.Client, error) {

	connString := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.qsyiuh0.mongodb.net/?retryWrites=true&w=majority", user, password)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))
	if err != nil {
		return nil, err
	}
	return client, nil
}
