package main

import "fmt"

func main(){
	var user, pass string
	var i int

	i = 0
	fmt.Scan(&user, &pass)
	for (user != "admin" || pass != "admin"){
	i = i + 1 
	fmt.Scan(&user, &pass)
	}
fmt.Println(i)
fmt.Println("login berhasil")
}