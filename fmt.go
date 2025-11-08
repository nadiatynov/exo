package main 

import "fmt" 

func main() { 
/*var x int 
x = 9
y:=3*/

 var (
	x int 
	y string 
	z bool
 )
  x = 17
  y ="je suis mineur "
  z = true 
 fmt.Printf("j'ai : %vans, et %v, Oui c'est %v!",x,y,z)

}