package tokenizer

import (
	"regexp"
	"strings"
)

// TreeBankTokenizer splits a sentences into words
//
// This implementation is a port of the SED script of Rober McIntyer avalaible here:
// https://gist.github.com/jdkato/fc8b8c4266dba22d45ac85042ae53b1e  in SED (original)
// https://gist.github.com/desilinguist/1506443#file-treebank2-heilman-py  in PYTHON
// The tokenizer expect one sentence to be tokenized.
type TreeBankTokenizer struct{}

// TreeBankeTokenizer constructor
func NewTreeBankTokenizer() *TreeBankTokenizer {
	return &TreeBankTokenizer{}
}

var contractions = []*regexp.Regexp{
	regexp.MustCompile(`(?i)\b(can)(not)\b`),
	regexp.MustCompile(`(?i)\b(d)('ye)\b`),
	regexp.MustCompile(`(?i)\b(d)('ye)\b`),
	regexp.MustCompile(`(?i)\b(gim)(me)\b`),
	regexp.MustCompile(`(?i)\b(gon)(na)\b`),
	regexp.MustCompile(`(?i)\b(got)(ta)\b`),
	regexp.MustCompile(`(?i)\b(lem)(me)\b`),
	regexp.MustCompile(`(?i)\b(mor)('n)\b`),
	regexp.MustCompile(`(?i)\b(wan)(na) `),
	regexp.MustCompile(`(?i)('t)(is)\b`),
	regexp.MustCompile(`(?i)('t)(was)\b`),
	regexp.MustCompile(`(?i)\b(whad)(dd)(ya)\b`),
	regexp.MustCompile(`(?i)\b(wha)(t)(cha)\b`),
}

var newlines = regexp.MustCompile(`(?:\n|\n\r|\r)`)
var spaces = regexp.MustCompile(`(?: {2,})`)

var startingQuotes = map[string]*regexp.Regexp{
	"$1 ``": regexp.MustCompile(`([ (\[{<)"`),
	"``":    regexp.MustCompile(`^(")`),
	" ``":   regexp.MustCompile(`( ")`),
}

var startingQuotes2 = map[string]*regexp.Regexp{
	" $1 ": regexp.MustCompile("(``)"),
}

var punctuation = map[string]*regexp.Regexp{
	" $1 $2":   regexp.MustCompile(`([:,])([^\d])`),
	" ... ":    regexp.MustCompile(`\.\.\.`),
	"$1 $2$3 ": regexp.MustCompile(`([^\.])(\.)([\]\)}>"\']*)\s*$`),
	"$1 ' ":    regexp.MustCompile(`([^'])' `),
}

var punctuation2 = []*regexp.Regexp{
	regexp.MustCompile(`([:,])$`),
	regexp.MustCompile(`([;@#$%&?!])`),
}

var brackets = map[string]*regexp.Regexp{
	" $1 ": regexp.MustCompile(`([\]\[\(\)\{\}\<\>])`),
	" -- ": regexp.MustCompile(`--`),
}
var endingQuotes = map[string]*regexp.Regexp{
	" '' ": regexp.MustCompile(`"`),
}

var endingQuotes2 = []*regexp.Regexp{
	regexp.MustCompile(`'(\S)(\'\')'`),
	regexp.MustCompile(`([^' ])('[sS]|'[mM]|'[dD]|') `),
	regexp.MustCompile(`([^' ])('ll|'LL|'re|'RE|'ve|'VE|n't|N'T) `),
}

// Tokenize splits a sentence into a slice of words
//
// The tokenizer implements PennTreeBank tokenizer :
// 1) Split on contraction
// 2) Split on-non terminating punctuation
// 3) Split on single quotes if followed by whitespace
// 4) Split on period appareaing at the end of lines
//
// NOTE: Not a sentences tokenizer
func (t TreeBankTokenizer) Tokenize(text string) []string {

	for subs, r := range startingQuotes {
		text = r.ReplaceAllString(text, subs)
	}

	for subs, r := range startingQuotes2 {
		text = r.ReplaceAllString(text, subs)
	}

	for subs, r := range punctuation {
		text = r.ReplaceAllString(text, subs)
	}

	for _, r := range punctuation2 {
		text = r.ReplaceAllString(text, " $1 ")
	}

	for subs, r := range brackets {
		text = r.ReplaceAllString(text, subs)
	}

	// Insert spaces at start and end, to make the rest of the tokenization easier
	text = " " + text + " "

	for subs, r := range endingQuotes {
		text = r.ReplaceAllString(text, subs)
	}

	for _, r := range endingQuotes2 {
		text = r.ReplaceAllString(text, "$1 $2 ")
	}

	for _, r := range contractions {
		text = r.ReplaceAllString(text, " $1 $2 ")
	}

	return strings.Split(text, " ")
}
