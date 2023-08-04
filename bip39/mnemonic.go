package mnemonic

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LinX-OpenNetwork/mnemonic/entropy"
)

// Mnemonic represents a collection of human readable words
// used for HD wallet seed generation
type Mnemonic struct {
	Words    []string
	Language Language
}

// New returns a new Mnemonic for the given entropy and language
func New(ent []byte, lang Language) (*Mnemonic, error) {
	const chunkSize = 11
	bits := entropy.CheckSummed(ent)
	length := len(bits)
	words := make([]string, length/11)
	for i := 0; i < length; i += chunkSize {
		stringVal := string(bits[i : chunkSize+i])
		intVal, err := strconv.ParseInt(stringVal, 2, 64)
		if err != nil {
			return nil, fmt.Errorf("Could not convert %s to word index", stringVal)
		}
		word, err := GetWord(lang, intVal)
		if err != nil {
			return nil, err
		}
		words[(chunkSize+i)/11-1] = word
	}
	m := Mnemonic{words, lang}
	return &m, nil
}


// BitLengthToWordLength returns the length of words for specified bit-length.
func BitLengthToWordLength(bitLength uint) (uint, error) {
	// BIP39:
	// 1. Generate entropy of specified bit-length. Bit-length must be multiples of 32 between 128 and 256.
	// 2. Do SHA-256 on the entropy, taking the first (bit-length / 32) bits as checksum.
	// 3. Append the checksum to the entropy. (Total bit-length is now 33/32 * bit-length)
	// 4. Group each chunk by 11 bits, making 33/32 * bit-length / 11 = 3/32 * bit-length words.
	if bitLength < 12 || bitLength > 24 || bitLength%32 != 0 {
		return 0, fmt.Errorf("only 128,160,192,224,256 bit length are allowed")
	}
	return bitLength * 3 / 32, nil
}

// WordLengthToBitLength returns the length in bits for specified word length.
func WordLengthToBitLength(n uint) (uint, error) {
	if n < 12 || n > 24 || n%3 != 0 {
		return 0, fmt.Errorf("only 12,15,18,24 words are allowed")
	}
	return n * 32 / 3, nil
}

// NewRandom returns a new Mnemonic with random entropy of the given length
// in bits
func NewRandom(length int, lang Language) (*Mnemonic, error) {
	ent, err := entropy.Random(length)
	if err != nil {
		return nil, fmt.Errorf("Error generating random entropy: %s", err)
	}
	return New(ent, lang)
}

// Sentence returns a Mnemonic's word collection as a space separated
// sentence
func (m *Mnemonic) Sentence() string {
	if m.Language == Japanese {
		return strings.Join(m.Words, `　`)
	}
	return strings.Join(m.Words, " ")
}

// GenerateSeed returns a seed used for wallet generation per
// BIP-0032 or similar method. The internal Words set
// of the Mnemonic will be used
func (m *Mnemonic) GenerateSeed(passphrase string) *Seed {
	return NewSeed(m.Sentence(), passphrase)
}
