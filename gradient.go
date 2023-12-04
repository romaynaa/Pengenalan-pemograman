package main

import "fmt"

func main(){
	var x1, y1, x2, y2 float64
	var nilai_gradient float64

	fmt.Scan(&x1, &y1, &x2, &y2)
	nilai_gradient = (y1 - y2) / (x1 - x2)
	fmt.Println(nilai_gradient)



}