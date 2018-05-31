package main

import "fmt"

func main() {
	m := map[string]int{
		"Jose":  32,
		"Maria": 75,
	}

	fmt.Println(m["Jose"])
	fmt.Println(m["Maria"])
}
