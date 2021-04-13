package main

import (
	"bytes"
	"fmt"
	"github.com/LampardNguyen234/go-binance/main/common"
	"testing"
)

func TestLoadAPIKey(t *testing.T) {
	numTries := 1000
	for i := 0; i < numTries; i++ {
		apiKey := APIKey{
			APIKey: common.RandPrintable(common.RandInt()%100),
			Secret: common.RandPrintable(common.RandInt()%100),
		}

		password := string(common.RandBytes(32))

		fileName := fmt.Sprintf("%v.bin", i)

		err := EncryptAndStore(password, fileName, apiKey)
		if err != nil {
			panic(err)
		}

		newAPIKey, err := LoadAPIKey(password, fileName)
		if err != nil {
			panic(err)
		}

		if !bytes.Equal(apiKey.ToBytes(), newAPIKey.ToBytes()) {
			fmt.Println(apiKey.ToBytes())
			fmt.Println(newAPIKey.ToBytes())
			panic("api and newAPI mismatch")
		}
	}

}
