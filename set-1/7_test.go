package cryptopals

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestDecryptAESECB(t *testing.T) {
	file, err := ioutil.ReadFile("./7.txt")
	if err != nil {
		panic(err)
	}
	cipher, err := base64.StdEncoding.DecodeString(string(file))
	if err != nil {
		panic(err)
	}
	res := decryptAESECB(cipher, []byte("YELLOW SUBMARINE"))
	fmt.Println(string(res))
}
