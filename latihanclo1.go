package main

import "fmt"

func main(){

	var x1, x2, x3, x4, x5 int
	var y1, y2, y3, y4 string

	fmt.Scan(&x1, &x2, &x3, &x4, &x5)
	y1 = (x1 + x2) % 4096 
}
