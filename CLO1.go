package main

import "fmt"

func main(){
	var n, dig1, dig2 int
	var status bool

	fmt.Scan(&n)
	
	for n > 0 {
		dig1 = n / 10 % 10
		dig2 = n % 10 
		status = (dig1 - dig2) == 1 && (dig2 - dig1) == 1
		n = n / 10
		
	

	}
	



	fmt.Println(status)
}