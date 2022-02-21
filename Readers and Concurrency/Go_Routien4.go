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

func main() {
	st := time.Now()
	wg := new(sync.WaitGroup)
	wg.Add(200)
	for i := 1; i <= 200; i++ {
		id := strconv.Itoa(i)
		go CallUrl(id, wg)
	}
	wg.Wait()
	diff := time.Since(st)
	fmt.Printf("Total Time taken %d", diff)
}