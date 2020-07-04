package passcat

import (
	"testing"
)

const (
	nChecks  = 1000
	nPhrases = 1000
	nWords   = 10
)

var lang = [13]string{
	"de",
	"eff",
	"es",
	"fr",
	"it",
	"jp",
	"nl",
	"no",
	"pl",
	"ro",
	"ru",
	"se",
	"tr",
}

func TestRollDie(t *testing.T) {
	t.Parallel()

	for i := 0; i < nChecks; i++ {
		res, err := rollDie()
		if err != nil {
			t.Fatal(err)
		}

		if res < 1 || res > 6 {
			t.Fatalf("Expected result to be in range 1-6, got %d.", res)
		}
	}
}

func TestRollWord(t *testing.T) {
	t.Parallel()

	for i := 0; i < nChecks; i++ {
		res, err := rollWord()
		if err != nil {
			t.Fatal(err)
		}

		if res < 11111 || res > 66666 {
			t.Fatalf("Expected result to be in range 11111-66666, got %d.", res)
		}
	}
}

func TestGeneratePassphrase(t *testing.T) {
	t.Parallel()

	for _, l := range lang {
		phrase, err := generatePassphrase(l, nWords)
		if err != nil {
			t.Fatal(err)
		}

		if len(phrase) != nWords {
			t.Fatalf("Expected %d words, got %d (%s).", nWords, len(phrase), l)
		}
	}
}

func TestGenerate(t *testing.T) {
	t.Parallel()

	for _, l := range lang {
		res, err := Generate(l, nWords, nPhrases)
		if err != nil {
			t.Fatal(err)
		}

		if len(res) != nPhrases {
			t.Fatalf("Expected %d passphrases, got %d (%s).", nPhrases, len(res), l)
		}
	}
}

func TestGenerateError(t *testing.T) {
	t.Parallel()

	_, err := Generate("unknown", nWords, nPhrases)
	if err == nil {
		t.Fatal("Expected error, got nil.")
	}
}
