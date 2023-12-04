package main

import "fmt"

func main(){

	var a, b, c string
	var status bool

	fmt.Scan(&a, &b, &c)
	status = (a == b || b == c || c == a)
	fmt.Println(status)
}