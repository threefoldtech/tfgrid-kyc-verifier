package configs

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	// DB
	MongoURI     string `mapstructure:"mongo_uri"`
	DatabaseName string `mapstructure:"database_name"`
	// Server
	Port string `mapstructure:"port"`
	// Idenfy
	Idenfy IdenfyConfig `mapstructure:"idenfy"`
	// TFChain
	TFChain TFChainConfig `mapstructure:"tfchain"`
	// IP limiter
	IPLimiter LimiterConfig `mapstructure:"ip_limiter"`
	// Client limiter
	IDLimiter LimiterConfig `mapstructure:"id_limiter"`
	// Verification
	Verification VerificationConfig `mapstructure:"verification"`
	// Other
	ChallengeWindow int64 `mapstructure:"challenge_window"`
}

type IdenfyConfig struct {
	APIKey          string   `mapstructure:"api_key"`
	APISecret       string   `mapstructure:"api_secret"`
	BaseURL         string   `mapstructure:"base_url"`
	CallbackSignKey string   `mapstructure:"callback_sign_key"`
	WhitelistedIPs  []string `mapstructure:"whitelisted_ips,omitempty"`
}

type TFChainConfig struct {
	WsProviderURL string `mapstructure:"ws_provider_url"`
}

type VerificationConfig struct {
	SuspiciousVerificationOutcome string `mapstructure:"suspicious_verification_outcome"`
	ExpiredDocumentOutcome        string `mapstructure:"expired_document_outcome"`
	MinBalanceToVerifyAccount     uint64 `mapstructure:"min_balance_to_verify_account"`
}

type LimiterConfig struct {
	MaxTokenRequests int `mapstructure:"max_token_requests"`
	TokenExpiration  int `mapstructure:"token_expiration"`
}

func LoadConfig() (*Config, error) {
	// replacer

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	err := viper.BindEnv("mongo_uri")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}
	err = viper.BindEnv("database_name")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}
	err = viper.BindEnv("port")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}
	err = viper.BindEnv("idenfy.api_key")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}
	err = viper.BindEnv("idenfy.api_secret")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}
	err = viper.BindEnv("idenfy.base_url")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}
	err = viper.BindEnv("idenfy.callback_sign_key")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}
	err = viper.BindEnv("idenfy.whitelisted_ips")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}
	err = viper.BindEnv("tfchain.ws_provider_url")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}
	err = viper.BindEnv("suspicious_verification_outcome")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}
	err = viper.BindEnv("expired_document_outcome")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}
	err = viper.BindEnv("challenge_window")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}
	err = viper.BindEnv("verification.min_balance_to_verify_account")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}
	err = viper.BindEnv("verification.suspicious_verification_outcome")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}
	err = viper.BindEnv("verification.expired_document_outcome")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}
	err = viper.BindEnv("ip_limiter.max_token_requests")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}
	err = viper.BindEnv("ip_limiter.token_expiration")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}
	err = viper.BindEnv("id_limiter.max_token_requests")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}
	err = viper.BindEnv("id_limiter.token_expiration")
	if err != nil {
		return nil, fmt.Errorf("error binding env variable: %w", err)
	}

	// Set default values
	// viper.SetDefault("port", "8080")
	// viper.SetDefault("max_token_requests_per_minute", 4)
	// viper.SetDefault("suspicious_verification_outcome", "verified")
	// viper.SetDefault("expired_document_outcome", "unverified")
	// viper.SetDefault("mongo_uri", "mongodb://localhost:27017")
	// viper.SetDefault("database_name", "tfgrid-kyc-db")
	// viper.SetDefault("idenfy.base_url", "https://ivs.idenfy.com")
	// viper.SetDefault("tfchain.ws_provider_url", "wss://tfchain.grid.tf")
	// viper.SetDefault("min_balance_to_verify_account", 10000000)
	// viper.SetDefault("challenge_window", 120)

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	fmt.Printf("%+v\n", config)
	return &config, nil
}
