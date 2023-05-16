package infrastructure

import (
	"context"
	"time"

	"github.com/unq-arq2-ecommerce-team/payments-service/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoPaymentRepository struct {
	client  *mongo.Client
	dbName  string
	timeout time.Duration
}

func NewMongoPaymentRepository(client *mongo.Client, dbName string, timeout time.Duration) *mongoPaymentRepository {
	return &mongoPaymentRepository{
		client:  client,
		dbName:  dbName,
		timeout: timeout,
	}
}

func (r *mongoPaymentRepository) Save(payment *domain.Payment) (*domain.Payment, error) {
	_, err := r.client.Database(r.dbName).Collection("products").InsertOne(context.Background(), payment)
	return payment, err
}

func (r *mongoPaymentRepository) Update(payment *domain.Payment) (*domain.Payment, error) {
	_, err := r.client.Database(r.dbName).Collection("products").UpdateOne(context.Background(), map[string]string{"id": payment.ID}, map[string]interface{}{"$set": map[string]interface{}{"method": payment.Method}})
	return payment, err
}

func (r *mongoPaymentRepository) Find(id string) (*domain.Payment, error) {
	var payment domain.Payment
	err := r.client.Database(r.dbName).Collection("products").FindOne(context.Background(), map[string]string{"id": id}).Decode(&payment)
	return &payment, err
}
