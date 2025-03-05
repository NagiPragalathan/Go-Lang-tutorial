package main

import "fmt"

func main(){
	var a, b string = "Hello", "World"
	fmt.Print(a) // print is used to print a string without a new line
	fmt.Print(b)
	fmt.Println()
	
	fmt.Println("Hello, World!") // println is used to print a new line

	fmt.Printf("a is %v and b is %v", a, b) // formatted string


	var i = 15.5
	var txt = "Hello World!"
  
	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)
  
	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)
}
