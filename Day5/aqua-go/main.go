package main

import (
	"aqua5/aqua-go/a"
	"fmt"
	"time"
)

type Z struct {
	x string
	y string
}

type student struct {
	name string
}

// Method dari struct
func (s *student) helloName() {
	fmt.Printf("Hell0 there, %s", s.name)
}
func (s *student) nameChange(namaBaru string) {
	s.name = namaBaru

}

func numbers() {
	for i := 1; 1 <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("number %d\n", i)
	}
}
func alphabets() {
	for i := 'a'; 1 <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("alphabet %c\n", i)
	}
}

func main() {
	var numberA int = 1000
	var numberB *int = &numberA

	fmt.Println("numberA Value:", numberA)
	fmt.Println("numberA Address:", &numberA)
	fmt.Println("numberB Value:", *numberB)
	fmt.Println("numberB Address:", &numberB)

	numberA = 4
	fmt.Println("numberA Value:", numberA)
	fmt.Println("numberA Address:", &numberA)
	fmt.Println("numberB Value:", *numberB)
	fmt.Println("numberB Address:", &numberB)

	// cara ngubah variabel (2)
	*numberB = 50
	fmt.Println("numberA Value:", numberA)
	fmt.Println("numberA Address:", &numberA)
	fmt.Println("numberB Value:", *numberB)
	fmt.Println("numberB Address:", &numberB)

	// pointer nyimpen struct
	x := Z{
		x: "Bla",
		y: "Ble",
	}

	fmt.Printf("%p\n", &x)
	fmt.Printf("%p\n", &x.x)
	fmt.Printf("%p\n", &x.y)

	// inisialisasi
	var std1 = student{name: "hassan"}
	std1.nameChange("opang")
	std1.helloName()

	// import test.go
	a.Hello()

	// Go Routines
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Printf("Terminating")
}
