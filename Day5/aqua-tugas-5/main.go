package main

import (
	"fmt"
)

type point struct {
	ID     int
	barang string
	harga  int
}

func main() {

	uangTotal := 100000
	var point_choices []point
	var a, b, c, d, e, f, g int
	var isBerhasil string

	points := []point{
		{ID: 1, barang: "Benih Lele", harga: 50000},
		{ID: 2, barang: "Pakan lele", harga: 25000},
		{ID: 3, barang: "Probiotik A", harga: 75000},
		{ID: 4, barang: "Probiotik Nila B", harga: 10000},
		{ID: 5, barang: "Pakan Nila", harga: 20000},
		{ID: 6, barang: "Benih Nila", harga: 20000},
		{ID: 7, barang: "Cupang", harga: 5000},
		{ID: 8, barang: "Benih Nila", harga: 30000},
		{ID: 9, barang: "Benih Cupang", harga: 10000},
		{ID: 10, barang: "Probiotik B", harga: 10000},
	}

	fmt.Print("masukkan ID maksimal 7 barang, tidak boleh ada yang sama: ") // disini sayaa bingung gimana caranya bisa minta banyak input tanpa tau berapa input yang bakal dimasukin user
	_, err := fmt.Scan(&a, &b, &c, &d, &e, &f, &g)
	if err != nil {
		fmt.Println(err)
	}
	idInput := []int{a, b, c, d, e, f, g}
	for i := 0; i < 7; i++ {
		// tadi abis debug
		// fmt.Print(points[idInput[i]-1])
		// fmt.Println(idInput[i])
		point_choices = append(point_choices, points[idInput[i]-1])
	}

	isBought, sum := farmerBuy(point_choices, uangTotal)

	if !isBought {
		fmt.Println("Petani tidak dapat membeli barang yang telah dipilih")
		isBerhasil = "tidak berhasil"
	} else {
		isBerhasil = "berhasil"
	}
	fmt.Printf("Petani %s membeli barang dengan total belanjaan %d", isBerhasil, sum)

}

// func parser(s string) []int {
// 	nums := []int{}
// 	numWhile := []int{}
// 	for i := 0; i < len(s); i++ {
// 		if s[i] != 32 {
// 			numDumm, _ := strconv.Atoi(string(s[i]))
// 			numWhile = append(numWhile, numDumm)
// 		} else {
// 			numDumm, _ := strconv.Atoi(string(s[i]))
// 			nums = append(nums, numDumm)
// 		}
// 	}
// 	return nums
// }

// desc: menentukan apakah barang yang dipilih dapat dibeli atau engga
// input: struct point
// output: boolean, true bila boleh, false jika gaboleh
// syarat: hanya boleh membeli 1 item masing masing barang
//      	dan jumlah total modal adalah 100.000
func farmerBuy(item []point, modal int) (bool, int) {
	totalBeli := 0
	var canBuy bool
	for i := 0; i < len(item); i++ {
		totalBeli += item[i].harga
	}
	if totalBeli > modal {
		canBuy = false
	} else {
		canBuy = true
	}
	return canBuy, totalBeli
}

func urutHargaBarang(p []point) {
	items := []int{}
	for i := 0; i < len(p); i++ {
		items = append(items, p[i].harga)
	}
}
