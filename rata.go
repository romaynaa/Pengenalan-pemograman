package main

import "fmt"

func main(){
	var i, n, jam, jumlah int

	fmt.Scan(&n)

	for i = 1; i <= n; i++ {
	fmt.Scan(&jam)
	jumlah = jumlah + jam
}
	fmt.Println(jumlah/n)
}