package cryptopals

import (
	"log"
	"testing"
)

func TestSingleByteXOR(t *testing.T) {
	result, err := SingleByteXOR("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		log.Fatalln(err)
	}
	if result != "Cooking MC's like a pound of bacon" {
		t.Fail()
	}
}
