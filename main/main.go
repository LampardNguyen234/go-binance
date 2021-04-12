package main

import (
	"encoding/json"
	"fmt"
	"github.com/LampardNguyen234/go-binance/main/myaes"
	"github.com/LampardNguyen234/go-binance/v2"
	"io/ioutil"
	"os"
	"strings"
)

const DEFAULT_API_KEY_FILE = "api.enc"

type APIKey struct{
	apiKey string
	secret string
}

func EncryptAndStore(password, filePath string, data interface{}) error {
	if len(filePath) == 0 {
		return fmt.Errorf("file path is empty")
	}
	f, err := os.Open(filePath)
	if err != nil{
		if strings.Contains(err.Error(), "not found") {
			f, err = os.Create(filePath)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
	defer f.Close()

	toBeStored, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal data error: %v", err)
	}

	cptx, err := myaes.AESEncrypt([]byte(password), toBeStored)
	if err != nil {
		return fmt.Errorf("AESEncrypt returns an error: %v", err)
	}

	_, err = f.Write(cptx)

	return err
}

func LoadAPIKey(password, filePath string) (*APIKey, error) {
	if len(filePath) == 0 {
		return nil, fmt.Errorf("file path is empty")
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	msg, err := myaes.AESDecrypt([]byte(password), data)
	if err != nil {
		return nil, fmt.Errorf("AESDecrypt returns an error: %v", err)
	}

	var apiKey APIKey
	err = json.Unmarshal(msg, &apiKey)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal API key from msg %v: %v", string(msg), err)
	}

	return &apiKey, nil
}

func main() {
	bClient := new(binance.Client)

}