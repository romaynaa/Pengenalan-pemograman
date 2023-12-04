package main

import "fmt"

func main(){
	var t, jam, menit, detik int
	fmt.Scan(&t)
	jam = t / 3600 
	menit = t % 3600 / 60
	detik = t % 60

	fmt.Println(jam, menit, detik)

	var 
}