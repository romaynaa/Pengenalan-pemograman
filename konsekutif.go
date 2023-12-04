package main

import "fmt"

func main(){
	var x, i, dig1, dig2 int
	var status bool

	fmt.Scan(&x)
	i = 0
	for x > 0 {
		dig1 = x / 10
		dig2 = x / 10 % 10
		status = (dig1 - dig2) == 1 || (dig2 - dig1) == 1
		x = x / 10
		i = i + 1
	
	}
	fmt.Println(status)
}