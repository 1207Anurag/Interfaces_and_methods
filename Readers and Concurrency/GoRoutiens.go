package main

import ("fmt"
       "time"
)

func main(){
	fmt.Println("Main Starts")
	go hello(1)
	go hello(2)
	go hello(3)
	go hello(4)
	go hello(5) 
	//these go routins are running in parllels

	/* After using go in front of hello 
	this will run in its seperate routien
	Main Starts
    hello End       output is this
    go routine give us no output

	this happens because the routien which is keeping track of main is executed much earlier than the routien of the 
	hello function
	*/
	fmt.Println("hello End")
	time.Sleep(100*time.Millisecond)
}
func hello(i int){
	for k:=0;k<4;k++{
		fmt.Println(i,"Hello")
	}
}
//each and every go program has at least one go routin which runs the main function 
