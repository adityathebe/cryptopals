package cryptopals

import (
	"encoding/hex"
	"io/ioutil"
)

func singleXOR(org []byte, x byte) []byte {
	for i := range org {
		org[i] = org[i] ^ x
	}
	return org
}

func measureScore(input []byte, scoreRef map[rune]float64) float64 {
	var score float64 = 0
	for _, char := range string(input) {
		score += scoreRef[char]
	}
	return score
}

func buildScoreFromCorpus(sourceFile string) map[rune]float64 {
	fileContent, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		panic(err)
	}
	score := make(map[rune]float64)
	for _, char := range string(fileContent) {
		score[char] = score[char] + 1
	}

	// Normalize
	var totalScore float64
	for x := range score {
		totalScore += score[x]
	}

	for x := range score {
		score[x] = score[x] / totalScore
	}

	return score
}

// SingleByteXOR a
func SingleByteXOR(hexStr string) (string, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	scoreRef := buildScoreFromCorpus("./test-data.txt")
	var bestScore float64 = 0
	var bestOne string
	for i := 0; i < 256; i++ {
		xored := singleXOR(bytes, byte(i))
		score := measureScore(xored, scoreRef)
		if score > bestScore {
			bestScore = score
			bestOne = string(xored)
		}
	}
	return bestOne, nil
}
