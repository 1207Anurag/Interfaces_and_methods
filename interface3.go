//type assersion example

package main

import "fmt"
/*func myFun(a interface{}){
	val:=a.(string)
	fmt.Println(val)
}
func main(){
	var val interface{

	}="Geeks for geeks"
	myFun(val)
}
*/
func main(){
//perform type assersion
var val interface{
}="geeks for geeks"
myFun(val)

}
func myFun(a interface{}){
	val,ok:=a.(string)
	fmt.Println(val,ok)
}
//use value and ok to prevent from program panic