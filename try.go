package main

import "fmt"

func main(){
v:=vertex{3,4}
fmt.Println(func_1())
scale(5)
fmt.Println()

}
type vertex struct{
	X,Y float64
}
func(v vertex)func_1(f float64){
v.X=v.X*f
v.Y=v.Y*f
}
func scale(v vertex,f float64) float64{
return math.sqrt(v.X*v.X+v.Y*v.Y)+f;
}