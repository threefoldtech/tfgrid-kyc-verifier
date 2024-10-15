package idenfy

import (
	"context"
	"net/http"

	"example.com/tfgrid-kyc-service/internal/configs"
)

type Idenfy struct {
	client          *http.Client
	accessKey       string
	secretKey       string
	baseURL         string
	callbackSignKey []byte
}

func New(config configs.IdenfyConfig) *Idenfy {
	return &Idenfy{
		baseURL:         config.BaseURL,
		client:          &http.Client{},
		accessKey:       config.APIKey,
		secretKey:       config.APISecret,
		callbackSignKey: []byte(config.CallbackSignKey),
	}
}

func (c *Idenfy) CreateVerificationSession(ctx context.Context) (interface{}, error) {
	var data interface{}

	return data, nil
}

func (c *Idenfy) ProcessVerificationResult(ctx context.Context, sessionID string) (interface{}, error) {
	var data interface{}

	return data, nil
}
