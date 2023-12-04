package main

import "fmt"

func main(){

	var gram, kg, pon, ons float64
	
	fmt.Scan(&gram)

	kg = gram / 1000
	pon = gram / 453.592
	ons = gram / 28.3495

	fmt.Println(kg, pon, ons)
}