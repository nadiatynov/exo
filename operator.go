package main 

import "fmt"

func main() {
	var (
		x int 
		y int
		z int 
	)
	x = 17
	y = -3
	z = 6
	fmt.Println(x + y + z)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x % y) 
	fmt.Println(x / y)
}