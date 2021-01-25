package cryptopals

import (
	"errors"
)

// Fixed XOR
// Write a function that takes two equal-length buffers and produces their XOR combination.
// If your function works properly, then when you feed it the string:

// XOR takes two equal-length buffers and produces their XOR combination
func XOR(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("Length mismatch")
	}
	result := make([]byte, len(a))
	for i := range result {
		result[i] = a[i] ^ b[i]
	}
	return result, nil
}
