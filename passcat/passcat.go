// Package passcat lets you generate cryptographically secure,
// memorable passphrases using the diceware algorithm.
//
// See https://en.wikipedia.org/wiki/Diceware
// for details on the Diceware algorithm.
package passcat

import (
	"crypto/rand"
	"math"
	"math/big"
)

// Number of sides on a traditional die.
var sides = big.NewInt(6)

// 'Generate' generates nPhrases collections of nWords words.
func Generate(lang string, nWords, nPhrases int) ([][]string, error) {
	res := make([][]string, 0, nPhrases)
	for i := 0; i < nPhrases; i++ {
		p, err := generatePassphrase(lang, nWords)
		if err != nil {
			return nil, err
		}

		res = append(res, p)
	}

	return res, nil
}

// 'generatePassphrase' generates a collection of nWords words.
func generatePassphrase(lang string, nWords int) ([]string, error) {
	phrase := make([]string, 0, nWords)
	w, err := getWordList(lang)
	if err != nil {
		return nil, err
	}

	for i := 0; i < nWords; i++ {
		digits, err := rollWord()
		if err != nil {
			return nil, err
		}

		word := w.WordAt(digits)
		phrase = append(phrase, word)
	}

	return phrase, nil
}

// 'rollWord' rolls and agregate 5 dice
// to represent a word in the selected wordlist.
func rollWord() (int, error) {
	var res int
	for i := 5; i > 0; i-- {
		roll, err := rollDie()
		if err != nil {
			return 0, err
		}
		res += roll * int(math.Pow(10, float64(i-1)))
	}

	return res, nil
}

// 'rollDie' rolls a single die and returns a value in range 1-6.
func rollDie() (int, error) {
	roll, err := rand.Int(rand.Reader, sides)
	if err != nil {
		return 0, err
	}

	return int(roll.Int64()) + 1, err
}
