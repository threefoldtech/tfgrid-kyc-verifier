package services

import (
	"context"
	"errors"
	"fmt"

	"example.com/tfgrid-kyc-service/internal/clients/idenfy"
	"example.com/tfgrid-kyc-service/internal/clients/substrate"
	"example.com/tfgrid-kyc-service/internal/models"
	"example.com/tfgrid-kyc-service/internal/repository"
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

func (s *tokenService) GetorCreateVerificationToken(ctx context.Context, clientID string) (*models.Token, bool, error) {
	token, err := s.repo.GetToken(ctx, clientID)
	if err != nil {
		return nil, false, err
	}
	// check if token is not nil and not expired or near expiry (2 min)
	if token != nil { //&& time.Since(token.CreatedAt)+2*time.Minute < time.Duration(token.ExpiryTime)*time.Second {
		return token, false, nil
	}
	fmt.Println("token is nil or expired")
	// check if user account balance satisfies the minimum required balance, return an error if not
	hasRequiredBalance, err := s.AccountHasRequiredBalance(ctx, clientID)
	if err != nil {
		return nil, false, err // todo: implement a custom error that can be converted in the handler to a 500 status code
	}
	if !hasRequiredBalance {
		return nil, false, errors.New("account does not have the required balance") // todo: implement a custom error that can be converted in the handler to a 402 status code
	}
	newToken, err := s.idenfy.CreateVerificationSession(ctx, clientID)
	if err != nil {
		return nil, false, err
	}
	fmt.Println("new token", newToken)
	err = s.repo.SaveToken(ctx, &newToken)
	if err != nil {
		fmt.Println("warning: was not able to save verification token to db", err)
	}

	return &newToken, true, nil
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
