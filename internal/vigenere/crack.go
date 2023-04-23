package vigenere

import (
	"github.com/arielril/vigenere/internal/frequency"
	"github.com/projectdiscovery/gologger"
)

func Crack(msg string) (string, error) {
	gologger.Silent().Msgf("message length: %d\n", len(msg))

	m := make([]int, len(msg))
	for i := 0; i < len(msg); i++ {
		m[i] = int(msg[i] - 'a')
	}

	bestFit := 1e100
	bestKey := ""

	for i := 1; i <= 26; i++ {

		key := make([]byte, i)

		fit := getFrequencyEveryNthPartition(m, key)

		if fit < bestFit {
			bestFit = fit
			bestKey = string(key)
			gologger.Debug().Msgf("best key so far: %s\n", bestKey)
		}
	}

	gologger.Silent().Msgf("found key: %s\n", bestKey)

	decodedMessage, err := Decode(msg, bestKey)
	return decodedMessage, err
}

func sumSlice(a []float64) (sum float64) {
	for _, f := range a {
		sum += f
	}
	return
}

func getKeyFrequency(letterCount []float64) int {
	sum := sumSlice(letterCount)
	bestFit := 1e100
	bestFrequency := 0

	for rotate := 0; rotate < 26; rotate++ {
		fit := 0.0
		for i := 0; i < 26; i++ {
			d := letterCount[(i+rotate)%26]/sum - frequency.EnglishLetterFrequency[i]
			fit += d * d / frequency.EnglishLetterFrequency[i]
		}

		if fit < bestFit {
			bestFit = fit
			bestFrequency = rotate
		}
	}

	return bestFrequency
}

func getFrequencyEveryNthPartition(msg []int, key []byte) float64 {
	messageLength := len(msg)
	keyLength := len(key)
	letterCount := make([]float64, 26)
	letterFrequency := make([]float64, 26)

	for j := 0; j < keyLength; j++ {

		for k := 0; k < 26; k++ {
			letterCount[k] = 0.0
		}

		for i := j; i < messageLength; i += keyLength {
			letterCount[msg[i]]++
		}

		keyFrequency := getKeyFrequency(letterCount)
		key[j] = byte(keyFrequency + 'a')
		for i := 0; i < 26; i++ {
			letterFrequency[i] += letterCount[(i+keyFrequency)%26]
		}
	}

	sum := sumSlice(letterFrequency)
	freq := 0.0

	for i := 0; i < 26; i++ {
		d := letterFrequency[i]/sum - frequency.EnglishLetterFrequency[i]
		freq += d * d / frequency.EnglishLetterFrequency[i]
	}

	return freq
}
