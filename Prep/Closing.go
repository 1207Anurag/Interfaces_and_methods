//Select statement in go lang

package main
import "fmt"
 
func main(){
	if v:=5*10; v>20{
		fmt.Println("v is greter than 20")
	} else{
		fmt.Println("v is less than 20")
	}
	fmt.Println(v)
}