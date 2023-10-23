package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	overrideconfigPath := flag.String("path", "./override.toml", "config for overriding for default test config")
	flag.Parse()

	if *overrideconfigPath == "" {
		log.Println("override.toml path is required")
		os.Exit(1)
	}
	cData, err := os.ReadFile(*overrideconfigPath)
	if err != nil {
		log.Println("unable to read the toml at ", *overrideconfigPath, "error - ", err)
		os.Exit(1)
	}
	// convert the data to Base64 encoded string
	encoded := base64.StdEncoding.EncodeToString(cData)
	// set the env var
	if os.Setenv("BASE64_TEST_CONFIG_OVERRIDE", encoded) != nil {
		os.Exit(1)
	}
	fmt.Println("Successfully set the env var BASE64_TEST_CONFIG_OVERRIDE with the contents of ", *overrideconfigPath, "as Base64 encoded string")
}
