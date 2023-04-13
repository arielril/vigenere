package kasiski

import (
	"math"
	"strings"

	"github.com/arielril/vigenere/internal/frequency"
	"github.com/projectdiscovery/gologger"
)

func GetPossibleKey(possibleKeyLength int, msg string) string {
	gologger.Debug().Msg("[possible-key] starting the discovery of the possible key")
	possibleKey := ""

	for i := 0; i < possibleKeyLength; i++ {
		currentLetterSequence := ""
		sum := i
		for sum < len(msg) {
			currentLetterSequence += string(msg[sum])
			sum += possibleKeyLength
		}
		gologger.Debug().Msgf("[possible-key] current letter sequence: %s\n", currentLetterSequence)

		currentLetterSequenceLen := len(currentLetterSequence)
		gologger.Debug().Msgf("[possible-key] current letter sequence length: %d\n", currentLetterSequenceLen)

		chiSquareCollection := make([]float64, 26)
		for j := 0; j < 26; j++ {

			shiftedLetterSequence := ""

			for k := 0; k < currentLetterSequenceLen; k++ {
				currentLetter := currentLetterSequence[k]
				currentLetter = (currentLetter - byte(j) - 'a') % 26
				currentLetter += 'a'

				shiftedLetterSequence += string(currentLetter)
			}

			totalChiAmount := 0.0

			for l := 0; l < 26; l++ {
				letter := string(l + 'a')

				observedCount := strings.Count(shiftedLetterSequence, letter)
				expected := frequency.EnglishLetterFrequency[l] * float64(currentLetterSequenceLen)
				squaredDifference := math.Pow(float64(observedCount)-expected, 2)
				currentChi := squaredDifference / expected
				totalChiAmount += currentChi
			}

			chiSquareCollection[j] += totalChiAmount
		}

		min := 0

		for m := 0; m < 26; m++ {
			if chiSquareCollection[m] < chiSquareCollection[min] {
				min = m
			}
		}

		possibleKey += string(min + 'a')
	}

	return possibleKey
}

// func AttemptToCrackWithKeyLength(msg string, keyLength int) string {
// 	return ""
// }

// // GetNthSubKeyLetters return every nth letter for each keyLength
// func GetNthSubKeyLetters(startPosition, keyLength int, msg string) string {
// 	letters := make([]string, 0)
// 	for i := startPosition - 1; i < len(msg); {
// 		letters = append(letters, string(msg[i]))
// 		i += keyLength
// 	}

// 	return strings.Join(letters, "")
// }
