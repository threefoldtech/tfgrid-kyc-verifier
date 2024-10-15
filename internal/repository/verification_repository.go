package repository

import (
	"context"
	"time"

	"example.com/tfgrid-kyc-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoVerificationRepository struct {
	collection *mongo.Collection
}

func NewMongoVerificationRepository(db *mongo.Database) VerificationRepository {
	return &MongoVerificationRepository{
		collection: db.Collection("verifications"),
	}
}

func (r *MongoVerificationRepository) SaveVerification(ctx context.Context, verification *models.Verification) error {
	verification.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, verification)
	return err
}

func (r *MongoVerificationRepository) GetVerification(ctx context.Context, clientID string) (*models.Verification, error) {
	var verification models.Verification
	err := r.collection.FindOne(ctx, bson.M{"clientID": clientID}).Decode(&verification)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &verification, nil
}

func (r *MongoVerificationRepository) DeleteVerification(ctx context.Context, clientID string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"clientID": clientID})
	return err
}
