package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/iamcryptoki/passcat-go/passcat"
)

var (
	nWords   = flag.Int("w", 7, "Number of words to use in the passphrase.")
	nPhrases = flag.Int("p", 1, "Number of passphrases to generate.")
	sep      = flag.String("s", " ", "Character to use between words.")
	lang     = flag.String("l", "eff", "Specify a wordlist.")
)

func main() {
	flag.Parse()

	res, err := passcat.Generate(*lang, *nWords, *nPhrases)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	for _, p := range res {
		fmt.Println(strings.Join(p, *sep))
	}
}
