package kzcrypto

import (
	"crypto/rand"
	"math/big"
)

const Letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"

func GenerateRandomString(length int) (string, error) {
	secret := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(Letters))))
		if err != nil {
			return "", err
		}
		secret[i] = Letters[num.Int64()]
	}

	return string(secret), nil
}
