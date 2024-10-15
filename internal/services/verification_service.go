package services

import (
	"context"

	"example.com/tfgrid-kyc-service/internal/repository"
	"example.com/tfgrid-kyc-service/internal/responses"
)

type verificationService struct {
	repo repository.VerificationRepository
}

func NewVerificationService(repo repository.VerificationRepository) VerificationService {
	return &verificationService{repo: repo}
}

func (s *verificationService) GetVerificationData(ctx context.Context, clientID string) (*responses.VerificationDataResponse, error) {
	// build responses.VerificationDataResponse from models.Verification
	verificationData := &responses.VerificationDataResponse{}
	return verificationData, nil
}

func (s *verificationService) GetVerificationStatus(ctx context.Context, clientID string) (*responses.VerificationStatusResponse, error) {
	verificationStatus := &responses.VerificationStatusResponse{}
	return verificationStatus, nil
}
