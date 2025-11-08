package main 

import "github.com/01-edu/z01"
	
func main() {
	for i := 'a'; i <= 'z'; i++ { 
  z01.PrintRune(i)
   z01.PrintRune(' ') 
	}
 z01.PrintRune('\n')
	for j := 'z'; j >='a'; j-- {
		z01.PrintRune(j)
		z01.PrintRune(' ')
		
	}
  z01.PrintRune('\n')
	for m := '0'; m <= '9'; m++ {
		z01.PrintRune(m)
		z01.PrintRune(' ')
		
	}

}