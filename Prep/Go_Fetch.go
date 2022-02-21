package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)
var wg sync.waitGroup

func resSize(url string,num chan int){
	defer wg.Done()

	response,err:=http.Get(url)
	if err!=nil{
		log.Fatal(err)
	}
	defer response.Body.close()

	body,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		log.Fatal(err)
	}
	nums<-len(body)sudo docker exec -it mysql bash

}
func main(){
	nums:=make(chan int)
    wg.Add(1)
	go resSize("https://www.golangprograms.com",nums)
	fmt.Println(<-nums)  //getting the content form the channel
	wg.wait()
	//now the channel is continously getting the values 
	//so we have to close the connection 
	close(nums)

}