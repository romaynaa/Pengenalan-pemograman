package main

import "fmt"

func main(){
	var R, r, s int
	var status bool
	fmt.Scan(&R, &r, &s)

	status = (R + r) < s 
	fmt.Println(status)
}