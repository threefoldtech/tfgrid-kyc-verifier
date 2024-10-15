package services

import (
	"context"

	"example.com/tfgrid-kyc-service/internal/responses"
)

type TokenService interface {
	CreateToken(ctx context.Context, clientID string) (*responses.TokenResponse, error)
	GetToken(ctx context.Context, clientID string) (*responses.TokenResponse, error)
	DeleteToken(ctx context.Context, clientID string) error
	AccountHasRequiredBalance(ctx context.Context, address string) (bool, error)
}

type VerificationService interface {
	GetVerificationData(ctx context.Context, clientID string) (*responses.VerificationDataResponse, error)
	GetVerificationStatus(ctx context.Context, clientID string) (*responses.VerificationStatusResponse, error)
}
