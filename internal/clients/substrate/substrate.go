package substrate

import (
	"fmt"
	"sync"

	"example.com/tfgrid-kyc-service/internal/configs"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/vedhavyas/go-subkey/v2"
)

type Substrate struct {
	api *gsrpc.SubstrateAPI
	mu  sync.Mutex // TODO: Check if SubstrateAPI is thread safe
}

func New(config configs.TFChainConfig) (*Substrate, error) {
	api, err := gsrpc.NewSubstrateAPI(config.WsProviderURL)
	if err != nil {
		return nil, fmt.Errorf("substrate connection error: failed to initialize Substrate client: %w", err)
	}

	chain, _ := api.RPC.System.Chain()
	nodeName, _ := api.RPC.System.Name()
	nodeVersion, _ := api.RPC.System.Version()
	fmt.Println("conected to chain:", chain, "| nodeName:", nodeName, "| nodeVersion:", nodeVersion)

	c := &Substrate{
		api: api,
		mu:  sync.Mutex{},
	}
	return c, nil
}

func (c *Substrate) GetAccountBalance(address string) (uint64, error) {
	_, pubkeyBytes, err := subkey.SS58Decode(address)
	if err != nil {
		return 0, fmt.Errorf("failed to decode ss58 address: %w", err)
	}
	account, err := types.NewAddressFromAccountID(pubkeyBytes)
	if err != nil {
		return 0, fmt.Errorf("failed to create AccountID: %w", err)
	}
	meta, err := c.api.RPC.State.GetMetadataLatest()
	if err != nil {
		return 0, fmt.Errorf("failed to get metadata: %w", err)
	}
	// Create a storage key for the account's balance
	key, err := types.CreateStorageKey(meta, "System", "Account", account.AsAccountID.ToBytes())
	if err != nil {
		return 0, fmt.Errorf("failed to create storage key: %w", err)
	}

	// Query the storage
	var accountInfo types.AccountInfo
	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		return 0, fmt.Errorf("failed to get storage: %w", err)
	}
	if !ok {
		return 0, nil // account not found
	}

	return accountInfo.Data.Free.Uint64(), nil
}
