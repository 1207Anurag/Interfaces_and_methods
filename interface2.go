package main

import "fmt"

func main(){
var t tanker
t=tank{10,14}
fmt.Println(t.area())
fmt.Println(t.volume())
}

//every tanker will have same property which is volume and area.....so we can create a interface 
type tanker interface{
	volume() float64
	area() float64
}
//for calculating the volume of the tank
func(t tank) volume()float64{
	return 2*t.radius*t.height
}
func(t tank) area()float64{
	return 3.14*t.radius*t.height
}

// for calculating the area of the tank

//properties of tank
type tank struct{
	height float64
	radius float64
}