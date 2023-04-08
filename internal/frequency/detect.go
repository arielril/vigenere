package frequency

import (
	"fmt"
	"os"
	"strings"

	"github.com/projectdiscovery/gologger"
)

func DetectIfIsEnglish(msg string) bool {
	wordPercentage := 20
	letterPercentage := 85
	wordsMatch := getEnglishCount(msg) >= wordPercentage

	return wordsMatch && (100 >= letterPercentage)
}

func getEnglishCount(msg string) int {
	englishDictionary := readDictionary("./dictionary/english.txt")

	match := 0
	for _, word := range englishDictionary {
		if strings.Contains(msg, word) {
			match += 1
		}
	}
	if match > 0 {
		fmt.Println("english words matched=", match)
	}
	return match
}

func readDictionary(dictPath string) []string {
	data, err := os.ReadFile(dictPath)
	if err != nil {
		gologger.Fatal().Msgf("could not read english dictionary: %s\n", err)
	}

	r := make([]string, 0)
	for _, line := range strings.Split(string(data), "\n") {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		r = append(r, strings.TrimSpace(line))
	}

	return r
}
