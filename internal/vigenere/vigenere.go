package vigenere

import (
	"fmt"
	"strings"

	"github.com/arielril/vigenere/internal/frequency"
	"github.com/projectdiscovery/gologger"
)

// Encode encodes a msg with a key using the vigenere cipher
func Encode(msg, key string) (string, error) {
	// encryption formula
	// E_i = (P_i + K_i) % 26

	cipheredMessage := make([]string, len(msg))

	for i := 0; i < len(msg); i++ {
		pI := msg[i] - 'a'
		kI := key[i%len(key)] - 'a'
		eI := (pI + kI) % 26

		gologger.Debug().Msgf("P_i=%v\n", pI)
		gologger.Debug().Msgf("K_i=%v\n", kI)
		gologger.Debug().Msgf("E_i=%v\n", eI)
		gologger.Debug().Msgf("----------\n")

		cipheredMessage[i] = fmt.Sprintf("%c", eI+'a')
	}

	return strings.Join(cipheredMessage, ""), nil
}

// Decode decodes an encrypted msg with a key using the vigenere cipher
func Decode(msg, key string) (string, error) {
	// decryption formula
	// D_i = (E_i - K_i) % 26

	decipheredMessage := make([]string, len(msg))
	for i := 0; i < len(msg); i++ {
		eI := msg[i] - 'a'
		kI := key[i%len(key)] - 'a'
		dI := (eI - kI + 26) % 26

		// gologger.Debug().Msgf("E_i=%v\n", eI)
		// gologger.Debug().Msgf("K_i=%v\n", kI)
		// gologger.Debug().Msgf("D_i=%v\n", dI)
		// gologger.Debug().Msg("----------\n")

		decipheredMessage[i] = fmt.Sprintf("%c", dI+'a')
	}

	return strings.Join(decipheredMessage, ""), nil
}

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

func getKeyFrequency(a []float64) int {
	sum := sumSlice(a)
	bestFit := 1e100
	bestFrequency := 0

	for rotate := 0; rotate < 26; rotate++ {
		fit := 0.0
		for i := 0; i < 26; i++ {
			d := a[(i+rotate)%26]/sum - frequency.EnglishLetterFrequency[i]
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
