// //select

// package main

// import "fmt"

// func main(){
// ninja1,ninja2:=make(chan string),make(chan string)
// go ElectCaptian(ninja1,"Ninja1")
// go ElectCaptian(ninja2,"Ninja2")

// p1:=<-ninja1
// p2:=<-ninja2
// fmt.Println(p1,p2)
// }
// func ElectCaptian(ninja chan string,message string){
// ninja<-message
// }

package main

import "fmt"

func main(){
ninja1,ninja2:=make(chan string),make(chan string)
//this channel would'nt really help us to run the elections
//in term of this we are going to use select to do the same

go ElectCaptian(ninja1,"Ninja1")
go ElectCaptian(ninja2,"Ninja2")

// in this condition it will print only one single message


select{
case message:=<-ninja1:
	fmt.Println(message)
case message:=<-ninja2:
	fmt.Println(message)
}
}
func ElectCaptian(ninja chan string,message string){
ninja<-message
}

//as we run this programm again and again 
//eventually we will end up getting ninja1 tooo
//roughly we get 50-50 changes of getting ninja1 and ninja2
