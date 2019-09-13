package main

import (
	"fmt"

	"github.com/hjuhel-cdpq/jloom/encoder"
)

func main() {
	doc := encoder.Document{"c'est un test je suis , pas mal inquier", "M'en parles pas ! Je ne cesse de me demandé comment va se passer le splut unicode"}

	out := doc.Tokenize()

	fmt.Println(len(out.Title))
	fmt.Printf("%+v\n", *out)
}
