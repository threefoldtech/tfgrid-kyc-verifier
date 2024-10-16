package services

import (
	"context"

	"example.com/tfgrid-kyc-service/internal/responses"
)

type TokenService interface {
	GetorCreateVerificationToken(ctx context.Context, clientID string) (*responses.TokenResponseWithStatus, error)
	DeleteToken(ctx context.Context, clientID string) error
	AccountHasRequiredBalance(ctx context.Context, address string) (bool, error)
}

type VerificationService interface {
	GetVerificationData(ctx context.Context, clientID string) (*responses.VerificationDataResponse, error)
	GetVerificationStatus(ctx context.Context, clientID string) (*responses.VerificationStatusResponse, error)
	ProcessVerificationResult(ctx context.Context, clientID string) error
	ProcessDocExpirationNotification(ctx context.Context, clientID string) error
	IsUserVerified(ctx context.Context, clientID string) (bool, error)
}
