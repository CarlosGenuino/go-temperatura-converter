package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("O numero %d é divisivel por 3\n", i)
		}
	}

	// Parte 2
	for i := 0; i <= 100; i++ {
		multiploDe3 := i%3 == 0
		multiploDe5 := i%5 == 0

		if multiploDe3 {
			fmt.Printf("PIN\n")
		} else if multiploDe5 {
			fmt.Printf("PAN\n")
		} else {
			fmt.Println(i)
		}

	}
}
