//examples for channels

//programm will print the sum of squares and cubes of individual digits of a number

// 123
// square=(1*1)+(2*2)+(3*3)
// cubes=(1*1*1)+(2*2*2)+(3*3*3)
// output=square+cube=50

package main
import "fmt"
func main(){
k:=123
//step 1-> create a channel
sqch:=make(chan int)
cubech:=make(chan int)
go Square(k,sqch)
go Cubes(k,cubech)
squares,cubes:=<-sqch,<-cubech
fmt.Println(squares+cubes)


}
// function which will find the sum of square of the number
func Square(num int,squareop chan int) {
	sum:=0
	for num>0{
	rem:=num%10
	sum+=rem*rem
	num/=10
	}
	squareop <-sum
}
func Cubes(num int,cubeop chan int){
	sum:=0
	for num>0{
		rem:=num%10
		sum+=rem*rem*rem
		num/=10
	}
	cubeop <-sum
}
//function which will find the sum of cubes of the number