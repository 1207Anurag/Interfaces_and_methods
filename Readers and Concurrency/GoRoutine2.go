package main

import ("fmt"
       "time"
)

func main(){
	fmt.Println("Main Starts")
	msg:="hello"
	go func(m string){   //later updating
		fmt.Println(m)
	}(msg)   //updating later
	//after passing the argument if we try to change the
	//value it wont affect and prevent programm from the race condition 
	msg="world"  //later updating
	fmt.Println("hello End")
	 time.Sleep(100*time.Millisecond)
}
func hello(i int){
	for k:=0;k<4;k++{
		fmt.Println(i,"Hello")
	}
}
//spinning up go routin using anonymous function
//what we have done in this programm is we have introduced...a race condition in which same piece of data is accessed
//in the main go routien as well as in routien in the anymous function as well

//race cond is not helthy...coz they will give us unexpeted output and application will end up having lot of bugs
//how to fix this--> pass an arugument in this anonymous function 
