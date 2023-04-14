package frequency

import "strings"

func GetLetterFrequencyFromString(msg string) []int {
	letterFrequency := make([]int, 26)

	for i := 0; i < 26; i++ {
		letter := string('a' + i)
		letterFrequency[i] = strings.Count(msg, letter)
	}

	return letterFrequency
}
