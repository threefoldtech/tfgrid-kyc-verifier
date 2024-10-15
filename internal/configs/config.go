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
	MaxTokenRequestsPerMinute int `mapstructure:"max_token_requests_per_minute"`
	// Other
	SuspiciousVerificationOutcome string `mapstructure:"suspicious_verification_outcome"`
	ExpiredDocumentOutcome        string `mapstructure:"expired_document_outcome"`
	ChallengeWindow               int64  `mapstructure:"challenge_window"`
	MinBalanceToVerifyAccount     uint64 `mapstructure:"min_balance_to_verify_account"`
}

type IdenfyConfig struct {
	APIKey          string   `mapstructure:"api_key"`
	APISecret       string   `mapstructure:"api_secret"`
	BaseURL         string   `mapstructure:"base_url"`
	CallbackSignKey string   `mapstructure:"callback_sign_key"`
	WhitelistedIPs  []string `mapstructure:"whitelisted_ips"`
}

type TFChainConfig struct {
	WsProviderURL string `mapstructure:"ws_provider_url"`
}

func LoadConfig() (*Config, error) {
	// Replace dots with underscores for nested keys
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Make Viper read environment variables
	viper.AutomaticEnv()

	// Set default values
	viper.SetDefault("port", "8080")
	viper.SetDefault("max_token_requests_per_minute", 4)
	viper.SetDefault("suspicious_verification_outcome", "verified")
	viper.SetDefault("expired_document_outcome", "unverified")
	viper.SetDefault("mongo_uri", "mongodb://localhost:27017")
	viper.SetDefault("database_name", "tfgrid-kyc-db")
	viper.SetDefault("idenfy.base_url", "https://ivs.idenfy.com/api/v2")
	viper.SetDefault("tfchain.ws_provider_url", "wss://tfchain.grid.tf")
	viper.SetDefault("min_balance_to_verify_account", 10000000)
	viper.SetDefault("challenge_window", 120)

	config := &Config{}
	err := viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", config)
	return config, nil
}
