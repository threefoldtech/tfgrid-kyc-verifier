package services

import (
	"context"

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

func NewTokenService(repo repository.TokenRepository, idenfy *idenfy.Idenfy, substrate *substrate.Substrate, requiredBalance uint64) TokenService {
	return &tokenService{repo: repo, idenfy: idenfy, substrate: substrate, requiredBalance: requiredBalance}
}

func (s *tokenService) CreateToken(ctx context.Context, clientID string) (*responses.TokenResponse, error) {
	token := &responses.TokenResponse{}

	return token, nil
}

func (s *tokenService) GetToken(ctx context.Context, clientID string) (*responses.TokenResponse, error) {
	token := &responses.TokenResponse{}
	return token, nil
}

func (s *tokenService) DeleteToken(ctx context.Context, clientID string) error {
	return s.repo.DeleteToken(ctx, clientID)
}

func (s *tokenService) AccountHasRequiredBalance(ctx context.Context, address string) (bool, error) {
	balance, err := s.substrate.GetAccountBalance(address)
	if err != nil {
		return false, err
	}
	return balance >= s.requiredBalance, nil
}
