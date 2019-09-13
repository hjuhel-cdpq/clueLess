package encoder

import (
	"strings"
	"unicode"
)

type Document struct {
	Title string
	Body  string
}

type FieldDocument struct {
	Title []string
	Body  []string
}

// Create a new FieldDocument from a Document with Title and Body tokenized
func (doc *Document) Tokenize() *FieldDocument {

	isNotLetter := func(r rune) bool { return !unicode.IsLetter(r) }

	return &FieldDocument{
		Title: strings.FieldsFunc(doc.Title, isNotLetter),
		Body:  strings.FieldsFunc(doc.Body, isNotLetter),
	}
}
