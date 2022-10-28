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
		fmt.Printf("Petani %s membeli barang dengan total belanjaan %d\n", isBerhasil, sum)
	} else {
		isBerhasil = "berhasil"
		fmt.Printf("Petani %s membeli barang dengan total belanjaan %d\n", isBerhasil, sum)
		fmt.Println("Total produk dengan modal Rp 100.000: ")
	}

	fmt.Println("Total produk dengan modal Rp 100.000: ")
	for i := 0; i < len(point_choices); i++ {
		fmt.Printf("%s - %d\n", point_choices[i].barang, point_choices[i].harga)
	}
	fmt.Println("---------------------------------------------------------------------")
	idHargaTerkecil, idHargaTerbesar := urutHargaBarang(point_choices)
	fmt.Printf("Produk termurah: %s Rp%d\n", point_choices[idHargaTerkecil].barang, point_choices[idHargaTerkecil].harga)
	fmt.Printf("Produk termahal: %s Rp%d\n", point_choices[idHargaTerbesar].barang, point_choices[idHargaTerbesar].harga)
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("Daftar produk dengan harga Rp 10.000: ")
	item10k := harga10k(point_choices)
	for i := 0; i < len(item10k); i++ {
		fmt.Printf("%s - %d\n", point_choices[i].barang, point_choices[i].harga)
	}
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

// desc: membuat fungsi untuk menentukan harga paling murah pada barang barang di point efishery, tanpa sort built-in function
// input: struct point
// output: ID (int) dengan harga terkecil dan harga terendah
func urutHargaBarang(p []point) (int, int) {
	// inisialisasi items
	items := p
	for i := 0; i < len(items); i++ {
		fmt.Print("i: ", i)
		if i != 6 {
			for j := i + 1; j < len(items); j++ {
				fmt.Println("j: ", j)
				if items[i].harga > items[j].harga {
					items[i], items[j] = items[j], items[i]
				}

			}
		}

	}
	return items[0].ID, items[len(items)].ID

}

// desc: menampilkan barang dengan harga 10.000
// input: struct point
// output: slices of struct point dengan harga 10.000
func harga10k(p []point) []point {
	var item10k []point

	for i := 0; i < len(p); i++ {
		if p[i].harga == 10000 {
			item10k = append(item10k, p[i])
		}
	}
	return item10k
}
