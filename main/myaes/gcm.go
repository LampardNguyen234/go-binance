package myaes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

func AESEncrypt(password, msg []byte) ([]byte, error) {
	if len(password) < 8 {
		return nil, fmt.Errorf("length of password must be at least 8, got %v", len(password))
	}

	if len(msg) == 0 {
		return nil, fmt.Errorf("msg is empty")
	}

	key := sha256.Sum256(password)

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte,gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	cptx := gcm.Seal(nonce, nonce, msg, nil)

	return cptx, nil
}

func AESDecrypt(password, data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("cipher text is empty")
	}

	key := sha256.Sum256(password)

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce, cptx := data[:gcm.NonceSize()], data[gcm.NonceSize():]

	msg, err := gcm.Open(nil, nonce, cptx, nil)

	return msg, err
}