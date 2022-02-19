package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func CreateAndCompareSHA256(data []byte, expect string) bool {
	hashed := sha256.Sum256(data)
	result := hashed[:]
	return hex.EncodeToString(result) == expect
}
