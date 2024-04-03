package main

import (
	"strings"
)

// Function to construct scan URL using the TOML configuration
func constructScanURL(networkName string, transactionHash string, config *Config) string {
	networkConfig, ok := config.Networks[networkName]
	if !ok {
		return "" // Return an empty string for unknown or unsupported networks
	}
	return networkConfig.URL + transactionHash
}

func constructMessageURL(messageID string) string {
	// Construct the message URL using the TOML configuration
	return "https://ccip.chain.link/msg/" + messageID
}

func findTokenNameByAddressInsensitive(tokens map[string]string, address string) (string, bool) {
	for tokenAddress, tokenName := range tokens {
		if strings.EqualFold(tokenAddress, address) {
			return tokenName, true
		}
	}
	return "", false
}

// Reverse lookup to find a token address by its name in a case-insensitive manner
func findTokenAddressByName(networkConfig *Config, networkName, tokenName string) (string, bool) {
	network, exists := networkConfig.Networks[networkName]
	if !exists {
		return "", false // Network not found
	}

	for address, name := range network.Tokens {
		if strings.EqualFold(name, tokenName) {
			return address, true
		}
	}

	return "", false // Token name not found
}

func resolveTokenAddress(networkName, feeToken string, config *Config) string {
	// Check if the feeToken is a known token name and a network name is provided
	if networkName != "" && feeToken != "" {
		// Convert token name to address if possible
		if address, found := findTokenAddressByName(config, networkName, feeToken); found {
			return strings.ToLower(address) // Token name successfully resolved to an address
		}
	}
	// Return the feeToken as is if it's not a recognized token name or no network name is provided
	return strings.ToLower(feeToken)
}


