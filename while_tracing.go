package main

import "fmt"

func main(){
	var a, b, reminder int

fmt.Scan(&a, &b)
	for b > 0 {
	reminder = a % b
	a = b
	b = reminder
	}
fmt.Println("GCD(",a,",",b,")=" ,a)
}