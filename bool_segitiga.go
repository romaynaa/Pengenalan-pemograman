package main

import "fmt"

func main(){
	var a, b, c int // suatu angka dapat membentuk segitiga
	//apabila dua garis terkecil jika dijumlahkan hasilnya
	//lebih besar dari garis yang diperpanjang//
	var status bool
	fmt.Scan(&a, &b, &c)
	status = (a + b > c) && (a + c > b) && (b + c > a)
	fmt.Println(status)

}