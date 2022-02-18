package main

import (
	"fmt"
	"io"
	"strings"
)
func main(){
	//creating a new reader
	b:=strings.NewReader("hello, Reader!!")
	r:=make([]byte,8)
	for {
		n,err:=r.Read(b)
fmt.Println(n,err,b)
fmt.Printf(b[:n])
if err==io.EOF{
	break
}
	}
	
}