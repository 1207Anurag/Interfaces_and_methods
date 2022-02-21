package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func CallUrl(id string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts" + id)
	if err != nil {
		fmt.Println("error accessing URL")
	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("cannot read body of Response")
		} else {
			fmt.Println(string(body))
		}
	}
}
func CallUrl1(id string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts" + id)
	if err != nil {
		fmt.Println("error accessing URL")
	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("cannot read body of Response")
		} else {
			fmt.Println(string(body))
		}
	}
}
func CallUrl2(id string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts" + id)
	if err != nil {
		fmt.Println("error accessing URL")
	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("cannot read body of Response")
		} else {
			fmt.Println(string(body))
		}
	}
}
func CallUrl3(id string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts" + id)
	if err != nil {
		fmt.Println("error accessing URL")
	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("cannot read body of Response")
		} else {
			fmt.Println(string(body))
		}
	}
}
func CallUrl4(id string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts" + id)
	if err != nil {
		fmt.Println("error accessing URL")
	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("cannot read body of Response")
		} else {
			fmt.Println(string(body))
		}
	}
}
848170675
1153234471
func main() {
	st := time.Now()
	wg := new(sync.WaitGroup)
	wg.Add(200)
	for i := 1; i <= 200; i++ {
		id := strconv.Itoa(i)
		go CallUrl(id, wg)
		go CallUrl1(id,wg)
		go CallUrl2(id,wg)
		go CallUrl3(id, wg)
		go CallUrl4(id,wg)
	}
	wg.Wait()
	diff := time.Since(st)
	fmt.Printf("Total Time taken %d", diff)
}