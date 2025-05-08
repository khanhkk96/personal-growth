package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateNumberOTP(level int) (string, error) {
	otp := ""

	for i := 0; i < level; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		otp += fmt.Sprintf("%d", num.Int64())
	}

	return otp, nil
}

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func GenerateStringOTP(level int) (string, error) {
	result := make([]byte, level)

	for i := range result {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[index.Int64()]
	}

	return string(result), nil
}
