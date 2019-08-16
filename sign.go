package hpool

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func Sign(secretKey, query string) string {
	// add secrete key
	query += "&secret_key=" + secretKey
	sum := sha256.Sum256([]byte(query))
	hexed := hex.EncodeToString(sum[:])
	return strings.ToUpper(hexed)
}
