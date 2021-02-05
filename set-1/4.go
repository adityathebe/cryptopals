package cryptopals

import (
	"bufio"
	"encoding/hex"
	"os"
	"strings"
)

func detectSingleCharacterXOR(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var bestScore float64 = 0
	var answer string
	scanner := bufio.NewScanner(file)
	scoreRef := buildScoreFromCorpus("./test-data.txt")
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		x, score, _ := singleCharacterXOR(text, scoreRef)
		if score > bestScore {
			bestScore = score
			answer = x
		}
	}
	return answer
}

func singleCharacterXOR(hexStr string, scoreRef map[rune]float64) (string, float64, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", 0, err
	}
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
	return bestOne, bestScore, nil
}
