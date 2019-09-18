package tokenizer

import "regexp"

// TreeBankTokenizer splits a sentences into words
//
// This implementation is a port of the SED script of Rober McIntyer avalaible here:
// https://gist.github.com/jdkato/fc8b8c4266dba22d45ac85042ae53b1e  as SED
// https://gist.github.com/desilinguist/1506443#file-treebank2-heilman-py  as PYTHON
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

var startingQuotes = []*regexp.Regexp{}
