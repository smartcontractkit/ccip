package main

import (
	"strings"

	"github.com/BurntSushi/toml"
)

// TokenMap type for mapping token addresses to names
type TokenMap map[string]string

// LoadTokenMapFromToml loads the token mappings from a TOML file and converts all keys to lowercase.
func LoadTokenMapFromToml(path string) (TokenMap, error) {
	var tokenMap TokenMap
	if _, err := toml.DecodeFile(path, &tokenMap); err != nil {
		return nil, err
	}

	// Create a new map with lowercase keys
	lowerCaseTokenMap := make(TokenMap)
	for key, value := range tokenMap {
		lowerCaseTokenMap[strings.ToLower(key)] = value
	}

	return lowerCaseTokenMap, nil
}

// GetTokenName looks up the friendly name for a given token address.
func GetTokenName(tokenMap TokenMap, address string) (string, bool) {

	// downcase tokenMap[address]

	// return tokenMap[address]	

	name, exists := tokenMap[address]
	return name, exists
}

