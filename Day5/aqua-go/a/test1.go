package a

import "fmt"

func Hello() { // yang ini diexport
	fmt.Println("Hello WOrld!!")
}

func hi() { // yang ini ga di export
	fmt.Println("hi")
}
