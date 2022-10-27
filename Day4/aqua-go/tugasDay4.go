package main

import (
	"fmt"
	"strings"
)

func rectNtree(n int) {
	if n%2 == 0 { // genap
		for i := 0; i < n; i++ {
			fmt.Println(strings.Repeat("*", n))
		}
	} else { // genap
		for i := 0; i < n; i++ {
			fmt.Println(strings.Repeat("*", i+1))
		}

	}
}

func main() {
	fmt.Println("Masukkan n: ")

	// var then variable name then variable type
	var n int

	// Taking input from user
	fmt.Scanln(&n)
	rectNtree(n)

}
