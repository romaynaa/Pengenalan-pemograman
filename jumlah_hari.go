package main

import "fmt"

func main(){
	var x int
	fmt.Scan(&x)
	fmt.Println("%d bulan %d minggu %d hari", x / 30 ,(x % 30) / 7, (x % 30) % 7)
}