package main 

import "github.com/01-edu/z01"
	
func main() {
	for i := 'a'; i <= 'z'; i++ { 
  z01.PrintRune(i)
   z01.PrintRune(' ') 
	}

	for j := 'z'; j >='a'; j-- {
		z01.PrintRune(j)
		z01.PrintRune(' ')
		
	}

}