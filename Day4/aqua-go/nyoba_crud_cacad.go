package main

import "fmt"

type database struct {
	name string
	age  int
}

func getSome(data []database) (string, []database) {
	headers := "Content-Type: application/json"
	return headers, data
}

func postSome(data []database, inputData []database) (string, []database) {
	headers := "Content-Type: application/json"
	data = append(data, inputData...)
	return headers, data
}
func mauin() {
	var data = []database{
		{name: "jason miraj", age: 50},
		{name: "bakso bakso", age: 2},
	}

	var dataInput = []database{
		{name: "okoJ kaP", age: 46},
	}

	for i := 1; i < 10; i++ {

		//if masukan get
		fmt.Println(getSome(data))

		// if masukan post
		fmt.Println(postSome(data, dataInput))

	}
}
