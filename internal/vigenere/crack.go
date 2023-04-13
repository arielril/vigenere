package vigenere

// func crackWithKeyLength(msg string, keyLength int) {
// 	alphabetLetters := "abcdefghijklmnopqrstuvwxyz"
// 	numOfLetterAttemptsPerSubkey := 4

// 	allKeyAndFrequencyMatchList := make(pair.PairFrequencyScoreAndKeyList, 0)

// 	for nth := 1; nth <= keyLength+1; nth++ {
// 		nthLetters := kasiski.GetNthSubKeyLetters(nth, keyLength, msg)

// 		keyAndFrequencyMatchList := make(pair.PairFrequencyScoreAndKeyList, 0)
// 		for _, possibleKey := range alphabetLetters {
// 			decryptedText, err := Decode(nthLetters, string(possibleKey))
// 			if err != nil {
// 				gologger.Warning().Msgf("could not decrypt msg: %s\n", err)
// 				continue
// 			}

// 			keyAndFrequencyMatchList = append(keyAndFrequencyMatchList, pair.PairFrequencyScoreAndKey{
// 				Key:   string(possibleKey),
// 				Value: frequency.GetEnglishFrequencyScore(decryptedText),
// 			})
// 		}

// 		sort.Sort(sort.Reverse(keyAndFrequencyMatchList))
// 		allKeyAndFrequencyMatchList = append(
// 			allKeyAndFrequencyMatchList,
// 			keyAndFrequencyMatchList[:numOfLetterAttemptsPerSubkey]...,
// 		)
// 	}

// 	for _, freqAndMatch := range allKeyAndFrequencyMatchList {
// 		gologger.Info().Msgf("possible letters for letter %v: %v\n", freqAndMatch.Key, freqAndMatch.Value)
// 	}

// 	everyCombinationOfLikelyLetters := cartesian.NewCartesianProduct(makeIntSlice(numOfLetterAttemptsPerSubkey, keyLength))

// 	// fmt.Println("allKeyAndFrequencyMatchList=", allKeyAndFrequencyMatchList)
// 	// fmt.Println("everyCombinationOfLikelyLetters=", everyCombinationOfLikelyLetters)
// 	for _, combination := range everyCombinationOfLikelyLetters.Values() {
// 		possibleKey := strings.Builder{}
// 		for i := 0; i < keyLength; i++ {
// 			// fmt.Println("allfreqscore key rune=", allKeyAndFrequencyMatchList[i].Key)
// 			// fmt.Println("possible key rune=", combination[i])
// 			possibleKey.WriteString(allKeyAndFrequencyMatchList[combination[i].(int)].Key)
// 		}

// 		// gologger.Silent().Msgf("attempting with key %s\n", possibleKey.String())

// 		decryptedText, err := Decode(msg, possibleKey.String())
// 		if err != nil {
// 			gologger.Error().Msgf("could not decode vigenere: %s\n", err)
// 		}
// 		// fmt.Println("decrypted text=", decryptedText)

// 		if frequency.DetectIfIsEnglish(decryptedText) {
// 			gologger.Silent().Msgf("possible used key %s\n", possibleKey.String())
// 			gologger.Silent().Msgf("decryption: %s\n", decryptedText)
// 		}
// 	}
// }

// func makeIntSlice(n, repeat int) []any {
// 	r := make([]any, 0)

// 	for i := 0; i < repeat; i++ {
// 		a := make([]int, 0)
// 		for j := 0; j < n; j++ {
// 			a = append(a, j)
// 		}
// 		r = append(r, a)
// 	}

// 	return r
// }
