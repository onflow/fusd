package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/onflow/flow-go-sdk"

	"github.com/onflow/fusd/lib/go/templates"
)

type manifest struct {
	Network   string     `json:"network"`
	Templates []template `json:"templates"`
}

func (m *manifest) addTemplate(t template) {
	m.Templates = append(m.Templates, t)
}

type template struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Source    string     `json:"source"`
	Arguments []argument `json:"arguments"`
	Network   string     `json:"network"`
	Hash      string     `json:"hash"`
}

type argument struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	SampleValue string `json:"sampleValue"`
}

type templateGenerator func(env templates.Environment) []byte

func generateTemplate(
	id, name string,
	env templates.Environment,
	generator templateGenerator,
	arguments []argument,
) template {
	source := generator(env)

	h := sha256.New()
	h.Write(source)
	hash := h.Sum(nil)

	return template{
		ID:        id,
		Name:      name,
		Source:    string(source),
		Arguments: arguments,
		Network:   env.Network,
		Hash:      hex.EncodeToString(hash),
	}
}

func generateManifest(env templates.Environment) *manifest {
	m := &manifest{
		Network: env.Network,
	}

	sampleAmount := "92233720368.54775808"

	m.addTemplate(generateTemplate(
		"FUSD.01", "Setup FUSD Vault",
		env,
		templates.SetupFUSDVaultTransaction,
		[]argument{},
	))

	m.addTemplate(generateTemplate(
		"FUSD.02", "Transfer FUSD",
		env,
		templates.TransferFUSDTransaction,
		[]argument{
			{
				Type:        "UFix64",
				Name:        "amount",
				Label:       "Amount",
				SampleValue: sampleAmount,
			},
			{
				Type:        "Address",
				Name:        "recipient",
				Label:       "Recipient",
				SampleValue: sampleAddress(env.Network),
			},
		},
	))

	return m
}

func sampleAddress(network string) string {
	var address flow.Address

	switch network {
	case testnet:
		address = flow.NewAddressGenerator(flow.Testnet).NextAddress()
	case mainnet:
		address = flow.NewAddressGenerator(flow.Mainnet).NextAddress()
	}

	return fmt.Sprintf("0x%s", address.Hex())
}
