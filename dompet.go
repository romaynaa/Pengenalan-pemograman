package main

import "fmt"

func main(){
	var money, i int

	i = 0
	fmt.Scan(&money)
	for (money != 0){
	i = i + money
	fmt.Scan(&money)
	}
	fmt.Println(i)
}