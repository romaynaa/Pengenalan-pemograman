package main

import "fmt"

func main(){
	var tot_belanja, diskon float64
	var ikut_asesmen bool
	fmt.Scan(&tot_belanja, &ikut_asesmen)
	if ikut_asesmen {
    	diskon = tot_belanja * 0.35
		tot_belanja = tot_belanja - diskon
	}
	fmt.Println(tot_belanja)
}