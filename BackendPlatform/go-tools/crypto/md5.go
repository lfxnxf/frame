package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 字符串md5
func Md5(s string) string {
	h := md5.New()
	n, err := h.Write([]byte(s))
	if err != nil || n == 0 {
		return ""
	}
	cipherStr := h.Sum(nil)
	sign := hex.EncodeToString(cipherStr)
	return sign
}

// Md5B byte
func Md5B(s []byte) string {
	h := md5.New()
	n, err := h.Write(s)
	if err != nil || n == 0 {
		return ""
	}
	cipherStr := h.Sum(nil)
	sign := hex.EncodeToString(cipherStr)
	return sign
}
