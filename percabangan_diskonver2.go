package main

import "fmt"

func main(){
	var diskon, tot_belanja float64
	var ikut_asesmen bool
	fmt.Scan(&tot_belanja, &ikut_asesmen)
	if ikut_asesmen {
		diskon = tot_belanja * 0.35
	}else{
		diskon = 0
	}
	fmt.Println(tot_belanja - diskon)
}