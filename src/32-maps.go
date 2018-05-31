package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Jose"] = 32
	m["Maria"] = 75

	for k, v := range m {
		fmt.Println("Chave:", k, "Valor:", v)
	}

}
