package kasiski

import (
	"sort"

	"github.com/arielril/vigenere/pair"
	"github.com/projectdiscovery/gologger"
)

const (
	minRepetitionSize = 1
	maxRepetitionSize = 10
	maxKeySize        = 16
)

// GetPossibleKeyLengths returns the most probable
// key lengths
func GetPossibleKeyLengths(msg string) []int {
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

			for i := 0; i < len(msg)-seqSize; i++ {
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
