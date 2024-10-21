package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"example.com/tfgrid-kyc-service/internal/models"
)

type MongoTokenRepository struct {
	collection *mongo.Collection
}

func NewMongoTokenRepository(db *mongo.Database) TokenRepository {
	repo := &MongoTokenRepository{
		collection: db.Collection("tokens"),
	}
	repo.createTTLIndex()
	return repo
}

func (r *MongoTokenRepository) createTTLIndex() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    bson.D{{"expiresAt", 1}},
			Options: options.Index().SetExpireAfterSeconds(0),
		},
	)

	if err != nil {
		fmt.Printf("Error creating TTL index: %v\n", err)
	}
}

func (r *MongoTokenRepository) SaveToken(ctx context.Context, token *models.Token) error {
	token.CreatedAt = time.Now()
	token.ExpiresAt = token.CreatedAt.Add(time.Duration(token.ExpiryTime) * time.Second)
	_, err := r.collection.InsertOne(ctx, token)
	fmt.Println("token saved to db", err)
	return err
}

func (r *MongoTokenRepository) GetToken(ctx context.Context, clientID string) (*models.Token, error) {
	var token models.Token
	fmt.Println("clientID from repo", clientID)
	err := r.collection.FindOne(ctx, bson.M{"clientId": clientID}).Decode(&token)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("no document found")
			return nil, nil
		}
		return nil, err
	}
	// calculate duration between createdAt and now then updae expiry time with remaining time
	duration := time.Since(token.CreatedAt)
	// protect against overflow
	if duration >= time.Duration(token.ExpiryTime)*time.Second {
		return nil, nil
	}
	remainingTime := time.Duration(token.ExpiryTime)*time.Second - duration
	token.ExpiryTime = int(remainingTime.Seconds())
	return &token, nil
}

func (r *MongoTokenRepository) DeleteToken(ctx context.Context, clientID string, scanRef string) error {
	res, err := r.collection.DeleteOne(ctx, bson.M{"clientId": clientID, "scanRef": scanRef})
	if err == nil {
		fmt.Println("token deletion succeeded. deleted count: ", res.DeletedCount)
	}
	return err
}
