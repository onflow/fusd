package contracts

import (
	"fmt"
	"strings"

	"github.com/onflow/fusd/lib/go/contracts/internal/assets"
)

const (
	fusdFilename                    = "FUSD.cdc"
	placeholderFungibleTokenAddress = "0xFUNGIBLETOKENADDRESS"
)

func FUSD(fungibleTokenAddress string) []byte {
	code := assets.MustAssetString(fusdFilename)

	// Replace the fungible token address
	code = strings.ReplaceAll(
		code,
		placeholderFungibleTokenAddress,
		withHexPrefix(fungibleTokenAddress),
	)

	return []byte(code)
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
