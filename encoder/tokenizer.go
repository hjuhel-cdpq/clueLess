package encoder

import (
	"strings"
	"unicode"
)

type ShinglesBytes [][]byte

// Create a new FieldDocument from a Document with Title and Body tokenized
func BytesTokenizer(s string) ShinglesBytes {

	// Split document by spaces
	isNotLetter := func(r rune) bool { return !unicode.IsLetter(r) }
	sShingles := strings.FieldsFunc(s, isNotLetter)

	// Convert the string to a list of bytes
	sShinglesBytes := make([][]byte, len(sShingles))

	for i, shingle := range sShingles {
		sShinglesBytes[i] = []byte(shingle)

	}

	return sShinglesBytes
}
