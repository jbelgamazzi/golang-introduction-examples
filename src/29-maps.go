package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Jose"] = 32
	m["Maria"] = 75

	fmt.Println(m["Jose"])
	fmt.Println(m["Maria"])
}
