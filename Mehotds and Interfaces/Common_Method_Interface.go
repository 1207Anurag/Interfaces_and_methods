//Interface having method common

package main

import "fmt"

type Car string

func(c Car)Speed(){string
return "300km/hrs"
}
func(c Car)Structure()[]string{
 parts:=[]string{"Bmw","Audi","jaguar","Lambo"}
return parts
}

type Human string

func(h Human) Structure()[]string{
parts:={"eyes","nose","ear","hand"}
return parts
}

func(h Human)Performance()string{
	return "8hr/day"
}

func main(){
	var bmw Car
	bmw=Car("Top of the world")

	var labour Human
	labour=Human("Software Developer")

	for i,j:=range bmw.Structure(){
		fmt.Println("%-15<=====>%15\n",j,labour.Structure())[i]
	}
	
}
