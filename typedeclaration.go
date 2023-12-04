package main

import"fmt"

func main(){

	type NoKTP string

	var noKtpRoma NoKTP = "23432894932094678"
	fmt.Println(noKtpRoma)

	var i = 10
	i += 10 // i = i + 10 
	fmt.Println(i)

}