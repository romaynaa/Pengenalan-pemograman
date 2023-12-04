package main

import "fmt"

func main(){

	var tahun int
	var status bool
	fmt.Scan(&tahun)
	status = (tahun % 400 == 0) || (hasil % 4 == 0) && (hasil % 100 != 0)
	
fmt.Println(status)
}
