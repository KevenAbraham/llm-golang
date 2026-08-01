package main

import (
	"fmt"

	"llm-golang/internal/tokenizer"
)

func main() {
	token := tokenizer.NewTokenizer()
	token.Train("O gato subiu no telhado o gato dormiu")

	fmt.Println("A vocabulary size: ", token.VocabSize())

	ids := token.Encode("O gato subiu")
	fmt.Println(ids)

	for _, id := range ids {
		fmt.Println(token.Decode(id))
	}

	ids2 := token.Encode("O cachorro correu pelo telhado")

	fmt.Println("Mensgaem enviada: o cachorro correu pelo telhado ")
	fmt.Println("ids gerados: ", ids2)
	fmt.Println("decode: ", token.Decode(ids2[0]))
}
