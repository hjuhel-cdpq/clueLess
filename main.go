package main

import (
	"fmt"

	"github.com/hjuhel-cdpq/jloom/encoder"
)

func main() {
	doc := encoder.Document{"c'est un test je suis , pas mal inquier", "M'en parles pas ! Je ne cesse de me demandé comment va se passer le splut unicode"}
	tokenized := doc.Tokenize()

	// Transform the document to rune
	byteList := make([][]byte, len(tokenized.Body))
	for i, item := range tokenized.Body {
		r := []byte(item)
		byteList[i] = r
	}
	seed := uint64(42)
	h := encoder.BuildHash64Function(seed)
	hash := encoder.Minhash(byteList, h, 2)
	fmt.Println(hash)

}
