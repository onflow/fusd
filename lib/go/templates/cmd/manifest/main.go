package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/psiemens/sconfig"
	"github.com/spf13/cobra"

	"github.com/onflow/fusd/lib/go/templates"
)

type Config struct {
	Network string `default:"mainnet" flag:"network" info:"Flow network to generate for"`
}

const envPrefix = "FLOW"

const (
	testnet = "testnet"
	mainnet = "mainnet"
)

const (
	testnetFungibleTokenAddress = "9a0766d93b6608b7"
	testnetFUSDAddress          = "e223d8a629e49c68"
)

const (
	mainnetFungibleTokenAddress = "f233dcee88fe0abe"
	mainnetFUSDAddress          = "3c5959b568896393"
)

var conf Config

var cmd = &cobra.Command{
	Use:   "manifest <outfile>",
	Short: "Generate a JSON manifest of all FUSD transaction templates",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		env, err := getEnv(conf)
		if err != nil {
			exit(err)
		}

		manifest := generateManifest(env)

		b, err := json.MarshalIndent(manifest, "", "  ")
		if err != nil {
			exit(err)
		}

		outfile := args[0]

		err = ioutil.WriteFile(outfile, b, 0777)
		if err != nil {
			exit(err)
		}
	},
}

func getEnv(conf Config) (templates.Environment, error) {

	if conf.Network == testnet {
		return templates.Environment{
			Network:              testnet,
			FungibleTokenAddress: testnetFungibleTokenAddress,
			FUSDAddress:          testnetFUSDAddress,
		}, nil
	}

	if conf.Network == mainnet {
		return templates.Environment{
			Network:              mainnet,
			FungibleTokenAddress: mainnetFungibleTokenAddress,
			FUSDAddress:          mainnetFUSDAddress,
		}, nil
	}

	return templates.Environment{}, fmt.Errorf("invalid network %s", conf.Network)
}

func init() {
	initConfig()
}

func initConfig() {
	err := sconfig.New(&conf).
		FromEnvironment(envPrefix).
		BindFlags(cmd.PersistentFlags()).
		Parse()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := cmd.Execute(); err != nil {
		exit(err)
	}
}

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
