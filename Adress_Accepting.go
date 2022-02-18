//Interface accepting the address 

package main

import "fmt"

func main(){

}

type Book struct {
	author,title string
}
type magazine struct{
	title string
	issue int
}
func(b *Book)Assign(n,t string){
	b.author=n
	b.title=t
}
func(b *Book)print(){
	fmt.println("Author:%s, Title: %s\n",b.author,b.title)

}
func(mag *magazine)Assign(t string,i int){
	mag.title=t
	mag.issue=i
}
func(mg *magazine)print(){
	fmt.Println("Title:%s, issue:%s\n",mg.title,mg.issue)
}
type Printer Interface{
	print()
}
func(b Book)Stringer()string{
return fmt.Sprintf(b.Author,b.title)
}
type Stringer interface{
	Stringer()
}
func decision(s Stringer){
	fmt.Println(b.Author,b.title)
}
func main(){
//declare instance of book

var b Book
var m Magazine
//if we dont use this functions to assign the values then
// b:=Books{"Helmet of pippers","Book of rabbits"}
// m:=Magazine{"rabbit Weekly",20}
//assigning the values via methods
b.Assign("Helmet of pippers","book of rabbits")
m.Assign("rabbit Weekly",20)
 
var i Printer
fmt.Println("calling the interfae")
//we are performing the type assertions
i=&b
i.Print()
i=&m
i.Print()

}