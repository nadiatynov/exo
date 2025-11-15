package main 

import (
"fmt" 
"math/rand"	
"time"
)

func main() { 
/*var x int 
x = 9
y:=3*/

/* var (
	x int 
	y string 
	z bool 
 )
  x = 17
  y ="je suis mineur "
  z = true 
  h = 15.669
 fmt.Printf("j'ai : %vans, et %v, Oui c'est %v!",x,y,z)

 fmt.Println('\n')
 
 fmt.Println(x , y , z) 


  for i:= 15; i >=0 ; i-- {
	if i%2 == 0 {
		continue
	}
	fmt.Println(i)
  } */

  //var list [n] type 
  var list [4] int 
  list[0] = 1
  list[1] = 6
  list[2] = 8
  list[3] = 19
fmt.Println(list) 

rand.Seed(time.Now().UnixNano())
var s  [6]int 
for i, h := range s{ 
 s[i]= rand.Int()
	fmt.Println(h) 
}
}