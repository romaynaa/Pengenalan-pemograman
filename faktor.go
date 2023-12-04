package main 

import "fmt"

func main(){
	var n, i, d int
	var s bool

	fmt.Scan(&n)

	for i = 1; i <=n; i++ {
		d = 0 + i
		s = n % i == 0
		fmt.Println(d, s)

	}
	
}