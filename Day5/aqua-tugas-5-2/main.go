package main

import (
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
	Desc string
	Jobs string
}

type Cache struct {
	DescString string
	JobList    string
	isFilled   bool
}

func fetch(url string) string {
	data, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer data.Body.Close()
	body, err := io.ReadAll(data.Body)
	if err != nil {
		panic(err)
	}

	// tadi mainan unmarshal, skarang kembalikan bodynya saja
	return string(body)
}

// Mencoba menghitung waktu
func calculateTime(start time.Time) {
	fmt.Println("dari calculate ", start)
	fmt.Printf("took %v\n", time.Since(start))
}

// mencoba agregasi
func (c *Cache) Aggregat() Aggr {
	defer calculateTime(time.Now())

	var wg sync.WaitGroup
	wg.Add(2)

	var aggr Aggr

	if c.isFilled { // kalo dia udah terisi, maka:
		aggr = Aggr{
			Desc: c.DescString,
			Jobs: c.JobList,
		}
		c.isFilled = false
	} else {
		go func() {
			c.DescString = fetch("https://workspace-rho.vercel.app/api/description")
			wg.Done()
		}()
		go func() {
			c.JobList = fetch("https://workspace-rho.vercel.app/api/jobs")
			wg.Done()
		}()

		wg.Wait()
		aggr = Aggr{
			Desc: c.DescString,
			Jobs: c.JobList,
		}
		c.isFilled = true
	}
	return aggr
}

func main() {
	chc := Cache{}
	data1 := chc.Aggregat()
	data2 := chc.Aggregat()
	data3 := chc.Aggregat()
	data4 := chc.Aggregat()

	fmt.Printf("data 1: %+v\n", data1)
	fmt.Println()
	fmt.Printf("data 2: %+v\n", data2)
	fmt.Println()
	fmt.Printf("data 3: %+v\n", data3)
	fmt.Println()
	fmt.Printf("data 4: %+v\n", data4)
}
