package main

import (
	"fmt"
)

// func add(a int, b int)  { // perameter
// fmt.Println(a + b)
// }
// func main (){ // function call argument
// 	add(1, 2);
// }


func hireOrder(a int,b int,f func(a int, b int) ){ // fhireOrder is a function that takes two integers and a function as parameters
	f(a, b);
}
func add(a int, b int)  {
	 add := a + b
	 fmt.Println(add)
	}
	func main (){
		hireOrder(1, 2, add); 
	}
	