package main

import (
	"fmt"

	"github.com/hjuhel-cdpq/clueLess/encoder"
	"github.com/hjuhel-cdpq/clueLess/encoder/tokenizer"
	"github.com/tenfyzhong/cityhash"
)

// Build a hash function from a seed
func BuildHash64Function(seed uint64) encoder.Hash64Function {

	curryHash := func(s []byte) uint64 {
		hash := cityhash.CityHash64WithSeed(s, seed)
		return hash
	}

	return curryHash
}

func add(t encoder.CountMedianSketchTable, s string, h1, h2 encoder.Hash64Function) (uint8, *encoder.MinWise) {

	// Create a new MinWise Struct
	mw := encoder.NewMinWise(h1, h2, 100)

	// Append the tokenized string to the mw
	for _, v := range tokenizer.BytesTokenizer(s) {
		mw.Push(v)
	}

	// Push the min wise to the table
	count_, err := t.Push(mw.GetSignature())
	if err != nil {
		fmt.Println("Mousp !")
	}

	count := count_.(uint8)
	return count, mw
}

func main() {

	// Geneate two hash functions to be used by tge min hash algo
	h1 := BuildHash64Function(42)
	h2 := BuildHash64Function(43)

	// Generate a count table to holds all the minhashs
	table := encoder.NewCountMedianSketchTable8(100)

	// Append the two strings...
	stringtest := "un deux trois quatre cinq six sept huit neuf dix onze douze treize quatorze quinze seize dixsept dixhuit dixneuf vingt"
	stringtest2 := "un deux troisq quatre cinqq six septq huit neufq dix onzeq douze treizeq quatorze quinzeq seize dixseptq dixhuit dixneufq vingt"

	count, mw := add(table, stringtest, h1, h2)
	fmt.Println(count)

	count2, mw2 := add(table, stringtest2, h1, h2)
	fmt.Println(count2)

	count3, _ := add(table, stringtest, h1, h2)
	fmt.Println(count3)

	fmt.Println(mw.Similarity(mw2))

}
