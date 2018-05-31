package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	pessoa := Pessoa{Nome: "Fulano", Idade: 1}
	fmt.Println(pessoa.Nome)
	fmt.Println(pessoa.Idade)

	Jose := &pessoa
	Jose.Nome = "Jose"
	Jose.Idade = 32

	fmt.Println(pessoa.Nome)
	fmt.Println(pessoa.Idade)
}
