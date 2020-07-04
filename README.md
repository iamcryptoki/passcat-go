# Passcat

Passcat lets you generate cryptographically secure, memorable passphrases using the diceware algorithm.

Contents
* [Installation](#installation)
* [Usage (library)](#usage-library)
* [Usage (CLI)](#usage-cli)
* [Available wordlists](#available-wordlists)

Installation
------------

```sh
$ go get -u github.com/iamcryptoki/passcat-go/passcat
```

Usage (library)
---------------

```golang
package main

import (
    "log"
    "strings"

    "github.com/iamcryptoki/passcat-go/passcat"
)

func main() {
    // Generate 8 words in French using the diceware algorithm.
    p1, err := passcat.Generate("fr", 8, 1)
    if err != nil  {
        log.Fatal(err)
    }

    log.Printf(strings.Join(p1, "-"))

    // Generate 5 passphrases in Japanese
    // using the diceware algorithm.
    p2, err := passcat.Generate("jp", 8, 5)
    if err != nil {
        log.Fatal(err)
    }

    for _, p := range p2 {
        log.Printf(strings.Join(p, "-"))
    }
}
```

Usage (CLI)
-----------

```sh
$ go install github.com/iamcryptoki/passcat-go/cmd/passcat
```

Basic usage:

```sh
$ passcat
risk freckled onslaught bouncing subpar winner syndrome
```

Specify the number of words to use in the passphrase:

```sh
$ passcat -w 10
provoke study crease fling dispense erasable drone monkhood trilogy outboard
```

Specify the number of passphrases to generate:

```sh
$ passcat -p 5
shelving ample marry fiscally osmosis boogeyman shrank
judiciary said eldercare utensil moonwalk frail swifter
capacity edginess rimless passably overhang jury backlash
attic ritalin headroom repost unhinge frenzied very
untangled gallstone gender anemic willfully calorie stainless
```

Specify a wordlist other than EFF:

```sh
$ passcat -l fr
logiez neuve brebis mastoc ortie tentes tilts
```

Available wordlists
-------------------

Language | Code
-------- | ----
EFF (English) | `eff`
Dutch | `nl`
French | `fr`
German | `de`
Italian | `it`
Japanese | `jp`
Norwegian | `no`
Polish | `pl`
Romanian | `ro`
Russian | `ru`
Spanish | `es`
Swedish | `se`
Turkish | `tr`

## License

This code is released under a free software [license](LICENSE.txt) and you are welcome to fork it.