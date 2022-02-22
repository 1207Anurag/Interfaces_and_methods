//Go Routines 

package main

import ("fmt"
"net/http"
"log"
"io/ioutil"
"sync"
)

func ReadUrl(url string, wg *sync.WaitGroup){
	
	defer wg.Done()
	fmt.Println("Step1 :",url)
	res,err:=http.Get(url)
	if err!=nil{
		log.Fatal(err)
	}
	//when everything is okay
	fmt.Println("Step2: ",url)
	defer res.Body.Close()

	fmt.Println("Step3 :",url)
  body,err:=ioutil.ReadAll(res.Body)
  if err!=nil{
	  log.Fatal(err)
  }
  fmt.Println("Step4 :", len(body))
}
func main(){
	wg:=make(sync.WaitGroup)
	wg.Add(1)

	func(){
		go ReadUrl()

	}()
	wg.Wait()
}