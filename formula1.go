package main

import "fmt"

func main(){
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	fmt.Println(int((a - b) / c))
}