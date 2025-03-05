package main
import "fmt"

func main(){
	var a int = 1
	var b float64 = 2.5
	var c string = "Hello"
	var d bool = true

	fmt.Println(a, b, c, d)
	fmt.Printf("a is %T, b is %T, c is %T, d is %T", a, b, c, d)
	
}