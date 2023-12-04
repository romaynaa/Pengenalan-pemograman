package main

import "fmt"

func main(){
	var kartu, diskon, cashback, status bool
	var n int
	fmt.Scan(&n, &status)

	kartu = status 
	diskon = n >= 100000
	cashback = (n >= 200000) && kartu
	fmt.Println(kartu, diskon, cashback)
}