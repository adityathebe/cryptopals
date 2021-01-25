package cryptopals

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestXOR(t *testing.T) {
	hexA, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	hexB, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	expected, _ := hex.DecodeString("746865206b696420646f6e277420706c6179")

	res, err := XOR(hexA, hexB)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(res, expected) {
		t.Error("Mismatch")
	}
}
