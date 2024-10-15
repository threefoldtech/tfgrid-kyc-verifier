package repository

import (
	"context"

	"example.com/tfgrid-kyc-service/internal/models"
)

type TokenRepository interface {
	SaveToken(ctx context.Context, token *models.Token) error
	GetToken(ctx context.Context, clientID string) (*models.Token, error)
	DeleteToken(ctx context.Context, clientID string) error
}

type VerificationRepository interface {
	SaveVerification(ctx context.Context, verification *models.Verification) error
	GetVerification(ctx context.Context, clientID string) (*models.Verification, error)
	DeleteVerification(ctx context.Context, clientID string) error
}
