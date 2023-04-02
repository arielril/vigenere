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
		pI := msg[i] - 97
		kI := key[i%len(key)] - 97
		eI := (pI + kI) % 26

		gologger.Debug().Msgf("P_i=%v\n", pI)
		gologger.Debug().Msgf("K_i=%v\n", kI)
		gologger.Debug().Msgf("E_i=%v\n", eI)
		gologger.Debug().Msgf("----------\n")

		cipheredMessage[i] = fmt.Sprintf("%c", eI+97)
	}

	return strings.Join(cipheredMessage, ""), nil
}

// Decode decodes an encrypted msg with a key using the vigenere cipher
func Decode(msg, key string) (string, error) {
	// decryption formula
	// D_i = (E_i - K_i) % 26

	decipheredMessage := make([]string, len(msg))
	for i := 0; i < len(msg); i++ {
		eI := msg[i] - 97
		kI := key[i%len(key)] - 97
		dI := (eI - kI + 26) % 26

		gologger.Debug().Msgf("E_i=%v\n", eI)
		gologger.Debug().Msgf("K_i=%v\n", kI)
		gologger.Debug().Msgf("D_i=%v\n", dI)
		gologger.Debug().Msg("----------\n")

		decipheredMessage[i] = fmt.Sprintf("%c", dI+97)
	}

	return strings.Join(decipheredMessage, ""), nil
}

func Crack(msg string) (string, error) {
	_ = kasiski.GetPossibleKeyLengths(msg)
	return "", nil
}
