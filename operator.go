package main 

import (
	"fmt"	
)

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

 /*
  rand.Seed(time.Now().UnixNano())
	if s := rand.Int() ; s % 2 == 0 {
		fmt.Println(s  , "est paire")
		}else{ 
    fmt.Println( s ,"est impaire") 
	}


	h := rand.Int() % 2
	if h == 0 {
	fmt.Println(h, "paire !")
	} else { 
	fmt.Println( h, "impaire !") 
	}
	*/
	age := 16
	if age > 18 {
		fmt.Println("je suis majeur")
     }else if age == 18{ 
	fmt.Println("je suis tout juste majeur")
	 }	else { 
		fmt.Println("je suis mineur")
	}
}