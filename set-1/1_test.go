package cryptopals

import "testing"

func TestAbs(t *testing.T) {
	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	b64Str, err := HexToBase64(hexStr)
	if err != nil {
		t.Fatal(err)
	}
	if b64Str != expected {
		t.Errorf("Base64 strings didn't match")
	}
}
