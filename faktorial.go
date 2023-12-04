package main

import "fmt"

func main(){

	var n, hasil,i, faktor int
	
	fmt.Scan(&n)
	hasil = n
	for i = 1; i < n; i++ {
		faktor = (n - i)
		hasil *= faktor
	
	}
	
	fmt.Println(float64(hasil))

}