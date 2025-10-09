package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// my_number := 42
	my_number := rand.Intn(100) + 1
	fmt.Println(my_number) // Not for prod version - only for test purpose

	for i := 0; i < 3; i = i + 1 {
		guess := ReadNumber()
		if guess == my_number {
			fmt.Println("Das war richtig!")
			return
		}
		if guess < my_number {
			fmt.Println("Die gesuchte Zahl ist größer.")
		} else {
			fmt.Println("Die gesuchte Zahl ist kleiner.")
		}
	}
	fmt.Println("Game Over!")
}

// ReadNumber liefert uns ein int.
func ReadNumber() int {
	var n int // alternativ: n := 0

	fmt.Println("Bitte gibt eine Zahl ein: ")
	fmt.Scan(&n)

	return n
}
