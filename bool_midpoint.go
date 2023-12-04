package main

import "fmt"

func main(){
	var a, b, c int
	var status bool
	fmt.Scan(&a, &b, &c)

	status = (((a + b) / 2) == c) || (((a + c) / 2) == b) || (((b + c) / 2) == a)
	fmt.Println(status)
}