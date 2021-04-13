package main

import (
	"fmt"
	"github.com/LampardNguyen234/go-binance/main/common"
	"github.com/LampardNguyen234/go-binance/v2"
	"golang.org/x/crypto/ssh/terminal"
)

func Init() error {
	fmt.Printf("Input the API key file: ")
	fileName, err := terminal.ReadPassword(0)
	if err != nil {
		return err
	}
	fmt.Printf("Reading API key from file %v\n", string(fileName))

	apiKey, err := LoadAPIKey(fileName)
	if err != nil {
		return err
	}
	fmt.Printf("\nLoad API key successfully!")

	common.BClient = binance.NewClient(apiKey.APIKey, apiKey.Secret)

	return nil
}
