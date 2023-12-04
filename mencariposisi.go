package main

import "fmt"

func main(){
	
	var v, s, t int
	var hasil int

	fmt.Scan(&v, &s, &t)
	hasil = (s * t) + v
	fmt.Println(hasil)
}