package main

import "fmt"

func main(){

}
func fibonacci(c,quit chan int){
	x,y:=0,1
	for{
		select{
		case c<-x:   //Receving
			x,y=y,x+y
		
	case <-quit:
		fmt.Println("quit")
		return 
	}
}
}
func main(){
	c:=make(chan int)
}