// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Circle struct {
// 	Radius float64
// }

// func (c Circle) Area() float64 {
// 	return math.Pi * math.Pow(c.Radius, 2)
// }

// func (c Circle) String() string {
// 	return fmt.Sprintf("Circle {Radius: %.2f}", c.Radius)
// }

// type Square struct {
// 	Width  float64
// 	Height float64
// }

// func (s Square) Area() float64 {
// 	return s.Width * s.Height
// }

// func (s Square) String() string {
// 	return fmt.Sprintf("Square {Width: %.2f, Height: %.2f}", s.Width, s.Height)
// }

// type Sizer interface {
// 	Area() float64
// 	fmt.Stringer
// }

// // type Shaper interface {
// // 	Sizer
// // 	fmt.Stringer
// // }

// func main() {
// 	c := Circle{Radius: 10}
// 	PrintArea(c)

// 	s := Square{Height: 10, Width: 5}
// 	PrintArea(s)

// 	l := Less(c, s)
// 	fmt.Printf("%v is the smallest\n", l)

// }

// func Less(s1, s2 Sizer) Sizer {
// 	if s1.Area() < s2.Area() {
// 		return s1
// 	}
// 	return s2
// }

// func PrintArea(s Sizer) {
// 	fmt.Printf("area of %s is %.2f\n", s.String(), s.Area())
// }

package main

import (
	"fmt"
	"strconv"
)

//like in java we can use try catch to handle the errors but in golang we use to have a function which 
//handles error 
func main(){
	//handling the error
	//if error is not nill it means program contains the error
	//if it is nill it means no error
	result,err:=returnError(false)
	if err!=nil{
    fmt.Println(err)
	} else{
	fmt.Println(result)
	}
}

type specialError struct{
	errorMessage string
	errorCode int
}
//error which is present in the line 96 is basically an interface....which we are going to implement 
//insde this function and the reviver will be the cutomized error...and Error()-> is the method which is 
//present inside the error interface
func (s specialError)Error() string{
	return s.errorMessage+" "+ strconv.Itoa(s.errorCode)
}

func returnError(returnError bool) (string,error){
	//a function whose return type is also an error
  if returnError{ //if there is any error
	 // return "",errors.New("Error Here")
	 //using special errors
	 return "",specialError{errorMessage:"Special Error",errorCode:123}
  }
	return "random results",nil
	//built in libraries of this 
//this is how we return errors in go	
}