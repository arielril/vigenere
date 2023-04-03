package vigenere

import (
	"sort"

	"github.com/arielril/vigenere/internal/frequency"
	"github.com/arielril/vigenere/internal/kasiski"
	"github.com/arielril/vigenere/pair"
	"github.com/projectdiscovery/gologger"
)

func crackWithKeyLength(msg string, keyLength int) {
	alphabetLetters := "abcdefghijklmnopqrstuvwxyz"
	numOfLetterAttemptsPerSubkey := 4

	allKeyAndFrequencyMatchList := make(pair.PairFrequencyScoreAndKeyList, 0)

	for nth := 1; nth <= keyLength+1; nth++ {
		nthLetters := kasiski.GetNthSubKeyLetters(nth, keyLength, msg)

		keyAndFrequencyMatchList := make(pair.PairFrequencyScoreAndKeyList, 0)
		for _, possibleKey := range alphabetLetters {
			decryptedText, err := Decode(nthLetters, string(possibleKey))
			if err != nil {
				gologger.Warning().Msgf("could not decrypt msg: %s\n", err)
				continue
			}

			keyAndFrequencyMatchList = append(keyAndFrequencyMatchList, pair.PairFrequencyScoreAndKey{
				Key:   string(possibleKey),
				Value: frequency.GetEnglishFrequencyScore(decryptedText),
			})
		}

		sort.Sort(keyAndFrequencyMatchList)
		allKeyAndFrequencyMatchList = append(
			allKeyAndFrequencyMatchList,
			keyAndFrequencyMatchList[:numOfLetterAttemptsPerSubkey]...,
		)
	}

	// for _, freqAndMatch := range allKeyAndFrequencyMatchList {
	// 	gologger.Info().Msgf("possible letters for letter %v: %v\n", freqAndMatch.Key, freqAndMatch.Value)
	// }
}
