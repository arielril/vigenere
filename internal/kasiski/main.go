package kasiski

import "fmt"

const (
	minRepetitionSize = 1
	maxRepetitionSize = 10
)

func GetPossibleKeyLengths(msg string) interface{} {
	// get all possible repetitions with size 1 to 10
	_ = getPossibleRepetitions(msg)

	return nil
}

func getPossibleRepetitions(msg string) interface{} {
	for seqSize := minRepetitionSize; seqSize <= maxRepetitionSize; seqSize++ {
		for seqStart := 0; seqStart <= len(msg)-seqSize; seqStart++ {
			seq := msg[seqStart : seqStart+seqSize]

			fmt.Println("seq=", seq)
		}
	}
	return nil
}
