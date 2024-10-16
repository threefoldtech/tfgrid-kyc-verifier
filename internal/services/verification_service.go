package services

import (
	"context"

	"example.com/tfgrid-kyc-service/internal/clients/idenfy"
	"example.com/tfgrid-kyc-service/internal/configs"
	"example.com/tfgrid-kyc-service/internal/repository"
	"example.com/tfgrid-kyc-service/internal/responses"
)

type verificationService struct {
	repo   repository.VerificationRepository
	idenfy *idenfy.Idenfy
	config *configs.VerificationConfig
}

func NewVerificationService(repo repository.VerificationRepository, idenfyClient *idenfy.Idenfy, config *configs.VerificationConfig) VerificationService {
	return &verificationService{repo: repo, idenfy: idenfyClient, config: config}
}

func (s *verificationService) GetVerificationData(ctx context.Context, clientID string) (*responses.VerificationDataResponse, error) {
	// build responses.VerificationDataResponse from models.Verification
	verificationData := &responses.VerificationDataResponse{}
	return verificationData, nil
}

func (s *verificationService) GetVerificationStatus(ctx context.Context, clientID string) (*responses.VerificationStatusResponse, error) {
	verification, err := s.repo.GetVerification(ctx, clientID)
	if err != nil {
		return nil, err
	}
	if verification == nil {
		return nil, nil
	}
	verificationStatus := &responses.VerificationStatusResponse{
		FraudTags:      verification.Status.FraudTags,
		MismatchTags:   verification.Status.MismatchTags,
		AutoDocument:   verification.Status.AutoDocument,
		ManualDocument: verification.Status.ManualDocument,
		AutoFace:       verification.Status.AutoFace,
		ManualFace:     verification.Status.ManualFace,
		ScanRef:        verification.ScanRef,
		ClientID:       verification.ClientID,
		Status:         string(verification.Status.Overall),
	}

	return verificationStatus, nil
}

func (s *verificationService) ProcessVerificationResult(ctx context.Context, clientID string) error {
	return nil
}

func (s *verificationService) ProcessDocExpirationNotification(ctx context.Context, clientID string) error {
	return nil
}

func (s *verificationService) IsUserVerified(ctx context.Context, clientID string) (bool, error) {
	verification, err := s.repo.GetVerification(ctx, clientID)
	if err != nil {
		return false, err
	}
	if verification == nil {
		return false, nil
	}
	return verification.Status.Overall == "APPROVED" || (s.config.SuspiciousVerificationOutcome == "APPROVED" && verification.Status.Overall == "SUSPECTED"), nil
}
