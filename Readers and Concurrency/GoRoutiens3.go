package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func CallUrl(i string) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + i)
	if err != nil {
		fmt.Printf("cannot access provided url")
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
	for i := 1; i <= 10; i++ {
		id := strconv.Itoa(i)
		go CallUrl(id)
	}
	time.Sleep(5* time.Second)
	diff := time.Since(st)   //storing the diffrence
	fmt.Printf("Time take to make 100 requests: %v", diff)
}