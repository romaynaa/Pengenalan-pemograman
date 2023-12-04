package main

import "fmt"

func main(){
	var x, n int
		n = 0
		for x % 2 == 0 {
		n = n + x
		fmt.Scan(&x)
		}
	fmt.Println(n)



	var x, n int
	var oke bool
		n = 0
		
		for !oke {
			n = n + x
			oke = (n%2 != 0)
		}
		fmt.Println(n)

	}
