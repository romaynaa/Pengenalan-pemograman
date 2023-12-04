package main

import "fmt"

func main (){
	var username, password string
	
	fmt.Scan(&username, &password)
	fmt.Println(len(username) > 6 && 
		username[0:6] == "@oplib" && len(username) <=15 && len(username) >= 8 && 
		(username[6] >=48 && 
		username[6] <=57 || 
		username[6] >=65 && 
		username[6] <= 90) && 
		username[6:] != password)
		


}