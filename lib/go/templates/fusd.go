package templates

import "github.com/onflow/fusd/lib/go/templates/internal/assets"

const (
	setupFUSDVaultFilename = "setup_fusd_vault.cdc"
	transferFUSDFilename   = "transfer_fusd.cdc"
)

func SetupFUSDVaultTransaction(env Environment) []byte {
	code := assets.MustAssetString(setupFUSDVaultFilename)

	return []byte(replaceAddresses(code, env))
}

func TransferFUSDTransaction(env Environment) []byte {
	code := assets.MustAssetString(transferFUSDFilename)

	return []byte(replaceAddresses(code, env))
}
