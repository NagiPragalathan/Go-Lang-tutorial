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
}