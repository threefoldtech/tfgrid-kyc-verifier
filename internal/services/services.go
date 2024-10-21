package services

import (
	"context"

	"example.com/tfgrid-kyc-service/internal/models"
)

type TokenService interface {
	GetorCreateVerificationToken(ctx context.Context, clientID string) (*models.Token, bool, error)
	DeleteToken(ctx context.Context, clientID string) error
	AccountHasRequiredBalance(ctx context.Context, address string) (bool, error)
}

type VerificationService interface {
	GetVerification(ctx context.Context, clientID string) (*models.Verification, error)
	GetVerificationStatus(ctx context.Context, clientID string) (*models.VerificationOutcome, error)
	ProcessVerificationResult(ctx context.Context, body []byte, sigHeader string, result models.Verification) error
	ProcessDocExpirationNotification(ctx context.Context, clientID string) error
	IsUserVerified(ctx context.Context, clientID string) (bool, error)
}

// The existing services (TokenService and VerificationService) already encapsulate most of the logic.
// This coordinator service would orchestrate operations between these services,
// encapsulating the logic that spans over them while keeping them focused on their specific domains.
type CoordinatorService interface {
	GetorCreateVerificationToken(ctx context.Context, clientID string) (*models.Token, bool, error)
}
