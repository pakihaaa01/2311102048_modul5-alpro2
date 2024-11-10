package main

import (
	"fmt"
)

func printbilanganganjil(n, current int) {
	if current > n {
		return
	}

	fmt.Print(current, " ")

	printbilanganganjil(n, current+2)
}

func main() {
	var n int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	printbilanganganjil(n, 1)
	fmt.Println()
}
