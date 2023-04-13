package kasiski

import (
	"sort"
	"strings"
	"sync"

	"github.com/arielril/vigenere/pair"
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
		/*
			gologger.Debug().Msgf("[possible-key-length] current split size: %d\n", splitSize)

			letterFrequencies := make([]float64, splitSize)

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

				indexOfCoincidence[splitSize-1] += letterFrequencies[startIdx]
			}

			indexOfCoincidence[splitSize-1] /= float64(splitSize)
		*/
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

// KasiskiExamination returns the most probable
// key lengths. Kasiski examination
func KasiskiExamination(msg string) []int {
	// get all possible repetitions with size 1 to 10
	repeatedSequencesMap := getPossibleRepetitions(msg)

	sequenceFactors := make(map[string][]int, 0)
	for sequence, sequenceSpacing := range repeatedSequencesMap {

		sequenceFactors[sequence] = make([]int, 0)

		for _, spacing := range sequenceSpacing {
			sequenceFactors[sequence] = append(sequenceFactors[sequence], getFactors(spacing)...)
		}
	}

	mostCommonFactors := getMostCommonFactors(sequenceFactors)
	sort.Sort(mostCommonFactors)

	mostProbableKeyLengths := make([]int, 0)
	for _, v := range mostCommonFactors {
		mostProbableKeyLengths = append(mostProbableKeyLengths, v.Key)
	}

	return mostProbableKeyLengths
}

func getPossibleRepetitions(msg string) map[string][]int {
	sequenceMapping := make(map[string][]int, 0)

	for seqSize := minRepetitionSize; seqSize <= maxRepetitionSize; seqSize++ {
		for seqStart := 0; seqStart <= len(msg)-seqSize; seqStart++ {

			seq := msg[seqStart : seqStart+seqSize]
			gologger.Debug().Msgf("seq= %s\n", seq)

			for i := seqStart + seqSize; i < len(msg)-seqSize; i++ {
				if msg[i:i+seqSize] == seq {
					if sequenceMapping[seq] == nil {
						sequenceMapping[seq] = make([]int, 0)
					}

					sequenceMapping[seq] = append(sequenceMapping[seq], i-seqStart)
				}
			}
		}
	}

	return sequenceMapping
}

func getFactors(value int) []int {
	factors := make([]int, 0)
	if value < 2 {
		return factors
	}

	for i := 2; i <= maxKeySize+1; i++ {
		if value%i == 0 {
			if i != 1 {
				factors = append(factors, i)
			}
			if value/i != 1 {
				factors = append(factors, value/i)
			}
		}
	}

	return factors
}

func getMostCommonFactors(sequenceFactors map[string][]int) pair.PairList {
	mostCommonFactors := make(map[int]int, 0)

	for _, factors := range sequenceFactors {
		for _, factor := range factors {
			mostCommonFactors[factor] += 1
		}
	}

	filteredCommonFactors := make(map[int]int, 0)
	for factor, count := range mostCommonFactors {
		if factor <= maxKeySize {
			filteredCommonFactors[factor] = count
		}
	}

	return pair.NewPairListFromMap(filteredCommonFactors)
}
