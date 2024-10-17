package idenfy

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"example.com/tfgrid-kyc-service/internal/configs"
	"example.com/tfgrid-kyc-service/internal/models"
	"github.com/valyala/fasthttp"
)

type Idenfy struct {
	client          *fasthttp.Client
	accessKey       string
	secretKey       string
	baseURL         string
	callbackSignKey []byte
}

const (
	VerificationSessionEndpoint = "/api/v2/token"
)

func New(config configs.IdenfyConfig) *Idenfy {
	return &Idenfy{
		baseURL:         config.BaseURL,
		client:          &fasthttp.Client{},
		accessKey:       config.APIKey,
		secretKey:       config.APISecret,
		callbackSignKey: []byte(config.CallbackSignKey),
	}
}

func (c *Idenfy) CreateVerificationSession(ctx context.Context, clientID string) (models.Token, error) { // TODO: Refactor
	url := c.baseURL + VerificationSessionEndpoint

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.Set("Content-Type", "application/json")

	// Set basic auth
	authStr := c.accessKey + ":" + c.secretKey
	auth := base64.StdEncoding.EncodeToString([]byte(authStr))
	req.Header.Set("Authorization", "Basic "+auth)

	jsonBody, err := json.Marshal(map[string]interface{}{
		"clientId":            clientID,
		"generateDigitString": true,
		"expiryTime":          30,
		"dummyStatus":         "APPROVED", // TODO: remove this after testing
	})
	if err != nil {
		return models.Token{}, fmt.Errorf("error marshaling request body: %w", err)
	}
	req.SetBody(jsonBody)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	fmt.Println("request", req)
	err = c.client.Do(req, resp)
	if err != nil {
		return models.Token{}, fmt.Errorf("error sending request: %w", err)
	}

	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		fmt.Println("response", resp)
		return models.Token{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}
	fmt.Println(string(resp.Body()))

	var result models.Token
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return models.Token{}, fmt.Errorf("error decoding response: %w", err)
	}

	fmt.Println(result)
	return result, nil
}

func (c *Idenfy) ProcessVerificationResult(ctx context.Context, sessionID string) (interface{}, error) {
	var data interface{}

	return data, nil
}
