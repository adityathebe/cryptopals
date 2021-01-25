package cryptopals

// Convert hex to base64

import (
	"encoding/base64"
	"encoding/hex"
)

// HexToBase64 converts hex string into base64 string
func HexToBase64(hexStr string) (string, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}
