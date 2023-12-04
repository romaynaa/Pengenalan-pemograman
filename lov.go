package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := n; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			fmt.Print(" ")
		}
		for k := n; k >= (2*i - n); k--{
			fmt.Print("*")
		}
		fmt.Println()

}
}