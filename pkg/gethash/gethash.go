package gethash

import (
	"crypto/sha256"
	"encoding/hex"
)

// Функция по созданию hash для отправки пользователю
func GetHash(toEmail string) string {
	hash := sha256.Sum256([]byte(toEmail))
	hexHash := hex.EncodeToString(hash[:])
	return hexHash
}
