package main

import (
	"fmt"
)

func printPattern(n, current int) {
	fmt.Print(current, " ")

	if current == 1 {
		return
	}

	printPattern(n, current-1)

	fmt.Print(current, " ")
}

func main() {
	var n int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	printPattern(n, n)
	fmt.Println()
}
