package util

import (
	"crypto/md5"
	"encoding/hex"
)


func HashString(s string) string {
	bytes := []byte(s)
	h := md5.New()
	h.Write(bytes)
	hashedBytes := h.Sum(nil)
	return hex.EncodeToString(hashedBytes)
}