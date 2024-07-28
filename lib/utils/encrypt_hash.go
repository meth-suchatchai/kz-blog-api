package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptedHash(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}
