package main

import "fmt"

func main(){
	var n int
	var kata string

	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
	fmt.Scan(&kata)
	fmt.Print(kata)
}
}