package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)

	for i := n; i >= 1; i-- {
		fmt.Print(i)
	}
	fmt.Println()
}
