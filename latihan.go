package main

import "fmt"

func main(){
	var bil int
	var status bool

	fmt.Scan(&bil)
	status = bil/2 == 2
	fmt.Println(status)

} 