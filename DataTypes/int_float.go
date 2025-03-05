/*
there are two types of float:
1. float32
2. float64

*/

package main
import "fmt"

func main(){
	var a float32 = 1.1
	var b float64 = 2.2

	fmt.Println(a, b)

	var x float32 = 123.78
	var y float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)
	
}
