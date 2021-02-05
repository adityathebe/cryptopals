package cryptopals

import "encoding/hex"

func repeatingKeyXOR(dataToEncrypt, key []byte) string {
	res := make([]byte, len(dataToEncrypt))
	keyIndex := 0
	for i := range dataToEncrypt {
		res[i] = dataToEncrypt[i] ^ key[keyIndex]
		keyIndex = (keyIndex + 1) % len(key)
	}
	return hex.EncodeToString(res)
}
