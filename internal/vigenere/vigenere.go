package vigenere

import (
	"fmt"
	"strings"

	"github.com/arielril/vigenere/internal/kasiski"
	"github.com/projectdiscovery/gologger"
)

// Encode encodes a msg with a key using the vigenere cipher
func Encode(msg, key string) (string, error) {
	// encryption formula
	// E_i = (P_i + K_i) % 26

	cipheredMessage := make([]string, len(msg))

	for i := 0; i < len(msg); i++ {
		pI := msg[i] - 'a'
		kI := key[i%len(key)] - 'a'
		eI := (pI + kI) % 26

		gologger.Debug().Msgf("P_i=%v\n", pI)
		gologger.Debug().Msgf("K_i=%v\n", kI)
		gologger.Debug().Msgf("E_i=%v\n", eI)
		gologger.Debug().Msgf("----------\n")

		cipheredMessage[i] = fmt.Sprintf("%c", eI+'a')
	}

	return strings.Join(cipheredMessage, ""), nil
}

// Decode decodes an encrypted msg with a key using the vigenere cipher
func Decode(msg, key string) (string, error) {
	// decryption formula
	// D_i = (E_i - K_i) % 26

	decipheredMessage := make([]string, len(msg))
	for i := 0; i < len(msg); i++ {
		eI := msg[i] - 'a'
		kI := key[i%len(key)] - 'a'
		dI := (eI - kI + 26) % 26

		gologger.Debug().Msgf("E_i=%v\n", eI)
		gologger.Debug().Msgf("K_i=%v\n", kI)
		gologger.Debug().Msgf("D_i=%v\n", dI)
		gologger.Debug().Msg("----------\n")

		decipheredMessage[i] = fmt.Sprintf("%c", dI+'a')
	}

	return strings.Join(decipheredMessage, ""), nil
}

func Crack(msg string) (string, error) {
	gologger.Silent().Msgf("message length: %d\n", len(msg))

	possibleKeyLength := kasiski.GetPossibleKeyLength(msg)
	gologger.Silent().Msgf("possible key length: %d\n", possibleKeyLength)

	possibleKey := kasiski.GetPossibleKey(possibleKeyLength, msg)
	gologger.Silent().Msgf("possible key: %s\n", possibleKey)

	decodedMsg, err := Decode(msg, possibleKey)
	if err != nil {
		gologger.Fatal().Msgf("could not decode msg with possible key: %s\n", err)
	}

	return decodedMsg, nil

	// possibleKeyLengths := kasiski.KasiskiExamination(msg)
	// gologger.Info().Msgf("kasiski most possible key lengths= %v\n", possibleKeyLengths)

	// for _, keyLength := range possibleKeyLengths {
	// 	gologger.Info().Msgf("attempting to crack with key length %v...\n", keyLength)

	// 	crackWithKeyLength(msg, keyLength)

	// for i := 1; i <= keyLength; i++ {
	// 	nthSubKeyLetters := kasiski.GetNthSubKeyLetters(i, keyLength, msg)
	// 	gologger.Debug().Msgf("got nth subkey letters= %v\n", nthSubKeyLetters)

	// 	frequencyScores := make(pair.PairFrequencyScoreAndKeyList, 0)
	// 	for _, possibleKey := range possibleLetters {
	// 		decryptedText, err := Decode(nthSubKeyLetters, string(possibleKey))
	// 		if err != nil {
	// 			gologger.Warning().Msgf("could not decrypt message: %s\n", err)
	// 		}

	// 		frequencyScores = append(frequencyScores, pair.PairFrequencyScoreAndKey{
	// 			Key:   string(possibleKey),
	// 			Value: frequency.GetEnglishFrequencyScore(decryptedText),
	// 		})
	// 	}

	// 	sort.Sort(frequencyScores)
	// 	gologger.Debug().Msgf("frequency scores: %#v\n", frequencyScores)

	// 	allFrequencyScores = append(allFrequencyScores, frequencyScores...)
	// }

	// for i := 0; i < len(allFrequencyScores); i++ {
	// 	gologger.Info().Msgf("possible letters for letter %v of the key: %v\n", i, allFrequencyScores[i].Value)
	// }

	// 	crackedMessage := ""
	// 	fmt.Println("cracked message=", crackedMessage)
	// }

}
