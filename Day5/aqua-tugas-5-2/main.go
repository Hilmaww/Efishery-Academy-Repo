package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Desc struct {
	Description string `json:"description"`
}
type Job struct {
	Title      string `json:"title"`
	WorkFrom   string `json:"work_from"`
	Department string `json:"department"`
}

func fetch(url string, d interface{}) error {
	data, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer data.Body.Close()
	body, err := io.ReadAll(data.Body)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(body, d); err != nil {
		return err
	}
	return err

}

func main() {
	var desc Desc
	var jobs []Job
	err := fetch("https://workspace-rho.vercel.app/api/description", &desc)
	if err != nil {
		panic(err)
	}
	err = fetch("https://workspace-rho.vercel.app/api/jobs", &jobs)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", desc)
	fmt.Printf("%+v\n", jobs)
}
