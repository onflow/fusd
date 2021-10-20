package contracts_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/onflow/fusd/lib/go/contracts"
)

const addrA = "0x0A"

func TestFUSDContract(t *testing.T) {
	contract := contracts.FUSD(addrA)
	assert.NotNil(t, contract)
	assert.Contains(t, string(contract), addrA)
}
