package main

import "fmt"

func main() {
	test_slice := []string{"Eduardo", "Escudero", "Arens"}
	out, success := ResliceAnything(test_slice, 1)
	if !success {
		panic(success)
	}

	fmt.Printf("%+q", out)
}
