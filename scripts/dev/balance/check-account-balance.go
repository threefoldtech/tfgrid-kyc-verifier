// Use substarte client to get account free balance for development use
package main

import (
	"fmt"

	"example.com/tfgrid-kyc-service/internal/clients/substrate"
	"example.com/tfgrid-kyc-service/internal/configs"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}
	substrateClient, err := substrate.New(config.TFChain)
	if err != nil {
		panic(err)
	}
	free_balance, err := substrateClient.GetAccountBalance("5DFkH2fcqYecVHjfgAEfxgsJyoEg5Kd93JFihfpHDaNoWagJ")
	if err != nil {
		panic(err)
	}
	fmt.Println("balance: ", free_balance)
}
