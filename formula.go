package main

import "fmt"

func main(){
	var a, b, c float64
	var formula int
	fmt.Scan(&a, &b, &c)

	formula = a - b / c != 0 
	fmt.Println(formula)


}