package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"example.com/tfgrid-kyc-service/internal/models"
)

type MongoTokenRepository struct {
	collection *mongo.Collection
}

func NewMongoTokenRepository(db *mongo.Database) TokenRepository {
	return &MongoTokenRepository{
		collection: db.Collection("tokens"),
	}
}

func (r *MongoTokenRepository) SaveToken(ctx context.Context, token *models.Token) error {
	token.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, token)
	return err
}

func (r *MongoTokenRepository) GetToken(ctx context.Context, authToken string) (*models.Token, error) {
	var token models.Token
	err := r.collection.FindOne(ctx, bson.M{"authToken": authToken}).Decode(&token)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &token, nil
}

func (r *MongoTokenRepository) DeleteToken(ctx context.Context, authToken string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"authToken": authToken})
	return err
}
