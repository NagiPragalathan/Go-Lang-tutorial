package main
import "fmt"

func main(){
	var a int = 1
	var b uint = 2
	var c int8 = 3
	var d uint8 = 4
	var e int16 = 5
	var f uint16 = 6
	var g int32 = 7
	var h uint32 = 8

	fmt.Println(a, b, c, d, e, f, g, h)
	fmt.Printf("a is %T, b is %T, c is %T, d is %T, e is %T, f is %T, g is %T, h is %T", a, b, c, d, e, f, g, h)
	
}
