package main

import "fmt"

func main(){
	var f, r, k, c float64
	fmt.Scan(&c)
		f = (c * 9 / 5) + 32
		r = c * 4 / 5
		k = c + 273.15
	fmt.Println(f, r, k)
}