package main

import "fmt"

func main(){

	var n, fn, fn_1, fn_2 int 
	fmt.Scan(&n)

	
	fn_1 = 0
	fn_2 = 1
	fn = fn_1 + fn_2
	fmt.Println(fn_1)
	for i := 1; i <= n; i++ { 
		fn = fn_1 + fn_2
		fn_2 = fn_1
		fn_1 = fn
		fmt.Println(fn)
}
	
	
}