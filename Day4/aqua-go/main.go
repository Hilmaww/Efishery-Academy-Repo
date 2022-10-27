package main

import "fmt"

type person struct {
	name string
	age  int
}
type student struct {
	matkul string
	person
}

func moin() {
	fmt.Println("Hellow world")

	name := "wock"
	number := 3
	var is_true bool = true
	distance := 9.2371
	fmt.Printf("I took %d %s %f km away. %t\n", number, name, distance, is_true)
	fmt.Printf("the distance: %.2f", distance)

	// deklarasi variabel
	var enters = `\nMemakan semua
		Lorem Ipsum Dolor
		Sit Amet
	`
	fmt.Printf("%s", enters)

	var namaDepan, namaBelakang = "ahmad", "dhani"
	fmt.Printf("%s %s \n", namaDepan, namaBelakang)

	var muridSatu student
	muridSatu.name = "Joko"
	muridSatu.age = 47
	muridSatu.matkul = "PSWD"

	fmt.Println(muridSatu.name)
	fmt.Println(muridSatu.age)
	fmt.Println(muridSatu.person.age)
	fmt.Println(muridSatu.matkul)

	students := []student{
		{matkul: "Elkom", person: person{name: "maryana", age: 18}},
	}

	fmt.Println(students)

	chicken := map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	var buah = []string{"apel", "melon", "gedhang", "nangka"}
	fmt.Println("buah[0:2] => ", buah[0:2])
	fmt.Println("buah[0:4] => ", buah[0:4])
	fmt.Println("buah[0:0] =› ", buah[0:0])
	fmt.Println("buah[4:4] => ", buah[4:4])
	// fmt.Println("buah[4:0]=> ", buah[4:0]) // tidak bisa, karena nilai a haru
	fmt.Println("buah[:]]=› ", buah[:])
	fmt.Println("buah[2:) => ", buah[2:])
	fmt.Println("buah(:2]=>", buah[:2])

}
