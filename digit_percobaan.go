package main

import "fmt"

func main(){
	var x int
	fmt.Scan(&x)

	angka_pertama := x / 1000 % 100
	angka_kedua := x / 100 % 10
	angka_ketiga := x / 10 % 0
	angka_keempat := x 
	fmt.Println("%d%d%d%d", angka_keempat, angka_ketiga, angka_kedua, angka_pertama)
} infalid