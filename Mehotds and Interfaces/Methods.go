
//with help of reciver argument, method can access the properties of reciver
//method-> struct and non struct type both

//And you are not allowed to create a method in which the receiver type is already defined in another package including inbuilt type like int, string, etc. If you try to do so, then the compiler will give an error.

//write a programm, in which if we make changes in the reciver it will reflect in the output
package main

import "fmt"
type author struct{
	name string
	branch string
}
func(a *author) myFunc_1(abranch string){
	(*a).branch=abranch
	}
	
	
	
func main(){
res:=author{
	name:"Anurag",
	branch:"IT",
}
p:=&res
p.myFunc_1("aerospace")
fmt.Println(p.branch)
k:=res
k.myFunc_1("mechanial")
fmt.Println(k.branch)



}
//method for the pointers
/*// method contais reciver argument in it
//with help of reciver argument, method can access the properties of reciver
//method-> struct and non struct type both

//And you are not allowed to create a method in which the receiver type is already defined in another package including inbuilt type like int, string, etc. If you try to do so, then the compiler will give an error.

//write a programm, in which if we make changes in the reciver it will reflect in the output
package main
 
import "fmt"
 
// Author structure
type author struct {
    name      string
    branch    string
}
 

func (a *author) show(abranch string) {
    (*a).branch = abranch
}
 
// Main function
func main() {
 
    // Initializing the values
    // of the author structure
    res := author{
        name:   "Sona",
        branch: "CSE",
    }
    // Creating a pointer
    p := &res
 
    // Calling the show method
    p.show("ECE")
    fmt.Println("Author's name: ", res.name)
    fmt.Println("Branch Name(After): ", res.branch)
	k:=res
	k.show("mechanical")
	fmt.Println(k.branch)
}*/