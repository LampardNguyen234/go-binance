package common

import (
	"fmt"
	"math/rand"
)

func RandInt() int {
	return rand.Int()
}

func RandInt64() int64 {
	return rand.Int63()
}

func RandBytes(length int) []byte {
	if length == 0 {
		return nil
	}

	res := make([]byte, length)

	_, err := rand.Read(res)
	if err != nil {
		return nil
	}

	return res
}

func RandPrintable(length int) string {
	res := ""
	for len(res) < length {
		r := 48 + RandInt() % 75
		res += fmt.Sprintf("%c", r)
	}

	return res
}
