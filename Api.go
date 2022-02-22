//Worker pool
// Concept->  

package main

import "fmt"

type Site struct{
	Url string
}
type Result struct {
	status int
}
func crawl(Worker_id int,jobs <-chan, results <-Result ){
  //jobs receiving, and results-> sending
  for site:=range jobs{
	  fmt.Printf("Worker_ID",Worker_id)
	  res,err:=http.Get(site.url)
	  if err!=nil{
		  log.Fatal(err)
	  }
	  //we are sending the result through the results channel
	  results<-Result{status: res.StatusCode}
  }
}
func main(){
	jobs:=make(chan Site,3)
	results:=make(chan Result,3)
	//using these two channels we are going to communicate with the worker in the worker pool
	for w:=1;w<=3;w++{
		go crawl(w,jobs,results)
	}
	//adding the urls that we have to send the request
	urls:=[]string{
    "https://abc.net",
    "https://google.com",
	}
	//in here we are constantly sending new jobs to the worker pool

	for _,url:=range urls{
		jobs<-Site{Url:url}
	}
	//after we have send each and everything to the channel 
	//we should just close the channel
	close(jobs)

	//for getting the results of these jobs
	for a:=1;a<=2;i++{
		result:=<-results
		fmt.Println(result)
	}
}
