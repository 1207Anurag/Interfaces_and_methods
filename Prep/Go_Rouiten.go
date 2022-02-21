//GO routiene Example using Url and http request

package main

import (
	"fmt"
	"time"
	"net/http"
	"net/http"
	"log"
	"io/ioutil"
)
func main(){
	go responseSize("https://www.golangprograms.com")
	go responseSize("https://coderwall.com")
	go responseSize("https://stackoverflow.com")
	time.Sleep(10*time.Second)
}
func responseSize(url string){
	//for getting the order of which one is executing first
  fmt.Println("Step1:",url)
  response,err:=http.Get(url)
  if err!=nil{
	  //means there exist an error
	  log.Fatal(err)
  }
  //if the program will be error free
  fmt.Println("Step2:",url)
  defer response.Body.Close()

  fmt.Println("Step3:",url)
  body,err:=ioutil.ReadAll(response.Body)
  if err!=nil{
	  log.Fatal(err)
  }
  fmt.Println("Step4:",len(body))
}