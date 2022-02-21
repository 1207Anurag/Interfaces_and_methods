package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.waitGroup

func responseSize(url string){
	
//schedule the call to waitGroup to tell that goroutine is completed
defer wg.Done()

fmt.Println("Step1: ",url)
response,err:=http.Get(url)
if err!=nil{
	log.Fatal(err)
}
fmt.Println("Step 2: ",url)
defer response.Body.close()

fmt.Println("Step3 :",url)
body,err:=ioutil.ReadAll(response.Body)
if err!=nil{
	log.Fatal(err)
}
fmt.Println("Step4 :",len(body))
}
func main(){
wg.add(3)
fmt.Println("Start GoRoutine")
go responseSize("https://www.golangprograms.com")
go responseSize("https://stackoverflow.com")
go responseSize("https://coderwall.com")

wg.wait()
fmt.Println("Terminating programm")
}