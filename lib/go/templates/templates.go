package templates

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../transactions -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../transactions/...
//go:generate go run ./cmd/manifest/main.go ./cmd/manifest/manifest.go manifest-fusd.testnet.json --network testnet
//go:generate go run ./cmd/manifest/main.go ./cmd/manifest/manifest.go manifest-fusd.mainnet.json --network mainnet

import (
	"fmt"
	"strings"

	_ "github.com/kevinburke/go-bindata"
	_ "github.com/psiemens/sconfig"
	_ "github.com/spf13/cobra"
)

const (
	placeholderFungibleTokenAddress = "0xFUNGIBLETOKENADDRESS"
	placeholderFUSDAddress          = "0xFUSDADDRESS"
)

type Environment struct {
	Network              string
	FungibleTokenAddress string
	FUSDAddress          string
}

func withHexPrefix(address string) string {
	if address == "" {
		return ""
	}

	if address[0:2] == "0x" {
		return address
	}

	return fmt.Sprintf("0x%s", address)
}

func replaceAddresses(code string, env Environment) string {

	code = strings.ReplaceAll(
		code,
		placeholderFungibleTokenAddress,
		withHexPrefix(env.FungibleTokenAddress),
	)

	code = strings.ReplaceAll(
		code,
		placeholderFUSDAddress,
		withHexPrefix(env.FUSDAddress),
	)

	return code
}
