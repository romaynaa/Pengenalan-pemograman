package main 

import "fmt"

func main(){
	var x, n int
	var oke bool
		n = 0
		
		for !oke {
			n = n + x
			oke = (n%2 != 0)
		}
		fmt.Println(n)

	
}