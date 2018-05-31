package main

import "fmt"

var (
	nome, sobrenome = "Jose", "Silva"
)

func main() {
	fmt.Println("Olá " + nome + " " + sobrenome)
	// retorna Olá Jose Silva
}
