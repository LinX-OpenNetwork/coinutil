# Mnemonic
[![Build Status](https://travis-ci.org/LinX-OpenNetwork/coinutil.svg?branch=master)](https://travis-ci.org/LinX-OpenNetwork/coinutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/LinX-OpenNetwork/coinutil)](https://goreportcard.com/report/github.com/LinX-OpenNetwork/coinutil)
[![GoDoc](https://godoc.org/github.com/LinX-OpenNetwork/coinutil?status.svg)](https://godoc.org/github.com/LinX-OpenNetwork/coinutil)

A BIP 39 implementation in Go.

Features:

* Generating human readable sentences for seed generation - a la [BIP 32](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki)
* All languages mentioned in the [proposal](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki) supported.
* 128 bit (12 words) through 256 bit (24 words) entropy.

## [`coinutil`](https://godoc.org/github.com/LinX-OpenNetwork/coinutil) package

* Generates human readable sentences and the seeds derived from them.
* Supports all languages mentioned in the [BIP 39 proposal](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki).
* Supports ideogrpahic spaces for Japanese language.

Example:

```go
package main

import (
    "fmt"
    "github.com/LinX-OpenNetwork/coinutil/bip39"
)

func main() {
    // generate a random Mnemonic in English with 256 bits of entropy
    m, _ := bip39.NewRandom(256, bip39.English)

    // print the Mnemonic as a sentence
    fmt.Println(m.Sentence())

    // inspect underlying words
    fmt.Println(m.Words)

    // generate a seed from the Mnemonic
    seed := m.GenerateSeed("passphrase")

    // print the seed as a hex encoded string
    fmt.Println(seed)
}
```

## [`entropy`](https://godoc.org/github.com/LinX-OpenNetwork/coinutil/entropy) package

* Supports generating random entropy in the range of 128-256 bits
* Supports generating entropy from a hex string

Example:

```go
package main

import (
    "fmt"
    "github.com/LinX-OpenNetwork/coinutil/bip39"
    "github.com/LinX-OpenNetwork/coinutil/entropy"
)

func main() {
    // generate some entropy from a hex string
    ent, _ := entropy.FromHex("8197a4a47f0425faeaa69deebc05ca29c0a5b5cc76ceacc0")
    
    // generate a Mnemonic in bip39 with the generated entropy
    jp, _ := bip39.New(ent, bip39.Japanese)

    // print the Mnemonic as a sentence
    fmt.Println(jp.Sentence())

    // generate some random 256 bit entropy
    rnd, _ := entropy.Random(256)
    
    // generate a Mnemonic in Spanish with the generated entropy
    sp, _ := bip39.New(rnd, bip39.Spanish)

    // print the Mnemonic as a sentence
    fmt.Println(sp.Sentence())
}
```

# Installation

To install Mnemonic, use `go get`:

    go get github.com/LinX-OpenNetwork/coinutil

This will then make the following packages available to you:

    github.com/LinX-OpenNetwork/coinutil
    github.com/LinX-OpenNetwork/coinutil/bip32
    github.com/LinX-OpenNetwork/coinutil/bip39
    github.com/LinX-OpenNetwork/coinutil/entropy

Import the `mnemonic` package into your code using this template:

```go
package yours

import (
  "github.com/LinX-OpenNetwork/coinutil/bip39"
)

func MnemonicJam(passphrase string) {

  m := bip39.NewRandom(passphrase)

}
```

# Contributing

Please feel free to submit issues, fork the repository and send pull requests!

When submitting an issue, we ask that you please include a complete test function that demonstrates the issue.
