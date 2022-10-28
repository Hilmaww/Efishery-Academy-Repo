package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type point struct {
	ID     int
	barang string
	harga  int
}

func main() {
	bufferReader := bufio.NewReader(os.Stdin)
	//uangTotal := 100000
	var inputHargaBeli string

	// points := []point{
	// 	{ID: 1, barang: "Benih Lele", harga: 50000},
	// 	{ID: 2, barang: "Pakan lele", harga: 25000},
	// 	{ID: 3, barang: "Probiotik A", harga: 75000},
	// 	{ID: 4, barang: "Probiotik Nila B", harga: 10000},
	// 	{ID: 5, barang: "Pakan Nila", harga: 20000},
	// 	{ID: 6, barang: "Benih Nila", harga: 20000},
	// 	{ID: 7, barang: "Cupang", harga: 5000},
	// 	{ID: 8, barang: "Benih Nila", harga: 30000},
	// 	{ID: 9, barang: "Benih Cupang", harga: 10000},
	// 	{ID: 10, barang: "Probiotik B", harga: 10000},
	// }
	fmt.Print("Enter: ")
	inputHargaBeli, err := bufferReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	hasil := parser(inputHargaBeli)
	fmt.Print(hasil)
}

func parser(s string) []int {
	nums := []int{}
	numWhile := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] != 32 {
			numDumm, _ := strconv.Atoi(string(s[i]))
			numWhile = append(numWhile, numDumm)
		} else {
			numDumm, _ := strconv.Atoi(string(s[i]))
			nums = append(nums, numDumm)
		}
	}
	return nums
}

// desc: menentukan apakah barang yang dipilih dapat dibeli atau engga
// input: struct point
// output: boolean, true bila boleh, false jika gaboleh
// syarat: hanya boleh membeli 1 item masing masing barang
//      	dan jumlah total modal adalah 100.000
func farmerBuy(item []point) (bool, int) {
	totalBeli := 0
	for i := 0; i < len(item); i++ {
		totalBeli += item[i].harga
		if totalBeli > 100000 {
			return false, totalBeli
		}
	}
	return true, totalBeli
}
