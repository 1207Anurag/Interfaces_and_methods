// when a new Goroutine executed, the Goroutine call return immediately. The control does not wait for 
///Goroutine to complete their execution just like normal function they always move forward to the next line 
//after the Goroutine call and ignores the value returned by the Goroutine. 
package main

import (
	"fmt"
	"time"
)
func main(){
	 fmt.Println("Main start")
    go f("direct")
     fmt.Println("Main finished")
	 time.Sleep(1*time.Second)
//adding sleep() method in programm makes go routine sleep for 1 second
//in between this 1 sec new goroutien executes and display these numbers on the screen
	 

}

func f(from string){
	for i:=0;i<3;i++{
		fmt.Println(from,":",i)
	}
}

