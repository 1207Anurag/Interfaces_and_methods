package main
import (
	"fmt"
)
type shape interface{
	//interface basically groups types together based on their methods
	area() float64  //the struct square has the area function associated with it...so we can call area on it
	circum() float64 //same for this...if not we cannot call the interface on this struct
}
type square struct{
	length float64
}
type circle struct{
	radius float64
}
func(s square) area() float64{
	return s.length*s.length
}
func(s square) circum()float64{
	return s.length*4
}
func(c circle) area()float64{
	return c.radius*c.radius
}
func(c circle) circum()float64{
	return 2*c.radius
}
//s could be either only square type or circle type
//for one instance only one...
//so for 2 things we have to create two diffrent functions

//instead of doing this we can create 
//a interface Line 5

/*func printShapeInfo(s //square/ circle){
	fmt.Println(s,s.area())
	fmt.println(s,s.circum())
}*/

//now changing this above function
func printShapeInfo(s shape){
	fmt.Println(s,s.area())
	fmt.Println(s,s.circum())
}


func main(){
	//now we can create a slice
	shape:=[]shape{
		
		circle{radius:7.5},
		circle{radius:12.3},
		square{length:4.9},
	}
	for _,v:=range shape{
		printShapeInfo(v)
		fmt.Println("...")
	}
}

/*
Interface-> 
-> The interface is a custom type that is used to specify a set of one or more method signatures 
and the interface is abstract, so you are not allowed to create an instance of the interface.
ut you are allowed to create a variable of an interface type and this variable can be assigned 
with a concrete type value that has the methods the interface requires. Or in other 
words, the interface is a collection of methods as well as it is a custom type.
it is necessary to implement all the methods declared in the interface for implementing an interface

-> zero value of interface is null
-> when interface has zero methods such types of interface is called empty interface
-> types of interface -:
static and dynamic
static type is interface itself.....
interface do not contains static values...so it always points to dynamic values
why??-> coz A variable of the interface type containing the value of the Type which implements the interface, so the value of that Type is known as dynamic value and the type is the dynamic type
package main
  
import "fmt"
  
// Creating an interface
type tank interface {
  
    // Methods
    Tarea() float64
    Volume() float64
}
  
func main() {
  
    var t tank
    fmt.Println("Value of the tank interface is: ", t)
    fmt.Printf("Type of the tank interface is: %T ", t)
}

o/p -> Value of the tank interface is:  <nil>
       Type of the tank interface is: <nil> 

Here, in the above example, we have an interface named as a tank. In this example,
fmt.Println(“Value of the tank interface is: “, t) statement returns the dynamic value of the interface and fmt.Printf(“Type of the tank interface is: %T “, t) 
statement returns the dynamic type, i.e nil because here the interface does not know who is implementing it.
 a is the value or the expression of the interface and T is the type also known as asserted type.
  The type assertion is used to check that the dynamic type of its operand will match the asserted type or
   not. 

   Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
*/

