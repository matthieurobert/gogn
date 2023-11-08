package utils

import (
	"crypto/rand"
	"log"
	"math/big"
)

func GenerateRandomString(length int) string {
	possible := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	values := make([]byte, length)
	_, err := rand.Read(values)

	if err != nil {
		log.Fatal(err)
	}

	var result string
	for _, v := range values {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(possible))))
		result += string(possible[int(v)%int(index.Int64())])
	}
	return result
}
