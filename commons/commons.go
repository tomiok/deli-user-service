package commons

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
)

func StringUUID() string {
	return uuid.New().String()
}

func EncryptPass(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	sha256Hash := hex.EncodeToString(h.Sum(nil))
	return sha256Hash
}