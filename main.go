package main

import "fmt"
func concat(s1 string, s2 string) string {
	return s1 + s2
}
func main (){
	fmt.Println("hello world!")

	// variable 
	age:=20

	// if else
	if age > 17{
		fmt.Println("you can married")
	} else{
		fmt.Println("can't")
	}

	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")

}  


func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}
