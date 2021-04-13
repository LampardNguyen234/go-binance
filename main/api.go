package main

import (
	"encoding/json"
	"fmt"
	"github.com/LampardNguyen234/go-binance/main/common"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
	"strings"
)

type String interface {
	ToBytes() []byte
}

type APIKey struct {
	APIKey string `json:"APIKey"`
	Secret string `json:"Secret"`
}

func (apiKey APIKey) ToBytes() []byte {
	res, err := json.Marshal(apiKey)
	if err != nil {
		return []byte{}
	}

	return res
}

func (apiKey *APIKey) FromBytes(data []byte) error {
	err := json.Unmarshal(data, &apiKey)
	if err != nil {
		return err
	}

	return nil
}

func EncryptAndStore(password, filePath string, data String) error {
	if len(filePath) == 0 {
		return fmt.Errorf("file path is empty")
	}

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			f, err = os.Create(filePath)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
	defer f.Close()

	toBeStored := data.ToBytes()

	cptx, err := common.AESEncrypt([]byte(password), toBeStored)
	if err != nil {
		return fmt.Errorf("AESEncrypt returns an error: %v", err)
	}

	_, err = f.Write(cptx)

	return err
}

func StoreAPIKey(filePath string, apiKey APIKey) error {
	fmt.Printf("Please type in the encryption password (mandatory): ")
	password, err := terminal.ReadPassword(0)
	if err != nil {
		return err
	}

	return EncryptAndStore(string(password), filePath, apiKey)
}

func LoadAPIKey(filePath string) (*APIKey, error) {
	fmt.Printf("Please type in the decryption password (mandatory): ")
	password, err := terminal.ReadPassword(0)
	if len(filePath) == 0 {
		return nil, fmt.Errorf("file path is empty")
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	msg, err := common.AESDecrypt(password, data)
	if err != nil {
		return nil, fmt.Errorf("AESDecrypt returns an error: %v", err)
	}

	var apiKey APIKey
	err = apiKey.FromBytes(msg)

	return &apiKey, err
}
