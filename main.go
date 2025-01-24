package main

import (
	"fmt"
)
var a int = 40
func main (){
	
	fmt.Println(a)
}  

func init() {
	fmt.Println("i'm form init");
	a = 100
	
}
