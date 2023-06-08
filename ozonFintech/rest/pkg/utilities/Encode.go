package utilities

import (
	"crypto/md5"
	"encoding/hex"
	"math/big"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
const length = 10

func EncodeBase63(n int64) string {
	if n == 0 {
		return string(alphabet[0])
	}
	result := make([]byte, 0, length)
	for n > 0 && len(result) < length {
		result = append(result, alphabet[n%63])
		n /= 63
	}
	for len(result) < length {
		result = append(result, alphabet[0])
	}
	reverse(result)
	return string(result)
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func HashLink(link string) int64 {
	h := md5.New()
	bi := big.NewInt(0)
	h.Write([]byte(link))
	hashBytes := h.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	bi.SetString(hashString, 16)
	return bi.Int64()
}
