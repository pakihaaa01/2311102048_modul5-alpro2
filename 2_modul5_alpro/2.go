package main

import "fmt"

func printStars(n int) {
	if n > 0 {
		printStars(n - 1)
		fmt.Print("*")
	}
}

func starPattern(rows int, current int) {
	if current > rows {
		return
	}
	printStars(current)
	fmt.Println()
	starPattern(rows, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&n)
	fmt.Println("Pola bintang:")
	starPattern(n, 1)
}
