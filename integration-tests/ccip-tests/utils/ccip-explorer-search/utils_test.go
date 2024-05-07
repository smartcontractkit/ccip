package main

import (
	"strings"
	"testing"
)

func mockConfig() *Config {
    return &Config{
        Networks: map[string]struct{
            URL    string            `toml:"url"`
            Tokens map[string]string `toml:"tokens"`
        }{
            "polygon-mainnet": {
                URL: "https://polygonscan.com/tx/",
                Tokens: map[string]string{
                    "0x123AbC": "LINK",
                },
            },
			"avalanche-mainnet": {
                URL: "https://polygonscan.com/tx/",
                Tokens: map[string]string{
                    "0x125": "LINK",
                },
            },
            "binance-smart-chain": {
                URL: "https://bscscan.com/tx/",
                Tokens: map[string]string{
                    "0xabc": "BNB",
                },
            },
        },
    }
}


func TestResolveTokenAddress(t *testing.T) {
	config := mockConfig()

	tests := []struct {
		networkName string
		feeToken    string
		expected    string
	}{
		// Test case: Known token name in known network
		{"polygon-mainnet", "LINK", "0x123abc"},
		// Test case: straight token address in known network
		{"polygon-mainnet", "0x123DeF", "0x123def"},
		// Test case: link token address in a different network
		{"avalanche-mainnet", "LINK", "0x125"},
		// Test case: link token address in a different network
		{"avalanche-mainnet", "0x125", "0x125"},
		// Test case: Unknown token name in known network but this was not going to work anyways
		{"polygon-mainnet", "ETH", "eth"},
		// Test case: Known token name in unknown network but this was not going to work anyways
		{"unknown-network", "LINK", "link"},
		// Test case: Empty network and token
		{"", "", ""},
		//Test case is unknown network and any token
		{"unknown-network", "0x129", "0x129"},
	}

	for _, test := range tests {
		result := resolveTokenAddress(test.networkName, test.feeToken, config)
		if result != test.expected {
			t.Errorf("resolveTokenAddress(%q, %q) = %q; want %q", test.networkName, test.feeToken, result, test.expected)
		}
	}
}

// Helper function to create a mock configuration for testing
func mockConfigForTokenLookup() *Config {
	return &Config{
		Networks: map[string]struct{
			URL    string            `toml:"url"`
			Tokens map[string]string `toml:"tokens"`
		}{
			"polygon-mainnet": {
				URL: "https://polygonscan.com/tx/",
				Tokens: map[string]string{
					"0x123": "LINK",
					"0x456": "WETH",
				},
			},
			"avalanche-mainnet": {
                URL: "https://polygonscan.com/tx/",
                Tokens: map[string]string{
                    "0x125": "LINK",
                },
            },
			"ethereum-mainnet": {
				URL: "https://etherscan.io/tx/",
				Tokens: map[string]string{
					"0x789": "DAI",
					"0xabc": "USDC",
				},
			},
		},
	}
}

func TestFindTokenAddressByName(t *testing.T) {
	config := mockConfigForTokenLookup()

	tests := []struct {
		networkName string
		tokenName   string
		expected    string
		expectFound bool
	}{
		{"polygon-mainnet", "LINK", "0x123", true},
		{"polygon-mainnet", "link", "0x123", true}, // Case-insensitive match
		{"ethereum-mainnet", "DAI", "0x789", true},
		{"ethereum-mainnet", "usdc", "0xabc", true}, // Case-insensitive match
		{"polygon-mainnet", "ETH", "", false},       // Token not found
		{"polygon-mainnet", "WETH", "0x456", true},      // WETH token found
		{"unknown-network", "LINK", "", false},      // Network not found
		{"avalanche-mainnet", "LINK", "0x125", true},    // Network not found
	}

	for _, test := range tests {
		address, found := findTokenAddressByName(config, test.networkName, test.tokenName)
		if found != test.expectFound {
			t.Errorf("findTokenAddressByName(%q, %q) found = %v; want %v", test.networkName, test.tokenName, found, test.expectFound)
		}
		if found && !strings.EqualFold(address, test.expected) {
			t.Errorf("findTokenAddressByName(%q, %q) = %q; want %q", test.networkName, test.tokenName, address, test.expected)
		}
	}
}

