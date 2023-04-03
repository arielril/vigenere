package frequency

import (
	"fmt"
	"sort"

	"github.com/arielril/vigenere/pair"
	"golang.org/x/exp/slices"
)

func GetEnglishFrequencyScore(msg string) int {
	frequencyOrder := getFrequencyOrder(msg)

	matchScore := 0
	commonLetterSequence := "etaoinshrdlcumwfgypbvkjxqz"

	for _, commonLetter := range commonLetterSequence[:6] {
		if slices.Contains[string](frequencyOrder, string(commonLetter)) {
			matchScore += 1
		}
	}

	for _, uncommonLetter := range commonLetterSequence[len(commonLetterSequence)-6:] {
		if slices.Contains[string](frequencyOrder, string(uncommonLetter)) {
			matchScore += 1
		}
	}

	return matchScore
}

// getLetterCount returns the count of each letter in the message
func getLetterCount(msg string) map[rune]int {
	letterCount := make(map[rune]int, 0)

	for i := 'a'; i < 'a'+26; i++ {
		letterCount[i] = 0
	}

	for _, r := range msg {
		letterCount[r] += 1
	}

	return letterCount
}

func getFrequencyOrder(msg string) []string {
	letterCount := getLetterCount(msg)

	frequencyLetterSliceMap := make(map[int][]rune, 0)
	for letter, count := range letterCount {
		if frequencyLetterSliceMap[count] == nil {
			frequencyLetterSliceMap[count] = make([]rune, 0)
		}

		frequencyLetterSliceMap[count] = append(frequencyLetterSliceMap[count], letter)
	}

	frequencyOrdered := pair.NewPairListRuleSliceFromMap(frequencyLetterSliceMap)
	sort.Sort(frequencyOrdered)

	frequencyAndKeyString := make(map[int]string, 0)
	for _, v := range frequencyOrdered {
		frequencyAndKeyString[v.Key] = string(v.Value)
	}
	fmt.Println("frequencyAndKeyString=", frequencyAndKeyString)

	frequencyLetters := make([]string, 0)
	for _, v := range frequencyAndKeyString {
		frequencyLetters = append(frequencyLetters, v)
	}
	fmt.Println("frequencyLetters=", frequencyLetters)

	return frequencyLetters
}
