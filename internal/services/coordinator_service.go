package services

import (
	"context"
	"errors"

	"example.com/tfgrid-kyc-service/internal/models"
)

type coordinatorService struct {
	tokenService        TokenService
	verificationService VerificationService
}

func NewCoordinatorService(tokenService TokenService, verificationService VerificationService) CoordinatorService {
	return &coordinatorService{tokenService: tokenService, verificationService: verificationService}
}

func (s *coordinatorService) GetorCreateVerificationToken(ctx context.Context, clientID string) (*models.Token, bool, error) {
	// check if user is unverified, return an error if not
	// this should be client responsibility to check if they are verified before requesting a new verification
	isVerified, err := s.verificationService.IsUserVerified(ctx, clientID)
	if err != nil {
		return nil, false, err
	}
	if isVerified {
		return nil, false, errors.New("user already verified") // TODO: implement a custom error that can be converted in the handler to a 400 status code
	}
	token, isNew, err := s.tokenService.GetorCreateVerificationToken(ctx, clientID)
	if err != nil {
		return nil, false, err
	}
	return token, isNew, nil
}
