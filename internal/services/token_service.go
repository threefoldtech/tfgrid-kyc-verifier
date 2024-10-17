package services

import (
	"context"
	"fmt"

	"example.com/tfgrid-kyc-service/internal/clients/idenfy"
	"example.com/tfgrid-kyc-service/internal/clients/substrate"
	"example.com/tfgrid-kyc-service/internal/repository"
	"example.com/tfgrid-kyc-service/internal/responses"
)

type tokenService struct {
	repo            repository.TokenRepository
	idenfy          *idenfy.Idenfy
	substrate       *substrate.Substrate
	requiredBalance uint64
}

func NewTokenService(repo repository.TokenRepository, idenfy *idenfy.Idenfy, substrateClient *substrate.Substrate, requiredBalance uint64) TokenService {
	return &tokenService{repo: repo, idenfy: idenfy, substrate: substrateClient, requiredBalance: requiredBalance}
}

func (s *tokenService) GetorCreateVerificationToken(ctx context.Context, clientID string) (*responses.TokenResponseWithStatus, error) {
	token, err := s.repo.GetToken(ctx, clientID)
	if err != nil {
		return nil, err
	}
	// check if token is not nil and not expired or near expiry (2 min)
	if token != nil { //&& time.Since(token.CreatedAt)+2*time.Minute < time.Duration(token.ExpiryTime)*time.Second {
		tokenResponse := &responses.TokenResponse{
			AuthToken:     token.AuthToken,
			ClientID:      token.ClientID,
			ScanRef:       token.ScanRef,
			ExpiryTime:    token.ExpiryTime,
			SessionLength: token.SessionLength,
			DigitString:   token.DigitString,
			TokenType:     token.TokenType,
		}
		tokenResponseWithStatus := &responses.TokenResponseWithStatus{
			Token:      tokenResponse,
			IsNewToken: false,
			Message:    "Existing valid token retrieved.",
		}
		fmt.Println("token from db", token)
		return tokenResponseWithStatus, nil
	}
	fmt.Println("token is nil or expired")
	newToken, err := s.idenfy.CreateVerificationSession(ctx, clientID)
	if err != nil {
		return nil, err
	}
	fmt.Println("new token", newToken)
	err = s.repo.SaveToken(ctx, &newToken)
	if err != nil {
		fmt.Println("warning: was not able to save verification token to db", err)
	}
	tokenResponse := &responses.TokenResponse{
		AuthToken:     newToken.AuthToken,
		ClientID:      newToken.ClientID,
		ScanRef:       newToken.ScanRef,
		ExpiryTime:    newToken.ExpiryTime,
		SessionLength: newToken.SessionLength,
		DigitString:   newToken.DigitString,
		TokenType:     newToken.TokenType,
	}
	tokenResponseWithStatus := &responses.TokenResponseWithStatus{
		Token:      tokenResponse,
		IsNewToken: true,
		Message:    "New token created",
	}
	return tokenResponseWithStatus, nil
}

func (s *tokenService) DeleteToken(ctx context.Context, clientID string) error {
	return s.repo.DeleteToken(ctx, clientID)
}

func (s *tokenService) AccountHasRequiredBalance(ctx context.Context, address string) (bool, error) {
	if s.requiredBalance == 0 {
		return true, nil
	}
	balance, err := s.substrate.GetAccountBalance(address)
	if err != nil {
		return false, err
	}
	return balance >= s.requiredBalance, nil
}
