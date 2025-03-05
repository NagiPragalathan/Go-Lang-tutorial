package main
import "fmt"

func main(){

	const a int = 1
	const b = 2

	fmt.Println(a, b)

	// multiple constants
	const (
		c int = 3
		d = 4
	)

	fmt.Println(c, d)

}