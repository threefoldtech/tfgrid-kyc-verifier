package main

import (
	"log"

	_ "example.com/tfgrid-kyc-service/api/docs"
	"example.com/tfgrid-kyc-service/internal/configs"
	"example.com/tfgrid-kyc-service/internal/server"
)

//	@title			TFGrid KYC API
//	@version		0.1.0
//	@description	This is a KYC service for TFGrid.
//	@termsOfService	http://swagger.io/terms/

// @contact.name	Codescalers Egypt
// @contact.url	https://codescalers-egypt.com
// @contact.email	info@codescalers.com
// @BasePath		/
func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	server := server.New(config)

	server.Start()
}
