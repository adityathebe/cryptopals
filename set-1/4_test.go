package cryptopals

import (
	"testing"
)

func TestDetectSingleCharXOR(t *testing.T) {
	answer := detectSingleCharacterXOR("./4.txt")
	if answer != "Now that the party is jumping\n" {
		t.Error("Did not match the expected")
	}
}
