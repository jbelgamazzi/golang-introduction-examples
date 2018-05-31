package main

import "fmt"

var (
	nome, sobrenome = "Jose", "Silva"
	idade           = 32
)

func main() {
	fmt.Println("Olá " + nome + " " + sobrenome)
	// retorna Olá Jose Silva
	fmt.Println(32)
	// retorn 32
}
