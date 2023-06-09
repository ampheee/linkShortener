package utilities

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
const length = 10

func EncodeBase63(n uint64) string {
	if n == 0 {
		return strings.Repeat(string(alphabet[0]), length)
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

func HashLink(link string) uint64 {
	//log := logger.GetLogger()
	h := sha256.New()
	bi := big.NewInt(0)
	h.Write([]byte(link))
	hashBytes := h.Sum(nil)
	check := bi.SetBytes([]byte(hex.EncodeToString(hashBytes)))
	return check.Uint64()
}
