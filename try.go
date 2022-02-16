//10-> Implicit Declaration 
package main
import "fmt"

// type I interface{
// 	F()
// }

// func main(){
// var i I=T{"Hello"}
// i.F()
// }
// //for creating method first create a struct

// type T struct{
// 	S string
// }
// func (t T) F(){
// 	fmt.Println(t.S)
// }

//11-> Interface Values
// Interface values can be thought as tuple of values and concrete type
 //(value,type)

//if we call method on interface..value executes the methods of same name on its underlying type


//12-> Interface values with nil underlying values


// package main

// import "fmt"

// func main(){
// var i I
// var t*T
// i=t
// describe(i)
// i.M()
// }
// //if a concrete value inside the interface is itself nil
// //then the method will called as nil reciever

// type T struct{
//    S string
// }
// func(t *T) M(){
// 	if t==nil{
// 		fmt.Println("<nil>")
// 		return
// 	}
// 	fmt.Println(t.S)
// }
// func describe(i I){
// 	fmt.Println("(%v, %T)",i,i)
// //as soon as i is triggerd...it will check go to the interface and check which method has appropiate methods and data type

// }
// type I interface{
// 	M()
// }

//13-> Nil interface

// package main

// import "fmt"
// type I interface{
//     M()
// }
// func main(){
// 	var i I
//      describe(i)
// 	 i.M()
// }

// func describe(i I){
// 	fmt.Println("(%v, %T)",i,i)
// }
//Run Time error

//14-> Empty Interface 

// func main(){
// 	var i interface{}
// 	describe(i)
// 	i=42
// 	describe(i)

// 	i="hello"
// 	describe(i)

// }
// func describe(i interface{}){
// 	fmt.Printf("(%v,%T)",i,i)
// }

//15-> Type asserstion 

//to access the values from the interface we use type assersion 

//t:=i.(T)--> interface value i holds concrete type T and assings underlying T value to variable t.


//if i doest not hold a t, statement will trigger a panic

// func main(){
// 	var i interface{}="Hello World"
// 	s:=i.(string)
// 	fmt.Println(s)
// 	//because we have a concrete value which is Hello world
// 	s,ok:=i.(string)
// 	fmt.Println(s,ok)
// 	f,ok:=i.(int)
// 	fmt.Println(f,ok)
// 	k := i.(float64) // panic
// 	fmt.Println(k)
// }

//16->Type switches

//using type switches we can construct that permits several type assertions in series
// func do(i interface{}){
// 	switch v:=i.(type){
// 	case int:
// 		fmt.Println("twice %v\n",v,v*2)
// 	case string:
// 		fmt.Println("%q is %v bytes long\n",v,len(v))
// 	}
// default:
// 	fmt.Println("i dont know the type\n",v)
// }

//17-> Stringers

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
// }

// func main() {
// 	a := Person{"Arthur Dent", 42}
// 	z := Person{"Zaphod Beeblebrox", 9001}
// 	fmt.Println(a, z)
// }

// func add(i,j interface{}){
// 	fmt.println(i+j)
// }


//Interface--> groups the methods which are gaving same properties in it


type I interface{
	printing()   
}
type T struct{
	S string
}
func(t T)printing(){
	fmt.Println(t.S)
}
type F float64
func(f F)printing(){
fmt.Println(f)
}
//one more function which will tell me the type and the value which we are gettin from the interface

func decide(i I){
	fmt.Printf("(%v,%T)\n",i,i)
}

func main(){
	var i I
	i=T{"Hello"}   //type struct
	decide(i)
	i.printing()

	i=F(56.3214)
	decide(i)
	i.printing()

}