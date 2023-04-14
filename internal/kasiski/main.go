package kasiski

import (
	"strings"
	"sync"

	"github.com/projectdiscovery/gologger"
)

const (
	minRepetitionSize = 1
	maxRepetitionSize = 20
	maxKeySize        = 16
)

func GetPossibleKeyLength(msg string) int {
	gologger.Debug().Msg("[possible-key-length] starting the discovery of the key length")
	max := 0
	indexOfCoincidence := make([]float64, maxRepetitionSize)

	var idxOfCoincidenceWg sync.WaitGroup

	for splitSize := minRepetitionSize; splitSize <= maxRepetitionSize; splitSize++ {

		idxOfCoincidenceWg.Add(1)
		go func(msg string, splitSize int, indexOfCoincidence []float64) {
			coincidenceIdx := computeIndexOfCoincidence(msg, splitSize)
			indexOfCoincidence[splitSize-1] = coincidenceIdx
			idxOfCoincidenceWg.Done()
		}(msg, splitSize, indexOfCoincidence)

	}
	idxOfCoincidenceWg.Wait()

	for i := 0; i < maxRepetitionSize; i++ {
		if indexOfCoincidence[i] > 0.06 {
			max = i
			break
		} else if indexOfCoincidence[i] > float64(indexOfCoincidence[max]) {
			max = i
		}

	}
	return max + 1
}

func computeIndexOfCoincidence(msg string, splitSize int) float64 {
	gologger.Debug().Msgf("[possible-key-length] current split size: %d\n", splitSize)

	letterFrequencies := make([]float64, splitSize)

	indexOfCoincidence := 0.0

	for startIdx := 0; startIdx < splitSize; startIdx++ {

		currentLetterSequence := ""
		sum := startIdx

		for sum < len(msg) {
			currentLetterSequence += string(msg[sum])
			sum += splitSize
		}

		letterSequenceLen := len(currentLetterSequence)
		gologger.Debug().Msgf("[possible-key-length] current letter sequence length: %d\n", letterSequenceLen)

		for i := 0; i < 26; i++ {
			letter := string(i + 'a')
			letterCountInLetterSequence := strings.Count(currentLetterSequence, letter)
			letterFrequency := float64(letterCountInLetterSequence) / float64(letterSequenceLen)
			frequencyInd := letterFrequency * (float64(letterCountInLetterSequence-1) / float64(letterSequenceLen-1))
			letterFrequencies[startIdx] += frequencyInd
		}

		indexOfCoincidence += letterFrequencies[startIdx]
	}

	return indexOfCoincidence / float64(splitSize)
}
