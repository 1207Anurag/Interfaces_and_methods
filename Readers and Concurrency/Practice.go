package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func main(){
for i:=1;i<=100;i++{
	id:=strconv.Itoa(i)
	go callerUrl(id)
}
}
func callerUrl(s string){
res,err:=http.GET("https://jsonplaceholder.typicode.com/posts/"+s)
if err!=nil{
	fmt.Println("Cannot Access the provided Url")
}
else{
	res,err:=ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println("cannot read the content from the body")
	}
	else{
		fmt.Println(string(body))
	}
}
}