package cryptopals

import (
	"crypto/aes"
	"log"
)

func decryptAESECB(cipher, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	out := make([]byte, len(cipher))
	if len(cipher)%block.BlockSize() != 0 {
		log.Fatalf("Invalid Length %d", len(cipher))
	}

	for i := 0; i < len(cipher); i += block.BlockSize() {
		block.Decrypt(out[i:], cipher[i:])
	}
	return out
}
