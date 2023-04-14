package kasiski

import (
	"fmt"

	"github.com/arielril/vigenere/internal/frequency"
	"github.com/projectdiscovery/gologger"
)

func GetPossibleKeys(possibleKeyLength int, msg string) []string {
	possibleKeys := transpose(msg, possibleKeyLength)

	return possibleKeys
}

func transpose(msg string, keyLen int) []string {
	fmt.Println("transpose key len", keyLen)
	keys := make([]string, 0)

	for i := 0; i < keyLen; i++ {
		for j := i; j < len(msg)-1; j += keyLen {
			key := msg[i : i+keyLen]
			keys = append(keys, key)
			gologger.Debug().Msgf("[transpose] possible key: %s\n", key)
		}
	}

	return keys
}

func GetPossibleKey(keyLen int, msg string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	key := ""

	for _, col := range partition(msg, keyLen) {
		scores := make(map[rune]float64, 0)

		for _, letter := range alphabet {

			transposed := make([]rune, 0)
			for _, c := range col {
				cIdx := (c - 'a') % 26
				transposed = append(transposed, rune(alphabet[cIdx]))
			}

			scores[letter] = frequencyScore(transposed)
		}

		var k rune
		maxV := 0.0
		for r, v := range scores {
			if v > maxV {
				maxV = v
				k = r
			}
		}
		key += string(k)

	}

	return key
}

func partition(txt string, num int) []string {
	cols := make([]string, num)

	for i, c := range txt {
		cols[i%num] += string(c)
	}
	return cols
}

func frequencyScore(tr []rune) float64 {
	score := 0.0
	for _, c := range tr {
		score += frequency.EnglishLetterFrequency[(c-'a')%26]
	}

	return score
}
