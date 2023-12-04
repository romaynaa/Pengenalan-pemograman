package main

import "fmt"

func main (){
	var bil, dig1, digmid, dig3, dig4 int
	var disc, cash bool
	fmt.Scan(&bil)

	dig1 = bil / 1000
	digmid = bil % 1000 / 10
	dig3 = bil % 100 / 10
	dig4 = bil % 10
	disc = digmid % 2 == 0
	cash = (dig1 + dig3) % dig4 == 0 && dig4 != 0

	fmt.Println(disc, cash)

}

