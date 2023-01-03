package config

import "os"

type ChainConfig struct {
	PublicMainnetURL string
	PublicTestnetURL string
	MainnetChainID   string
	TestnetChainID   string
}

// database - all DB variables
func chain() *ChainConfig {
	config := &ChainConfig{}

	env()

	config.PublicMainnetURL = os.Getenv("PUBLIC_MAINNET_URL")
	config.PublicTestnetURL = os.Getenv("PUBLIC_TESTNET_URL")
	config.MainnetChainID = os.Getenv("MAINNET_CHAIN_ID")
	config.TestnetChainID = os.Getenv("TESTNET_CHAIN_ID")

	return config
}
