/*
Variables:
int, float32, string, bool.

## in this method we declare the variable type and then the value.
var var_name var_type = value

## in this method we declare the variable without the type and the value.
x := 10 // variable is dynamic type

*/

package main

import "fmt"

func main(){
	var a = 10
	var b = 20
	c := a+b
	fmt.Println(c)

	var name_1 string = "John" // type is string
	var name_2 = "Doe" // type is inferred
	name_3 := "Jane" // type is inferred

	fmt.Println(name_1, name_2, name_3)

	// variable declaration without value
	var d string
	var e int
	var f bool

	fmt.Println(d, e, f)
}