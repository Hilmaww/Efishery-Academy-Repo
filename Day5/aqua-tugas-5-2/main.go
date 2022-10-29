package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Desc struct {
	Description string `json:"description"`
}
type Job struct {
	Title      string `json:"title"`
	WorkFrom   string `json:"work_from"`
	Department string `json:"department"`
}
type Aggr struct {
	Desc Desc
	Jobs []Job
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

// Mencoba menghitung waktu
func caculateTime(start time.Time) {
	fmt.Println("dari calculate ", start)
	fmt.Printf("took %v\n", time.Since(start))
}

// mencoba agregasi
func Aggregat() (Aggr, error) {
	defer caculateTime(time.Now())
	var wg sync.WaitGroup
	wg.Add(2)

	var desc Desc
	var descErr error
	var jobs []Job
	var jobsErr error

	go func() {
		descErr = fetch("https://workspace-rho.vercel.app/api/description", &desc)
		wg.Done()
	}()
	if descErr != nil {
		return Aggr{}, descErr
	}
	go func() {
		jobsErr = fetch("https://workspace-rho.vercel.app/api/description", &desc)
		wg.Done()
	}()

	wg.Wait()

	if jobsErr != nil {
		return Aggr{}, jobsErr
	}

	aggr := Aggr{
		Desc: desc,
		Jobs: jobs,
	}
	return aggr, nil
}

func main() {

	aggr, err := Aggregat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", aggr)
}
