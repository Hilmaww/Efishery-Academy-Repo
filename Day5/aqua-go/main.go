package main

import (
	"aqua5/aqua-go/a"
	"fmt"
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

}
