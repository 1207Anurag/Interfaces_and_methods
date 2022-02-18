//channels 
//bydirectional in nature by default
//only can transfer data of  same type

package main

import "fmt"

// func main(){
// //channel
// fmt.Println("Main started")
// ch:=make(chan int)    //creating a channel
// go myChannel(ch)     // func-> calling the my channel routine
// ch<-23               //sending this data to the channel (main)
// fmt.Println("Main finished")

// }
// func myChannel(ch chan int){
// 	fmt.Println(237+<-ch)  //reciving the value from the main channel
// }

//close()-> Channel


// func main(){
// c:=make(chan string)
// go myChan(c)
// for{
// 	res,ok:=<-c //receiving the data from the channel
// 		if ok==false{
// 			fmt.Println("The channel is Closed",ok)
// 			break
// 		}
// 		fmt.Println("The channel is opened ",res,ok)
	
// }
// }
// //creating a function which is going to send the string value to the  main goRoutine
// func myChan(ch chan string){
// 	for v:=0;v<4;v++{
// 	ch<-"Anurag"
// 	}
// 	close(ch)
// }   //it will send the data until the loop reaches to 4 after that the channel will be automatically closed

// func main(){
// myChan:=make(chan string,4)  //creating a channel

// 	myChan<-"Gfg"
// 	myChan<-"is "
// 	myChan<-"The"
// 	myChan<-"Best"
// 	fmt.Println("Length of the channel is", cap(myChan))
	
// //calling the anonymous function

// }

// func main(){
// s:=[]int{7,2,8,-9,4,0}
// c:=make(chan int)     //creating a channel
// go sums(s[:2],c)  //9
// go sums(s,c)     //12
// go sums(s[:4],c)  //8
// go sums(s[1:4],c) //1
// x:=<-c
// y:=<-c
// z:=<-c
// fmt.Println(x,y,z,x+y+z)  //ans-> 1+9=10
// }
// func sums(s []int,c chan int){
// 	sum:=0
// 	for _,v:=range s{
// 		sum+=v
// 	}
// 	c <-sum   //send the sum to c
// }


func main(){
c:=make(chan int,10)
go fibo(cap(c),c)
for{
	res,ok:=<-c
		if ok==false{
			fmt.Println("got an error",ok)
			break
		}
		fmt.Println(res)
}
}
func fibo(n int,c chan int){
	x,y:=0,1
	for i:=0;i<n;i++{
		c<-x
		x,y=y,x+y
	}

}