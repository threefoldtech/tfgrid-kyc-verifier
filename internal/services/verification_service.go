package services

import (
	"context"
	"fmt"

	"example.com/tfgrid-kyc-service/internal/clients/idenfy"
	"example.com/tfgrid-kyc-service/internal/configs"
	"example.com/tfgrid-kyc-service/internal/models"
	"example.com/tfgrid-kyc-service/internal/repository"
)

type verificationService struct {
	repo   repository.VerificationRepository
	idenfy *idenfy.Idenfy
	config *configs.VerificationConfig
}

func NewVerificationService(repo repository.VerificationRepository, idenfyClient *idenfy.Idenfy, config *configs.VerificationConfig) VerificationService {
	return &verificationService{repo: repo, idenfy: idenfyClient, config: config}
}

func (s *verificationService) GetVerification(ctx context.Context, clientID string) (*models.Verification, error) {
	verification, err := s.repo.GetVerification(ctx, clientID)
	if err != nil {
		return nil, err
	}
	return verification, nil
}

func (s *verificationService) ProcessVerificationResult(ctx context.Context, body []byte, sigHeader string, result models.Verification) error {
	err := s.idenfy.VerifyCallbackSignature(ctx, body, sigHeader)
	if err != nil {
		return err
	}
	err = s.repo.SaveVerification(ctx, &result)
	if err != nil {
		fmt.Printf("error saving verification to the database: %v", err)
		return err
	}
	// fmt the result
	fmt.Println(result)
	return nil
}

func (s *verificationService) ProcessDocExpirationNotification(ctx context.Context, clientID string) error {
	return nil
}

func (s *verificationService) IsUserVerified(ctx context.Context, clientID string) (bool, error) {
	verification, err := s.GetVerification(ctx, clientID)
	if err != nil {
		return false, err
	}
	if verification == nil {
		return false, nil
	}
	return verification.Status.Overall == "APPROVED" || (s.config.SuspiciousVerificationOutcome == "APPROVED" && verification.Status.Overall == "SUSPECTED"), nil
}
