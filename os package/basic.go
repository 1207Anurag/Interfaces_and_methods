package main

import ("fmt"
"os"
"log"
)

func main(){
	
	emptyFile,err:=os.Create("Anurag.txt")
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%v",emptyFile)
	emptyFile.Close()

	_,erros:=os.Stat("Anurag.txt")
	if err!=nil{
		if os.IsNotExist(erros){
			fmt.Println("File Does not exist")
		}
	}
	fmt.Println("File Exist")

	//changing the permission 
	err=os.Chmod("Anurag.txt",0777)
	if err!=nil{
		fmt.Println(err)
	}else{
	fmt.Println("Permission changed")
	}
}
//if the open fails....error string will be self explanatory

//the files data can be read using slice of bytes


