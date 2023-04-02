package kasiski

import (
	"strings"
)

func AttemptToCrackWithKeyLength(msg string, keyLength int) string {
	return ""
}

// GetNthSubKeyLetters return every nth letter for each keyLength
func GetNthSubKeyLetters(startPosition, keyLength int, msg string) string {
	letters := make([]string, 0)
	for i := startPosition - 1; i < len(msg); {
		letters = append(letters, string(msg[i]))
		i += keyLength
	}

	return strings.Join(letters, "")
}
