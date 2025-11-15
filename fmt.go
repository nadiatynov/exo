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
	h float32
 )
  x = 17
  y ="je suis mineur "
  z = true 
  h = 15.66
 fmt.Printf("j'ai : %vans, et %v, Oui c'est %v!",x,y,z)

 fmt.Println('\n')
 
 fmt.Println(x , y , z)

 fmt.Printf("j'ai : %v de moyenne",h)

}