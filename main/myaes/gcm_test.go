package myaes

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"testing"
)

func TestAESDecrypt(t *testing.T) {
	msg := []byte("Hello world!!! My name is AES GCM")
	fmt.Println("msg", msg)

	numTries := 10

	for i := 0; i < numTries; i++ {
		password := make([]byte, 32)
		_, err := rand.Read(password)
		if err != nil {
			panic(err)
		}

		cptx, err := AESEncrypt(password, msg)
		if err != nil {
			panic(err)
		}

		recoveredMsg, err := AESDecrypt(password, cptx)
		if err != nil {
			panic(err)
		}

		if !bytes.Equal(recoveredMsg, msg) {
			panic(fmt.Sprintf("recoveredMsg and msg mismatch: %v, %v", recoveredMsg, msg))
		}
	}
}
